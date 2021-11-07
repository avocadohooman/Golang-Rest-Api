package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/avocadohooman/Golang-Rest-Api-Todo/domain"
	"github.com/avocadohooman/Golang-Rest-Api-Todo/handlers"
	"github.com/avocadohooman/Golang-Rest-Api-Todo/postgres"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv" // for accessing .env variables
)

func main() {
	godotenv_err := godotenv.Load()

	if godotenv_err != nil {
		fmt.Printf("Error: %s", godotenv_err)
	}

	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "todo_dev",
	})

	// a defer statment defers the execution of a function until the surrounding function returns
	defer DB.Close()

	domainDB := domain.DB{
		UserRepo: postgres.NewUserRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	http_err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if http_err != nil {
		log.Fatalf("cannot start server %v\n", http_err)
	}
}
