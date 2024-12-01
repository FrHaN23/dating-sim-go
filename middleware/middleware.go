package middleware

import (
	"context"
	"dating-sim/dotenv"
	"dating-sim/res"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type key string

const TokenKey key = "decodedToken"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			log.Print("empty Auth Header")
			res.ResErrJson(
				w,
				http.StatusForbidden,
				errors.New("forbidden"),
			)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == "" {
			res.ResErrJson(
				w,
				http.StatusForbidden,
				errors.New("forbidden"),
			)
			return
		}

		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(dotenv.Env("JWT_SECRET")), nil
			}
			return nil, errors.New("unauthorized")
		})

		if err != nil || !parsedToken.Valid {
			log.Print(err)
			res.ResErrJson(
				w,
				http.StatusUnauthorized,
				errors.New("Unauthorized"),
			)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			res.ResErrJson(
				w,
				http.StatusUnauthorized,
				errors.New("Unauthorized"),
			)
			return
		}
		// Pass the decoded token into the context
		ctx := context.WithValue(r.Context(), TokenKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
