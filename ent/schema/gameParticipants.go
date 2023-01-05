package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GameParticipant holds the edge schema definition of the GameParticipant relationship.
type GameParticipant struct {
	ent.Schema
}

func (GameParticipant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "game_id"),
	}
}

// Fields of the GameParticipant.
func (GameParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Enum("rsvp_status").NamedValues(
			"YES", "YES",
			"NO", "NO",
			"MAYBE", "MAYBE",
		),
		field.Int("game_id"),
		field.Int("user_id"),
	}
}

// Edges of the GameParticipant.
func (GameParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("game", Game.Type).
			Required().
			Unique().
			Field("game_id"),
	}
}
