package repository

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/github-timall/stat/config"
	"log"
)

var (
	db *sql.DB
	err error
)

func init() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Yml.Database.Host,
		config.Yml.Database.Port,
		config.Yml.Database.User,
		config.Yml.Database.Password,
		config.Yml.Database.DbName,
	)
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Printf("FAILURE: %v", err)
	}
}