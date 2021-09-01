package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/twitter-clone")

/*ConectarBD Permite conectar la base de datos*/
func ConectarBD() *mongo.Client {
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
	log.Println("Conexion exitosa con la base de datos")
	return client
}

/*CheckConnection Permite chequear la base de datos*/
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
