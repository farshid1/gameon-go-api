// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"ledape.com/gameon/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// CreatedGames holds the value of the createdGames edge.
	CreatedGames []*Game `json:"createdGames,omitempty"`
	// ParticipatingGames holds the value of the participating_games edge.
	ParticipatingGames []*Game `json:"participating_games,omitempty"`
	// Participants holds the value of the participants edge.
	Participants []*GameParticipant `json:"participants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedCreatedGames       map[string][]*Game
	namedParticipatingGames map[string][]*Game
	namedParticipants       map[string][]*GameParticipant
}

// CreatedGamesOrErr returns the CreatedGames value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedGamesOrErr() ([]*Game, error) {
	if e.loadedTypes[0] {
		return e.CreatedGames, nil
	}
	return nil, &NotLoadedError{edge: "createdGames"}
}

// ParticipatingGamesOrErr returns the ParticipatingGames value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ParticipatingGamesOrErr() ([]*Game, error) {
	if e.loadedTypes[1] {
		return e.ParticipatingGames, nil
	}
	return nil, &NotLoadedError{edge: "participating_games"}
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ParticipantsOrErr() ([]*GameParticipant, error) {
	if e.loadedTypes[2] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldEmail, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		}
	}
	return nil
}

// QueryCreatedGames queries the "createdGames" edge of the User entity.
func (u *User) QueryCreatedGames() *GameQuery {
	return (&UserClient{config: u.config}).QueryCreatedGames(u)
}

// QueryParticipatingGames queries the "participating_games" edge of the User entity.
func (u *User) QueryParticipatingGames() *GameQuery {
	return (&UserClient{config: u.config}).QueryParticipatingGames(u)
}

// QueryParticipants queries the "participants" edge of the User entity.
func (u *User) QueryParticipants() *GameParticipantQuery {
	return (&UserClient{config: u.config}).QueryParticipants(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCreatedGames returns the CreatedGames named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCreatedGames(name string) ([]*Game, error) {
	if u.Edges.namedCreatedGames == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCreatedGames[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCreatedGames(name string, edges ...*Game) {
	if u.Edges.namedCreatedGames == nil {
		u.Edges.namedCreatedGames = make(map[string][]*Game)
	}
	if len(edges) == 0 {
		u.Edges.namedCreatedGames[name] = []*Game{}
	} else {
		u.Edges.namedCreatedGames[name] = append(u.Edges.namedCreatedGames[name], edges...)
	}
}

// NamedParticipatingGames returns the ParticipatingGames named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedParticipatingGames(name string) ([]*Game, error) {
	if u.Edges.namedParticipatingGames == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedParticipatingGames[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedParticipatingGames(name string, edges ...*Game) {
	if u.Edges.namedParticipatingGames == nil {
		u.Edges.namedParticipatingGames = make(map[string][]*Game)
	}
	if len(edges) == 0 {
		u.Edges.namedParticipatingGames[name] = []*Game{}
	} else {
		u.Edges.namedParticipatingGames[name] = append(u.Edges.namedParticipatingGames[name], edges...)
	}
}

// NamedParticipants returns the Participants named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedParticipants(name string) ([]*GameParticipant, error) {
	if u.Edges.namedParticipants == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedParticipants[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedParticipants(name string, edges ...*GameParticipant) {
	if u.Edges.namedParticipants == nil {
		u.Edges.namedParticipants = make(map[string][]*GameParticipant)
	}
	if len(edges) == 0 {
		u.Edges.namedParticipants[name] = []*GameParticipant{}
	} else {
		u.Edges.namedParticipants[name] = append(u.Edges.namedParticipants[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
