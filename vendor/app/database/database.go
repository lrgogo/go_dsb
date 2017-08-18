package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var SQL *sqlx.DB

func DSN() string {
	return "root:123456@tcp(localhost:3306)/dsb?parseTime=true"
}

func Connect() {
	var err error
	if SQL, err = sqlx.Connect("mysql", DSN()); err != nil {
		log.Println(err)
		return
	}
	if err = SQL.Ping(); err != nil {
		log.Println(err)
		return
	}
	log.Println("database connect success")
}
