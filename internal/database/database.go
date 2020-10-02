package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Db id database handler
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "game:game@tcp(game-db:3306)/game-db?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
}
