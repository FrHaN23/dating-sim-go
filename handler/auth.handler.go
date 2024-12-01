package handler

import (
	"dating-sim/dotenv"
	"dating-sim/models"
	"dating-sim/res"
	"dating-sim/types"
	"dating-sim/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Auth contains authentication-related handlers.
type Auth struct{}

// Signup handles user registration.
func (a *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	a.handleRequest(w, r, a.signupHandler)
}

// Login handles user login.
func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	a.handleRequest(w, r, a.loginUser)
}

// handleRequest abstracts out the method checking and delegates to the handler.
func (a *Auth) handleRequest(w http.ResponseWriter, r *http.Request, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	switch r.Method {
	case http.MethodPost:
		handlerFunc(w, r)
	case http.MethodOptions:
		res.PreFlightJson(w, http.StatusNoContent, "POST, OPTIONS")
	default:
		res.PreFlightJson(w, http.StatusMethodNotAllowed, "POST, OPTIONS")
	}
}

func (a *Auth) signupHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB body req
	if err := utils.DecodeJson(&req, w, r.Body, false); err != nil {
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Email == "" || req.Password == "" {
		res.ResErrJson(
			w,
			http.StatusBadRequest,
			errors.New("missing requried fields"),
		)
		return
	}

	curPas := req.Password
	req.Password, _ = utils.HashPassword(curPas)

	// create user
	if err := req.Create(); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to created user"),
		)
		return
	}

	// Respond with success
	res.ResOkJSON(w, &types.Res{
		Message: "ok",
		Data:    "user created",
	})
}

func (a *Auth) loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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

	user.Username = req.Username

	if err := user.GetUserByUsername(); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed to find user"),
		)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusUnauthorized,
			errors.New("username/password wrong"),
		)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"premium":  user.Premium,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(dotenv.Env("JWT_SECRET")))

	if err != nil {
		log.Print(err)
		res.ResErrJson(
			w,
			http.StatusInternalServerError,
			errors.New("failed signed token"),
		)
		return
	}

	res.ResOkJSON(w, &types.Res{
		Message: "ok",
		Data: map[string]string{
			"token": tokenString,
			"type":  "bearer",
		},
	})
}
