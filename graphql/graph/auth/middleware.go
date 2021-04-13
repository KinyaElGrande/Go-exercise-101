package auth

import (
	"context"
	"net/http"

	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/jwt"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/model"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/repository/userRepository"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//Validates jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "invalid token", http.StatusForbidden)
			}

			user := model.User{Username: username}
			id, err := userRepository.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id

			//put it into context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
