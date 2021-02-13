package database

import (
	"context"
	"log"
	"web-stash-api/config"
	"web-stash-api/ent"
	"web-stash-api/ent/migrate"
)

func OpenDb() *ent.Client {
	client, err := ent.Open(config.Cfg.Database.Driver, config.Cfg.Database.Address)
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	return client
}

func Migrate() {
	client := OpenDb()
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		log.Fatalf("failed to migrate schema resources: %v", err)
	}
}
