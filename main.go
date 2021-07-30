package main

import (
	"github.com/ojacome/twittor/db"
	"github.com/ojacome/twittor/handlers"
	"log"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Sin conexión a MongoDb")
		return
	}
	handlers.Manejadores()
}
