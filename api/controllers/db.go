package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Db is a database handler
var Db *sql.DB

func Init() {
	var err error
	// env から取得する
	Db, err = sql.Open("mysql", "game:game@tcp(db:3306)/game_db")
	if err != nil {
		log.Panic(err)
	}
}
