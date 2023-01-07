package main

import (
	"github.com/hramov/tg-bot-admin/internal/composite"
	"github.com/hramov/tg-bot-admin/pkg/api/middleware"
	"github.com/hramov/tg-bot-admin/pkg/client"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("cannot load .env file: %v. if it occured inside a container, don't worry", err)
	}

	pg, err := client.DatabaseFactory(client.Postgres)
	if err != nil {
		log.Printf("cannot start postgres: %v", err)
		os.Exit(3)
	}

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(middleware.Cors)

	composite.NewUser(pg, router)

	log.Println("server started on port:", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}
