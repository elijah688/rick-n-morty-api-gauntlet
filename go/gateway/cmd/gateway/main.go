package main

import (
	"log"
	"riki_gateway/internal/api"
	"riki_gateway/internal/config"
	"riki_gateway/internal/services"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := api.New(services.New(cfg)).Start(); err != nil {
		log.Fatal(err)
	}
}
