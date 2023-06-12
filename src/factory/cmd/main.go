package main

import (
	"github.com/hramov/tbotfactory/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("cannot load env: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("cannot get port from env: %v", err)
	}

	s := server.New(port)
	if err = s.Start(); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
