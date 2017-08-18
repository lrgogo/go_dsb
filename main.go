package main

import (
	"app/database"
	"app/server"
	"app/controller"
)

func main() {
	database.Connect()
	controller.Load()
	server.Run()
}
