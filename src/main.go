package main

import (
	"github.com/hramov/tg-bot-admin/src/domain/auth"
	"github.com/hramov/tg-bot-admin/src/domain/user"
	"github.com/hramov/tg-bot-admin/src/interface/api"
	"github.com/hramov/tg-bot-admin/src/interface/registry"
	domainContainer "github.com/hramov/tg-bot-admin/src/modules/container"
	"github.com/hramov/tg-bot-admin/src/modules/context"
	"github.com/hramov/tg-bot-admin/src/modules/data_source"
	"github.com/hramov/tg-bot-admin/src/modules/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("cannot load .env file: %v. if it occured inside a container, don't worry", err)
	}

	if err := logger.New(); err != nil {
		log.Printf("cannot start logger: %v", err)
		os.Exit(2)
	}

	customContext.New()
	domainContainer.New()

	pg, err := data_source.DatabaseFactory(data_source.Postgres)
	if err != nil {
		log.Printf("cannot start postgres: %v", err)
		os.Exit(3)
	}

	// initialize modules
	auth.New(registry.NewAuthRegistry(pg))
	user.New(registry.NewUserRegistry(pg))

	server := api.StartServer()
	if err := server.Run(os.Getenv("SERVER_PORT")); err != nil {
		log.Printf("cannot start server: %v", err)
		os.Exit(4)
	}

}
