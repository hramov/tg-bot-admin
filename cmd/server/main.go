package main

import (
	"github.com/hramov/tg-bot-admin/pkg/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("cannot load .env file: %v. if it occured inside a container, don't worry", err)
	}

	_, err := client.DatabaseFactory(client.Postgres)
	if err != nil {
		log.Printf("cannot start postgres: %v", err)
		os.Exit(3)
	}
}
