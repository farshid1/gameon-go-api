// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"ledape.com/gameon/ent/game"
	"ledape.com/gameon/ent/gameparticipant"
	"ledape.com/gameon/ent/user"
)

// GameParticipant is the model entity for the GameParticipant schema.
type GameParticipant struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// RsvpStatus holds the value of the "rsvp_status" field.
	RsvpStatus gameparticipant.RsvpStatus `json:"rsvp_status,omitempty"`
	// GameID holds the value of the "game_id" field.
	GameID int `json:"game_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameParticipantQuery when eager-loading is set.
	Edges GameParticipantEdges `json:"edges"`
}

// GameParticipantEdges holds the relations/edges for other nodes in the graph.
type GameParticipantEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameParticipantEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameParticipantEdges) GameOrErr() (*Game, error) {
	if e.loadedTypes[1] {
		if e.Game == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameParticipant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameparticipant.FieldGameID, gameparticipant.FieldUserID:
			values[i] = new(sql.NullInt64)
		case gameparticipant.FieldRsvpStatus:
			values[i] = new(sql.NullString)
		case gameparticipant.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GameParticipant", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameParticipant fields.
func (gp *GameParticipant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameparticipant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gp.CreatedAt = value.Time
			}
		case gameparticipant.FieldRsvpStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rsvp_status", values[i])
			} else if value.Valid {
				gp.RsvpStatus = gameparticipant.RsvpStatus(value.String)
			}
		case gameparticipant.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				gp.GameID = int(value.Int64)
			}
		case gameparticipant.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gp.UserID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the GameParticipant entity.
func (gp *GameParticipant) QueryUser() *UserQuery {
	return (&GameParticipantClient{config: gp.config}).QueryUser(gp)
}

// QueryGame queries the "game" edge of the GameParticipant entity.
func (gp *GameParticipant) QueryGame() *GameQuery {
	return (&GameParticipantClient{config: gp.config}).QueryGame(gp)
}

// Update returns a builder for updating this GameParticipant.
// Note that you need to call GameParticipant.Unwrap() before calling this method if this GameParticipant
// was returned from a transaction, and the transaction was committed or rolled back.
func (gp *GameParticipant) Update() *GameParticipantUpdateOne {
	return (&GameParticipantClient{config: gp.config}).UpdateOne(gp)
}

// Unwrap unwraps the GameParticipant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gp *GameParticipant) Unwrap() *GameParticipant {
	_tx, ok := gp.config.driver.(*txDriver)
	if !ok {
		panic("ent: GameParticipant is not a transactional entity")
	}
	gp.config.driver = _tx.drv
	return gp
}

// String implements the fmt.Stringer.
func (gp *GameParticipant) String() string {
	var builder strings.Builder
	builder.WriteString("GameParticipant(")
	builder.WriteString("created_at=")
	builder.WriteString(gp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("rsvp_status=")
	builder.WriteString(fmt.Sprintf("%v", gp.RsvpStatus))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.GameID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// GameParticipants is a parsable slice of GameParticipant.
type GameParticipants []*GameParticipant

func (gp GameParticipants) config(cfg config) {
	for _i := range gp {
		gp[_i].config = cfg
	}
}