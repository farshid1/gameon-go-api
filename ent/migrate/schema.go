// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "time", Type: field.TypeTime},
		{Name: "user_created_games", Type: field.TypeInt, Nullable: true},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_users_createdGames",
				Columns:    []*schema.Column{GamesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GameParticipantsColumns holds the columns for the "game_participants" table.
	GameParticipantsColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime},
		{Name: "rsvp_status", Type: field.TypeEnum, Enums: []string{"YES", "NO", "MAYBE"}},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "game_id", Type: field.TypeInt},
	}
	// GameParticipantsTable holds the schema information for the "game_participants" table.
	GameParticipantsTable = &schema.Table{
		Name:       "game_participants",
		Columns:    GameParticipantsColumns,
		PrimaryKey: []*schema.Column{GameParticipantsColumns[2], GameParticipantsColumns[3]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_participants_users_user",
				Columns:    []*schema.Column{GameParticipantsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "game_participants_games_game",
				Columns:    []*schema.Column{GameParticipantsColumns[3]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GamesTable,
		GameParticipantsTable,
		UsersTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = UsersTable
	GameParticipantsTable.ForeignKeys[0].RefTable = UsersTable
	GameParticipantsTable.ForeignKeys[1].RefTable = GamesTable
}
