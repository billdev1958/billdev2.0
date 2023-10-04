package main

import (
	"log"
	"os"

	"github.com/billdev1958/billdev2.0.git/api"
	"github.com/billdev1958/billdev2.0.git/db"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	store, err := db.NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(os.Getenv("PORT"), store)
	server.Run()
}
