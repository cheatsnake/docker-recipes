package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// get environment variables
	dbUser, _ := os.LookupEnv("POSTGRES_USER")
	dbPass, _ := os.LookupEnv("POSTGRES_PASSWORD")
	dbHost, _ := os.LookupEnv("POSTGRES_HOST")
	dbPort, _ := os.LookupEnv("POSTGRES_PORT")
	dbName, _ := os.LookupEnv("POSTGRES_DB")

	// URL example: "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// connect to database
	dbpool, err := pgxpool.New(context.Background(), url)
	handleError(err)
	defer dbpool.Close()

	// insert an item
	_, err = dbpool.Query(context.Background(), "insert into items (name, price) values ('Apple', 10);")
	handleError(err)

	var name string
	var price int

	// retrive an item
	err = dbpool.QueryRow(context.Background(), "select name, price from items where id=$1", 1).Scan(&name, &price)
	handleError(err)

	fmt.Printf("ITEM FOUND: name - %s, price - %d$ \n", name, price)
	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n", err)
		os.Exit(1)
	}
}
