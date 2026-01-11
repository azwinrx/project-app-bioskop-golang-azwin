package main

import (
	"log"

	"project-app-bioskop-golang-azwin/pkg/database"
	"project-app-bioskop-golang-azwin/pkg/utils"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to the database")
}
