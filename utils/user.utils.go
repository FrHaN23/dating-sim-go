package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// getProfileIDFromRequest extracts the profile ID from the URL path (e.g., /swipe/{id}/right).
func GetProfileIDFromRequest(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	profileIDStr, exists := vars["id"]
	if !exists {
		return 0, fmt.Errorf("profile ID not found in URL")
	}

	profileID, err := strconv.Atoi(profileIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid profile ID format")
	}

	return uint(profileID), nil
}

func GetUserIDFromContext(r *http.Request) (uint, error) {
	claims, err := ExtractToken(r)
	if err != nil {
		return 0, fmt.Errorf("failed extarct token")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user ID is invalid in token claims")
	}
	return uint(userID), nil
}

func GetUserPremiumStatus(r *http.Request) (bool, error) {
	claims, err := ExtractToken(r)
	if err != nil {
		return false, fmt.Errorf("failed extarct token")
	}
	premium, ok := claims["premium"].(string)
	if !ok {
		return false, fmt.Errorf("user ID is invalid in token claims")
	}
	return premium != "", nil
}
