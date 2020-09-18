package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/models"
	"github.com/vavilen84/class_booking/orm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ctx context.Context

func main() {
	db, err := sql.Open(os.Getenv("SQL_DRIVER"), os.Getenv("SQL_DSN"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = createMigrationsTableIfNotExists(db)
	if err != nil {
		log.Println(err)
	}

	err = migrateUp(db)
	if err != nil {
		log.Println(err)
	}
}

func getMigration(info os.FileInfo) (err error, m models.Migration) {
	filename := info.Name()
	splitted := strings.Split(info.Name(), "_")
	version, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Println(err)
		return
	}

	m = models.Migration{
		Filename:  filename,
		Version:   int64(version),
		CreatedAt: time.Now().Unix(),
	}
	return
}

func getMigrations() (err error, keys []int, list map[int64]models.Migration) {
	list = make(map[int64]models.Migration)
	keys = make([]int, 0)

	err = filepath.Walk("./"+constants.MigrationsFolder, func(path string, info os.FileInfo, err error) error {
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

func migrateUp(db *sql.DB) error {
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

func performMigrateTx(db *sql.DB, m models.Migration) error {
	ctx = context.TODO()
	tx, beginTxErr := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if beginTxErr != nil {
		log.Fatal(beginTxErr)
		return beginTxErr
	}

	execErr := orm.Insert(db, m)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
		return execErr
	}

	content, readErr := ioutil.ReadFile("./" + constants.MigrationsFolder + "/" + m.Filename)
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

func apply(db *sql.DB, k int, list map[int64]models.Migration) error {
	m := list[int64(k)]
	row := db.QueryRow(`SELECT id FROM public.`+constants.MigrationsTableName+` WHERE version = $1`, m.Version)
	var id int64
	err := row.Scan(&id)
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

func createMigrationsTableIfNotExists(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS public.` + constants.MigrationsTableName + `
		(
			id serial PRIMARY KEY,
			filename text NOT NULL,
			version integer NOT NULL,
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
