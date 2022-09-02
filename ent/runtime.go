// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"ledape.com/gameon/ent/game"
	"ledape.com/gameon/ent/schema"
	"ledape.com/gameon/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	gameMixin := schema.Game{}.Mixin()
	gameMixinFields0 := gameMixin[0].Fields()
	_ = gameMixinFields0
	gameFields := schema.Game{}.Fields()
	_ = gameFields
	// gameDescCreateTime is the schema descriptor for create_time field.
	gameDescCreateTime := gameMixinFields0[0].Descriptor()
	// game.DefaultCreateTime holds the default value on creation for the create_time field.
	game.DefaultCreateTime = gameDescCreateTime.Default.(func() time.Time)
	// gameDescUpdateTime is the schema descriptor for update_time field.
	gameDescUpdateTime := gameMixinFields0[1].Descriptor()
	// game.DefaultUpdateTime holds the default value on creation for the update_time field.
	game.DefaultUpdateTime = gameDescUpdateTime.Default.(func() time.Time)
	// game.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	game.UpdateDefaultUpdateTime = gameDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
}
