package utils

import (
	"dating-sim/middleware"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(r *http.Request) (jwt.MapClaims, error) {
	claims, ok := r.Context().Value(middleware.TokenKey).(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token not found in context")
	}
	return claims, nil
}
