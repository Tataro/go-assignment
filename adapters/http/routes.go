package http

import (
	"github.com/gorilla/mux"
	"gitlab.com/upaphong/go-assignment/engine"
	"gitlab.com/upaphong/go-assignment/handlers"
)

func RegisterRoutes(e engine.Engine, r *mux.Router) {
	// Knight routes
	r.HandleFunc("/knight", handlers.HandleListKnights(e)).Methods("GET")
	r.HandleFunc("/knight", handlers.HandleSaveKnight(e)).Methods("POST")
	r.HandleFunc("/knight/{id}", handlers.HandleGetKnight(e)).Methods("GET")
}
