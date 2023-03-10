// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"ledape.com/gameon/ent/game"
	"ledape.com/gameon/ent/user"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Time holds the value of the "time" field.
	Time time.Time `json:"time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges              GameEdges `json:"edges"`
	user_created_games *int
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// GameParticipants holds the value of the game_participants edge.
	GameParticipants []*User `json:"game_participants,omitempty"`
	// Participants holds the value of the participants edge.
	Participants []*GameParticipant `json:"participants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedGameParticipants map[string][]*User
	namedParticipants     map[string][]*GameParticipant
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// GameParticipantsOrErr returns the GameParticipants value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) GameParticipantsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.GameParticipants, nil
	}
	return nil, &NotLoadedError{edge: "game_participants"}
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) ParticipantsOrErr() ([]*GameParticipant, error) {
	if e.loadedTypes[2] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			values[i] = new(sql.NullInt64)
		case game.FieldTitle:
			values[i] = new(sql.NullString)
		case game.FieldCreateTime, game.FieldUpdateTime, game.FieldTime:
			values[i] = new(sql.NullTime)
		case game.ForeignKeys[0]: // user_created_games
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Game", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case game.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ga.CreateTime = value.Time
			}
		case game.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ga.UpdateTime = value.Time
			}
		case game.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ga.Title = value.String
			}
		case game.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				ga.Time = value.Time
			}
		case game.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_created_games", value)
			} else if value.Valid {
				ga.user_created_games = new(int)
				*ga.user_created_games = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCreator queries the "creator" edge of the Game entity.
func (ga *Game) QueryCreator() *UserQuery {
	return (&GameClient{config: ga.config}).QueryCreator(ga)
}

// QueryGameParticipants queries the "game_participants" edge of the Game entity.
func (ga *Game) QueryGameParticipants() *UserQuery {
	return (&GameClient{config: ga.config}).QueryGameParticipants(ga)
}

// QueryParticipants queries the "participants" edge of the Game entity.
func (ga *Game) QueryParticipants() *GameParticipantQuery {
	return (&GameClient{config: ga.config}).QueryParticipants(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return (&GameClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ga.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ga.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(ga.Title)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(ga.Time.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedGameParticipants returns the GameParticipants named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ga *Game) NamedGameParticipants(name string) ([]*User, error) {
	if ga.Edges.namedGameParticipants == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ga.Edges.namedGameParticipants[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ga *Game) appendNamedGameParticipants(name string, edges ...*User) {
	if ga.Edges.namedGameParticipants == nil {
		ga.Edges.namedGameParticipants = make(map[string][]*User)
	}
	if len(edges) == 0 {
		ga.Edges.namedGameParticipants[name] = []*User{}
	} else {
		ga.Edges.namedGameParticipants[name] = append(ga.Edges.namedGameParticipants[name], edges...)
	}
}

// NamedParticipants returns the Participants named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ga *Game) NamedParticipants(name string) ([]*GameParticipant, error) {
	if ga.Edges.namedParticipants == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ga.Edges.namedParticipants[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ga *Game) appendNamedParticipants(name string, edges ...*GameParticipant) {
	if ga.Edges.namedParticipants == nil {
		ga.Edges.namedParticipants = make(map[string][]*GameParticipant)
	}
	if len(edges) == 0 {
		ga.Edges.namedParticipants[name] = []*GameParticipant{}
	} else {
		ga.Edges.namedParticipants[name] = append(ga.Edges.namedParticipants[name], edges...)
	}
}

// Games is a parsable slice of Game.
type Games []*Game

func (ga Games) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
