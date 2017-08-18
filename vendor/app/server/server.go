package server

import (
	"log"
	"net/http"
	"time"
)

func Run() {
	log.Println("server run at ", time.Now())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
