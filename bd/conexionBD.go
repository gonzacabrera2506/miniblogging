package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion a bd
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gonzacabrera2506:Dios2021@cluster0.7wqpa.mongodb.net/microblogging?retryWrites=true&w=majority")

//ConectarBD es la funcion para conectar a la bd
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal((err.Error()))
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa a BD")
	return client
}

//ChequeoConnection hace un ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
