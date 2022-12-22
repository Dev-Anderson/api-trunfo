package main

import (
	"api-trunfo/database"
	"api-trunfo/server"
)

func main() {
	database.GetDatabase()
	s := server.NewServer()
	s.Run()
}
