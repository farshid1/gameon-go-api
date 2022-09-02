package gameon

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"ledape.com/gameon/ent"
	"ledape.com/gameon/ent/user"
)

func (r *gameResolver) CreatedBy(ctx context.Context, obj *ent.Game) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *gameResolver) Time(ctx context.Context, obj *ent.Game) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, loginInput *LoginInput) (*AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Signup(ctx context.Context, signupInput *SignupInput) (*AuthPayload, error) {
	_, err := r.client.User.
		Query().
		Where(user.EmailEQ(signupInput.Email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupInput.Password), 10)
	if err != nil {
		panic("failed to hash the password")
	}
	user := r.client.User.Create().
		SetEmail(signupInput.Name).
		SetName(signupInput.Name).
		SetPassword(string(hashedPassword)).
		SaveX(ctx)
	return CreateAuthTokenAndReturnAuthPayload(user)
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
