package routes

import (
	"dating-sim/handler"
	"dating-sim/middleware"

	"github.com/gorilla/mux"
)

// SetupUpgradeRoutes sets up routes for upgrading (e.g., to premium).
func SetupUpgradeRoutes(mux *mux.Router, handler *handler.Upgrade) {
	prefix := mux.PathPrefix("/api/v1/upgrade").Subrouter()
	prefix.Use(middleware.Middleware)
	{
		prefix.HandleFunc("/premium", handler.Premium)
	}
}
