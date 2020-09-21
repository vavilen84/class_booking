package main

import (
	"flag"
	"fmt"
	"github.com/vavilen84/class_booking/constants"
	"github.com/vavilen84/class_booking/helpers"
	"os"
	"path/filepath"
	"time"
)

func main() {
	namePtr := flag.String("n", "", "migration file name")
	flag.Parse()

	now := time.Now()
	nowUnix := now.Unix()

	file := filepath.Join(constants.MigrationsFolder, fmt.Sprintf("%d_%s.up.sql", nowUnix, *namePtr))
	_, err := os.Create(file)
	if err != nil {
		helpers.LogError(err)
	}
}
