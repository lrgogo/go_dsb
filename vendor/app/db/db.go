package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Open() {
	DB, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/dsb")
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db open ping success")
}
