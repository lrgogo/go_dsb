package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var mydb *sql.DB

func Open() {
	mydb, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/dsb")
	if err != nil {
		log.Fatal(err)
	}
	err = mydb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db open ping success")
}

func Close() {
	mydb.Close()
}
