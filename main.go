package main

import (
	"dating-sim/db"
	migrate "dating-sim/migration"
	"dating-sim/redis"
	"dating-sim/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	db.Conn()
	redis.Init()
	migrate.Db()
}

func main() {
	route := mux.NewRouter()
	routes.NewRoutes(route)

	listener := &http.Server{
		Addr:         ":5000",
		Handler:      route,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	defer func() {
		if r := recover(); r != nil {
			log.Print("recovered: error:", r)
		}
		listener.Close()
	}()
	log.Print("listening at " + listener.Addr)
	log.Fatal(listener.ListenAndServe())
}
