package main

import (
	database "webapi/databases"
	"webapi/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
