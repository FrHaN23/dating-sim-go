package routes

import (
	"dating-sim/db"
	"dating-sim/handler"

	"dating-sim/redis"

	"github.com/gorilla/mux"
)

// NewRoutes setups the complete routing for the app.
func NewRoutes(mux *mux.Router) {
	// Health check
	mux.HandleFunc("/health", handler.HealthCheck)

	SetupAuthRoutes(mux, &handler.Auth{})

	SetupSwipeRoutes(mux, &handler.Swipe{DB: db.DB, Redis: redis.Client})

	SetupUpgradeRoutes(mux, &handler.Upgrade{})

	SetupProfileRoutes(mux, &handler.Profile{})
}
