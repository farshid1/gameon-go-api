package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

func (Game) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Time("time"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("createdGames").
			Unique(),
		edge.To("participants", User.Type),
	}
}
