package gameon

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
	AccessUUID  string
	RefreshUUID string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
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
		if err == jwt.ErrTokenExpired {
			// TODO: handle the expired token case
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

// Middleware decodes the auth token and packs the session into context
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

func storeAuthToken(ctx context.Context, redisClient *redis.Client, userId int, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := redisClient.Set(ctx, td.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := redisClient.Set(ctx, td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func RecreateAccessToken(ctx context.Context, dbClient *ent.Client, redisClient *redis.Client, token string) (*AuthPayload, error) {
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("Invalid Refresh token sent")
	}
	refreshToken := splitToken[1]
	claims := &Claims{}

	refToken, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(GetEnvWithKey("AUTH_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !refToken.Valid {
		return nil, errors.New("Refresh token is no longer valid")
	}
	user, err := dbClient.User.
		Query().
		Where(user.IDEQ(int(claims.UserID))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	td, err := createToken(user)
	if err != nil {
		return nil, err
	}
	storeAuthToken(ctx, redisClient, user.ID, td)

	return &AuthPayload{
		RefreshToken: td.RefreshToken,
		Token:        td.AccessToken,
		User:         user,
	}, nil
}

func createToken(user *ent.User) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.New().String()
	td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RefreshUuid = uuid.New().String()
	claims := &Claims{
		Name:   user.Name,
		Email:  user.Email,
		UserID: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		AccessUUID: td.AccessUuid,
	}
	refreshClaims := &Claims{
		UserID: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		RefreshUUID: td.RefreshUuid,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		refreshClaims,
	)
	secret := []byte(GetEnvWithKey("AUTH_SECRET"))

	tokenString, signingErr :=
		token.SignedString(secret)
	if signingErr != nil {
		return nil, signingErr
	}
	td.AccessToken = tokenString

	refreshTokenString, signingErr :=
		refreshToken.SignedString(secret)
	if signingErr != nil {
		return nil, signingErr
	}
	td.RefreshToken = refreshTokenString

	return td, nil
}

func CreateAuthTokenAndReturnAuthPayload(ctx context.Context, redisClient *redis.Client, user *ent.User) (*AuthPayload, error) {
	td, err := createToken(user)
	if err != nil {
		return nil, err
	}
	storeAuthToken(ctx, redisClient, user.ID, td)

	return &AuthPayload{
		RefreshToken: td.RefreshToken,
		Token:        td.AccessToken,
		User:         user,
	}, nil
}
