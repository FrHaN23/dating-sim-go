package routes

import (
	"dating-sim/handler"
	"dating-sim/middleware"

	"github.com/gorilla/mux"
)

func SetupProfileRoutes(mux *mux.Router, handler *handler.Profile) {
	prefix := mux.PathPrefix("/api/v1/profile").Subrouter()
	prefix.Use(middleware.Middleware)
	{
		prefix.HandleFunc("/me", handler.Me)
		prefix.HandleFunc("/password", handler.ChangePassword)
	}
}
