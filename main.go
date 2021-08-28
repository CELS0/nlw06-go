package main

import (
	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
