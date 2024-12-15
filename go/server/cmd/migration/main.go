package main

import (
	"context"
	"log"
	"riki/internal/config"
	"riki/internal/db"
	"riki/internal/migrator"
	"riki/internal/riki_client"
)

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

	client := riki_client.NewRikiClient(cfg.RikiClientConfig)

	if err := migrator.New(client, db).Run(ctx); err != nil {
		log.Fatal(err)
	}

}
