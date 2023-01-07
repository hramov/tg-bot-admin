package main

import (
	"github.com/hramov/tg-bot-admin/internal/composite"
	"github.com/hramov/tg-bot-admin/pkg/api/middleware"
	"github.com/hramov/tg-bot-admin/pkg/db"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := logging.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Warning("cannot load .env file: %v. if it occurred inside a container, don't worry", err)
	}

	pg, err := db.DatabaseFactory(db.Postgres)
	if err != nil {
		logger.Fatal("cannot start postgres: %v", err)
		os.Exit(1)
	}

	logger.Info("create router")
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(middleware.Cors)

	logger.Info("initializing user module")
	composite.NewUser(pg, router, logger)

	logger.Info("server started on port:", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}
