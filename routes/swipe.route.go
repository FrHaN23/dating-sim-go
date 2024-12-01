package routes

import (
	"dating-sim/handler"
	"dating-sim/middleware"

	"github.com/gorilla/mux"
)

// SetupSwipeRoutes sets up routes for swipe actions (e.g., swipe right, swipe left).
func SetupSwipeRoutes(mux *mux.Router, handler *handler.Swipe) {
	prefix := mux.PathPrefix("/api/v1/swipe").Subrouter()
	prefix.Use(middleware.Middleware)
	{
		prefix.HandleFunc("/{id:[0-9]+}/right", handler.Right)
		prefix.HandleFunc("/{id:[0-9]+}/left", handler.Left)
	}
}
