package gameon

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"ledape.com/gameon/ent"
	"ledape.com/gameon/ent/user"
)

func (r *gameResolver) CreatedBy(ctx context.Context, obj *ent.Game) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *gameResolver) Time(ctx context.Context, obj *ent.Game) (string, error) {
	return obj.CreateTime.Format(time.RFC3339), nil
}

func (r *mutationResolver) Login(ctx context.Context, loginInput *LoginInput) (*AuthPayload, error) {
	user, err := r.client.User.
		Query().
		Where(user.EmailEQ(loginInput.Email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("User not found - failed querying user: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		return nil, fmt.Errorf("Invalid password provided")
	}

	return CreateAuthTokenAndReturnAuthPayload(user)
}

func (r *mutationResolver) Signup(ctx context.Context, signupInput *SignupInput) (*AuthPayload, error) {
	existingUser, err := r.client.User.
		Query().
		Where(user.EmailEQ(signupInput.Email)).
		Only(ctx)
	if existingUser != nil {
		return nil, fmt.Errorf("User Already exist - failed querying user: %w", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupInput.Password), 10)
	if err != nil {
		panic("failed to hash the password")
	}
	user := r.client.User.Create().
		SetEmail(signupInput.Email).
		SetName(signupInput.Name).
		SetPassword(string(hashedPassword)).
		SaveX(ctx)

	return CreateAuthTokenAndReturnAuthPayload(user)
}

func (r *mutationResolver) CreateGame(ctx context.Context, gameInput *GameInput) (*ent.Game, error) {
	user := ForContext(ctx)
	startTime, err := time.Parse("2006-01-02T15:04:05.000Z", gameInput.Time)
	if err != nil {
		return nil, err
	}
	game := r.client.Game.Create().
		SetTime(startTime).
		SetCreator(user).
		SetTitle(gameInput.Title).
		SaveX(ctx)
	return game, nil
}

func (r *queryResolver) UpcomingGames(ctx context.Context) ([]*ent.Game, error) {
	return r.client.User.Query().QueryParticipatingGames().AllX(ctx), nil
}

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type gameResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
