// routes/auth_routes.go
package routes

import (
	"dating-sim/handler" // Import the handler package

	"github.com/gorilla/mux"
)

// SetupAuthRoutes sets up the routes for authentication (e.g., signup, login).
func SetupAuthRoutes(mux *mux.Router, handler *handler.Auth) {
	prefix := mux.PathPrefix("/api/v1/auth").Subrouter()
	{
		prefix.HandleFunc("/signup", handler.Signup)
		prefix.HandleFunc("/login", handler.Login)
	}
}
