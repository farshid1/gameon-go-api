package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"ledape.com/gameon/ent"
	"ledape.com/gameon/utils"
)

func main() {
	utils.LoadEnv()
	get := utils.GetEnvWithKey

	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			get("DB_HOST"),
			get("DB_PORT"),
			get("DB_USER"),
			get("DB_NAME"),
			get("DB_PASSWORD"),
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
