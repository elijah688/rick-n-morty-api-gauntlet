package main

import (
	"context"
	"log"
	"riki/internal/api"
	"riki/internal/config"
	"riki/internal/db"
	"riki/internal/services"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = make(map[int]Item)
var currentID = 1

func main() {

	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}

	db, err := db.New(ctx, cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	if err := api.New(services.New(db)).Start(); err != nil {
		log.Fatal(err)
	}
}
