package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"ledape.com/gameon"
	"ledape.com/gameon/ent"
	"ledape.com/gameon/ent/migrate"
)

const defaultPort = "8080"

func main() {
	gameon.LoadEnv()
	get := gameon.GetEnvWithKey

	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			// "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			get("DB_USER"),
			get("DB_PASSWORD"),
			get("DB_HOST"),
			get("DB_PORT"),
			get("DB_NAME"),
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	port := get("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(gameon.Middleware(ctx, client))

	srv := handler.NewDefaultServer(gameon.NewSchema(client))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
