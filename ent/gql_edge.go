// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (ga *Game) Creator(ctx context.Context) (*User, error) {
	result, err := ga.Edges.CreatorOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryCreator().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ga *Game) Participants(ctx context.Context) ([]*User, error) {
	result, err := ga.Edges.ParticipantsOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryParticipants().All(ctx)
	}
	return result, err
}

func (u *User) CreatedGames(ctx context.Context) ([]*Game, error) {
	result, err := u.Edges.CreatedGamesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryCreatedGames().All(ctx)
	}
	return result, err
}

func (u *User) ParticipatingGames(ctx context.Context) ([]*Game, error) {
	result, err := u.Edges.ParticipatingGamesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryParticipatingGames().All(ctx)
	}
	return result, err
}