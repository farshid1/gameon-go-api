package gameon

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"ledape.com/gameon/ent"
	"ledape.com/gameon/ent/user"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type Claims struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID uint   `json:"userId"`
	jwt.StandardClaims
}

func validateAndGetUserID(reqToken string, w http.ResponseWriter) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		reqToken,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(GetEnvWithKey("AUTH_SECRET")), nil
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 0, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return 0, err
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return 0, fmt.Errorf("Invalid Token")
	}

	return claims.UserID, nil
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(ctx context.Context, client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			if reqToken == "" {
				next.ServeHTTP(w, r)
				return
			}
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			if reqToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			userId, err := validateAndGetUserID(reqToken, w)
			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}
			// get the user from the database
			user, err := client.User.Query().Where(user.IDEQ(int(userId))).Only(ctx)
			if err != nil {
				http.Error(w, "Invalid User ID", http.StatusForbidden)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			newRequest := r.Clone(ctx)

			// and call the next with our new context
			next.ServeHTTP(w, newRequest)
		})
	}
}

func ForContext(ctx context.Context) *ent.User {
	user, _ := ctx.Value(userCtxKey).(*ent.User)
	return user
}

func CreateAuthTokenAndReturnAuthPayload(user *ent.User) (*AuthPayload, error) {
	claims := &Claims{
		Name:   user.Name,
		Email:  user.Email,
		UserID: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(50 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	secret := []byte(GetEnvWithKey("AUTH_SECRET"))

	tokenString, signingErr :=
		token.SignedString(secret)
	if signingErr != nil {
		return nil, signingErr
	}

	return &AuthPayload{
		Token: tokenString,
		User:  user,
	}, nil
}
