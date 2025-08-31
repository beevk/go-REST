package main

import (
	"log"
	"net/http"
	"os"

	"github.com/beevk/go-todo/domain"
	"github.com/beevk/go-todo/handlers"
	"github.com/beevk/go-todo/postgres"
	"github.com/go-pg/pg/v10"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUser := os.Getenv("DB_USERNAME")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "mySecretPass"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "todo"
	}

	DB := postgres.New(&pg.Options{
		Addr:     dbHost + ":5432",
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	defer func(DB *pg.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Error closing DB connection: %v", err)
		}
	}(DB)

	domainDB := &domain.DB{UserRepo: postgres.NewUserRepo(DB)}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Can not start the server %v", err)
	}
}
