package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoCN objeto de conexión.
var MongoCN = conectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://portafolio_user:portafolio_password@cluster0.njduj.mongodb.net/twittor")

func conectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conectado a MongoDB.")
	return client
}

// CheckConnection verificar la conexión.
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
