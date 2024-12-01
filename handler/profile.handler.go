package handler

import (
	"dating-sim/models"
	"dating-sim/res"
	"dating-sim/types"
	"dating-sim/utils"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Profile struct{}

// Me handles the retrieval of the user's profile
func (p *Profile) Me(w http.ResponseWriter, r *http.Request) {
	p.handleRequest(w, r, p.meHandler, http.MethodGet)
}

// UpdateProfile handles updating the user's profile
func (p *Profile) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	p.handleRequest(w, r, p.updateProfileHandler, http.MethodPut)
}

// changePassword handles updating the user's password
func (p *Profile) ChangePassword(w http.ResponseWriter, r *http.Request) {
	p.handleRequest(w, r, p.changePasswordHandler, http.MethodPost)
}

func (p *Profile) handleRequest(w http.ResponseWriter, r *http.Request, handlerFunc func(w http.ResponseWriter, r *http.Request), allowedMethod string) {
	switch r.Method {
	case allowedMethod:
		handlerFunc(w, r)
	case http.MethodOptions:
		res.PreFlightJson(w, http.StatusNoContent, allowedMethod+", OPTIONS")
	default:
		res.PreFlightJson(w, http.StatusMethodNotAllowed, allowedMethod+", OPTIONS")
	}
}

func (p *Profile) meHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to find user"),
		)
		return
	}
	user.ID = userID
	user.Get()
	res.ResOkJSON(
		w,
		&types.Res{
			Message: "ok",
			Data:    user,
		},
	)
}

func (p *Profile) updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB body req
	if err := utils.DecodeJson(&user, w, r.Body, false); err != nil {
		res.ResErrJson(
			w,
			http.StatusRequestedRangeNotSatisfiable,
			errors.New("failed to parse json"),
		)
		return
	}
	defer r.Body.Close()

	if err := user.Update(); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to update user"),
		)
		return
	}
}

func (p *Profile) changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to find user"),
		)
		return
	}
	user.ID = userID

	user.GetWithPassword()

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB body req
	if err := utils.DecodeJson(&req, w, r.Body, false); err != nil {
		res.ResErrJson(
			w,
			http.StatusRequestedRangeNotSatisfiable,
			errors.New("failed to parse json"),
		)
		return
	}
	defer r.Body.Close()

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusUnauthorized,
			errors.New("username/password wrong"),
		)
		return
	}

	user.Password, _ = utils.HashPassword(req.NewPassword)

	if err := user.Update(); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to update user"),
		)
		return
	}

	res.ResOkJSON(w, &types.Res{
		Message: "ok",
		Data:    nil,
	})
}
