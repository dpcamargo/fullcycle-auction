package main

import (
	"context"
	"log"

	"github.com/dpcamargo/fullcycle-auction/configuration/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error trying to load env variables")
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error trying to connect to database")
	}
	
}
