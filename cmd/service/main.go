package main

import (
	"github.com/joho/godotenv"
	"go-project-layout/internal/durable"
	"go-project-layout/internal/model"
	"go-project-layout/internal/server"
	"log"
	"net/http"
	"os"
)

func init() {
	// setup logger
	durable.SetupLogger()

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := durable.ConnectDB(os.Getenv("DB_DSN")); err != nil {
		log.Fatal("Error connecting to database")
	}

	// migrate database
	if err := durable.Connection().AutoMigrate(
		&model.Users{}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	server.SetupRoutes(mux)

	middlewareMux := server.SetupMiddleware(mux)
	server.StartServer(middlewareMux, os.Getenv("SERVER_PORT"))
}
