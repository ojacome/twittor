package handlers

import (
	"github.com/gorilla/mux"
	"github.com/ojacome/twittor/middlewares"
	"github.com/ojacome/twittor/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// Manejadores control del cors, configuraciones extras y levanto el servidor.
func Manejadores() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckBD(routers.Registro)).Methods("POST")

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
