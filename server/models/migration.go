package models

import (
	"context"
	"database/sql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/database"
	"github.com/vavilen84/class_booking/helpers"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Migration struct {
	Version   int64  `json:"version" column:"version" validate:"required"`
	Filename  string `json:"filename" column:"filename" validate:"required"`
	CreatedAt int64  `json:"created_at" column:"created_at" validate:"required"`
}

func (Migration) GetTableName() string {
	return constants.MigrationsTableName
}

func (m Migration) GetId() string {
	// TODO ID not needed for Migration
	return ""
}

func getMigration(info os.FileInfo) (err error, m Migration) {
	filename := info.Name()
	splitted := strings.Split(info.Name(), "_")
	version, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Println(err)
		return
	}

	m = Migration{
		Filename:  filename,
		Version:   int64(version),
		CreatedAt: time.Now().Unix(),
	}
	return
}

func getMigrations() (err error, keys []int, list map[int64]Migration) {
	list = make(map[int64]Migration)
	keys = make([]int, 0)

	err = filepath.Walk("../"+constants.MigrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			helpers.LogError(err)
		}
		if !info.IsDir() {
			err, migration := getMigration(info)
			if err != nil {
				log.Println(err)
				return err
			}
			list[migration.Version] = migration
			keys = append(keys, int(migration.Version))
		}
		return nil
	})
	if err != nil {
		log.Print(err.Error())
		return
	}

	sort.Ints(keys)
	return
}

func MigrateUp(db *sql.DB) error {
	err, keys, list := getMigrations()
	for _, k := range keys {
		err = apply(db, k, list)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func performMigrateTx(db *sql.DB, m Migration) error {
	ctx := context.TODO()
	tx, beginTxErr := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if beginTxErr != nil {
		log.Fatal(beginTxErr)
		return beginTxErr
	}

	execErr := database.Insert(db, m)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}

	content, readErr := ioutil.ReadFile("../" + constants.MigrationsFolder + "/" + m.Filename)
	if readErr != nil {
		log.Fatal(readErr)
		return readErr
	}
	sqlQuery := string(content)
	_, execErr = tx.Exec(sqlQuery)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func apply(db *sql.DB, k int, list map[int64]Migration) error {
	m := list[int64(k)]
	row := db.QueryRow(`SELECT version FROM `+constants.MigrationsTableName+` WHERE version = ?`, m.Version)
	var version int64
	err := row.Scan(&version)
	if err == sql.ErrNoRows {
		err = performMigrateTx(db, m)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func CreateMigrationsTableIfNotExists(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS ` + constants.MigrationsTableName + `
		(
			version integer PRIMARY KEY,
			filename text NOT NULL,
			created_at integer NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
