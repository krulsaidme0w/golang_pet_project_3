package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"krulsaidme0w/library/internal/pkg/config"
	"krulsaidme0w/library/internal/pkg/storage"
)

const (
	count        = 1000
	maxOpenConns = 1000
	configPath   = "./"
	configName   = "config"
)

func main() {
	if err := config.Init(configPath, configName); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("psql.host"),
		viper.GetString("psql.port"),
		viper.GetString("psql.user"),
		viper.GetString("psql.password"),
		viper.GetString("psql.dbname")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(maxOpenConns)

	err = storage.FillLibraryDB(db, count)
	if err != nil {
		log.Fatal(err)
	}
}
