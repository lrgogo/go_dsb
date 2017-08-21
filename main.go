package main

import (
	"app/api"
	"log"
	"time"
	"net/http"
	"app/db"
)



func main() {
	db.Open()
	defer db.Close()

	api.Load()

	log.Println("server run at ", time.Now())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
