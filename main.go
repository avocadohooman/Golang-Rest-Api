package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/avocadohooman/Golang-Rest-Api-Todo/handlers"
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

	r := handlers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" { 
		port = "5001"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v\n", err)
	}
}