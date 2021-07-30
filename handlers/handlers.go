package handlers

import (
	"github.com/gorilla/mux"
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
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
