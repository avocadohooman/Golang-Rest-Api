package main

import (
	"github.com/avocadohooman/Golang-Rest-Api-Todo/postgres"
	"github.com/go-pg/pg/v9"
)

func main() {
	DB := postgres.New(&pg.Options{
		User: "postgres",
		Password: "postgres",
		Database: "todo_dev",
	})

	// a defer statment defers the execution of a function until the surrounding function returns
	defer DB.Close()
}