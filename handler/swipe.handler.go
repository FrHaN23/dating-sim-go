package handler

import (
	"context"
	"dating-sim/models"
	"dating-sim/res"
	"dating-sim/types"
	"dating-sim/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Swipe struct {
	DB    *gorm.DB
	Redis *redis.Client
}

const dailyLimit = 10

var ctx = context.Background()

// SwipeData holds the daily swipe data for a user
type SwipeData struct {
	UserID         uint
	SwipedProfiles []uint
	SwipeCount     int
	LastSwipeDate  time.Time
}

// Should create 1 method swipped with added params action (like or pass)
// swipe handles right swipe means like
func (s *Swipe) Right(w http.ResponseWriter, r *http.Request) {
	s.handleRequest(w, r, s.right)
}

// swipe handles left swipe means pass.
func (s *Swipe) Left(w http.ResponseWriter, r *http.Request) {
	s.handleRequest(w, r, s.left)
}

// handleRequest abstracts out the method checking and delegates to the handler.
func (s *Swipe) handleRequest(w http.ResponseWriter, r *http.Request, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	switch r.Method {
	case http.MethodPost:
		handlerFunc(w, r)
	case http.MethodOptions:
		res.PreFlightJson(w, http.StatusNoContent, "POST, OPTIONS")
	default:
		res.PreFlightJson(w, http.StatusMethodNotAllowed, "POST, OPTIONS")
	}
}

// right handles right swipe (like) action.
func (s *Swipe) right(w http.ResponseWriter, r *http.Request) {
	var swiper models.Swipe

	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get user id"),
		)
	}
	swipeData, err := s.getSwipeDataFromRedis(userID)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get user swipe data from Redis"),
		)
		return
	}

	// Check if the user premium and has already swiped 10 profiles today
	if isPremium, _ := utils.GetUserPremiumStatus(r); !isPremium && swipeData.SwipeCount >= dailyLimit {
		res.ResErrJson(
			w,
			http.StatusBadRequest,
			errors.New("reached the daily swipe limit"),
		)
		return
	}

	// Get profile ID from request URL
	profileID, err := utils.GetProfileIDFromRequest(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get profile id"),
		)
	}

	// Check if this profile was already swiped today
	if contains(swipeData.SwipedProfiles, profileID) {
		res.ResErrJson(
			w,
			http.StatusConflict,
			errors.New("profile already swiped"),
		)
		return
	}

	// Process the swipe action (like)
	err = swiper.ProcessSwipe(uint(userID), uint(profileID), true)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to process swipe"),
		)
		return
	}

	// Update the swipe count and swiped profiles in Redis
	swipeData.SwipeCount++
	swipeData.SwipedProfiles = append(swipeData.SwipedProfiles, profileID)

	// Save the updated data in Redis with 24-hour expiration
	err = s.saveSwipeDataToRedis(swipeData)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to save swipe data to Redis"),
		)
		return
	}

	// Respond with a success message
	res.ResOkJSON(w, &types.Res{
		Message: fmt.Sprintf("Profile %d swiped right", profileID),
		Data:    nil,
	})
}

// left handles left swipe (pass) action.
func (s *Swipe) left(w http.ResponseWriter, r *http.Request) {
	var swiper models.Swipe

	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get user id"),
		)
	}

	// Check the user's daily swipe count and swiped profiles from Redis
	swipeData, err := s.getSwipeDataFromRedis(userID)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get user swipe data from Redis"),
		)
		return
	}

	if isPremium, _ := utils.GetUserPremiumStatus(r); !isPremium && swipeData.SwipeCount >= dailyLimit {
		res.ResErrJson(
			w,
			http.StatusBadRequest,
			errors.New("reached the daily swipe limit"),
		)
		return
	}

	profileID, err := utils.GetProfileIDFromRequest(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get profile id"),
		)
	}

	if contains(swipeData.SwipedProfiles, profileID) {
		res.ResErrJson(
			w,
			http.StatusConflict,
			errors.New("profile already swiped"),
		)
		return
	}

	// Process the swipe action (pass)
	err = swiper.ProcessSwipe(uint(userID), uint(profileID), false)

	if err != nil {
		http.Error(w, "Failed to process left swipe", http.StatusInternalServerError)
		return
	}

	// Update the swipe count and swiped profiles in Redis
	swipeData.SwipeCount++
	swipeData.SwipedProfiles = append(swipeData.SwipedProfiles, profileID)

	err = s.saveSwipeDataToRedis(swipeData)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to save swipe data to Redis"),
		)
		return
	}

	res.ResOkJSON(w, &types.Res{
		Message: fmt.Sprintf("Profile %d swiped left", profileID),
	})
}

// getSwipeDataFromRedis retrieves the swipe data from Redis for the user
func (s *Swipe) getSwipeDataFromRedis(userID uint) (*SwipeData, error) {
	key := fmt.Sprintf("swipeData:%d", userID)
	data, err := s.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return &SwipeData{
			UserID:         userID,
			SwipedProfiles: []uint{},
			SwipeCount:     0,
			LastSwipeDate:  time.Now(),
		}, nil
	}
	if err != nil {
		return nil, err
	}

	var swipeData SwipeData
	err = json.Unmarshal([]byte(data), &swipeData)
	if err != nil {
		return nil, err
	}

	return &swipeData, nil
}

// saveSwipeDataToRedis saves the swipe data to Redis with an expiration time
func (s *Swipe) saveSwipeDataToRedis(swipeData *SwipeData) error {
	key := fmt.Sprintf("swipeData:%d", swipeData.UserID)
	data, err := json.Marshal(swipeData)
	if err != nil {
		return err
	}

	// Save the data to Redis with a 24-hour expiration
	return s.Redis.Set(ctx, key, data, 24*time.Hour).Err()
}

// contains checks if a profile ID is already in the list of swiped profiles
func contains(profiles []uint, profileID uint) bool {
	for _, id := range profiles {
		if uint(id) == uint(profileID) {
			return true
		}
	}
	return false
}
