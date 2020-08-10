package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("mysql", "root:abc123@tcp(localhost:3306)/user_demo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
