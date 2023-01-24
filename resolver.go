package gameon

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis/v9"
	"ledape.com/gameon/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client      *ent.Client
	redisClient *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, redisClient *redis.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client, redisClient},
	})
}
