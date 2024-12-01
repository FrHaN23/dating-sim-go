package handler

import (
	"dating-sim/models"
	"dating-sim/res"
	"dating-sim/types"
	"dating-sim/utils"
	"errors"
	"log"
	"net/http"
	"time"
)

type Upgrade struct{}

// swipe handles right swipe means like
func (u *Upgrade) Premium(w http.ResponseWriter, r *http.Request) {
	u.handleRequest(w, r, u.PremiumHandler)
}

// handleRequest abstracts out the method checking and delegates to the handler.
func (s *Upgrade) handleRequest(w http.ResponseWriter, r *http.Request, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	switch r.Method {
	case http.MethodGet:
		handlerFunc(w, r)
	case http.MethodOptions:
		res.PreFlightJson(w, http.StatusNoContent, "GET, OPTIONS")
	default:
		res.PreFlightJson(w, http.StatusMethodNotAllowed, "GET, OPTIONS")
	}
}

func (u *Upgrade) PremiumHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to get user id"),
		)
		return
	}

	now := time.Now()
	user.ID = userID

	user.Premium = &now

	if err = user.Update(); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to update user"),
		)
		return
	}
	res.ResOkJSON(
		w,
		&types.Res{
			Message: "ok",
			Data:    nil,
		},
	)
}
