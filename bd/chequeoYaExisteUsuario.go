package bd

import (
	"context"
	"time"

	"github.com/gonzacabrera2506/miniblogging/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email de parametro y se fija si ya exsite en BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("miniblogging")
	col := db.Collection("usuarios")

	// M es una función que formatea o mapea a bson lo que recibe como json
	condicion := bson.M{"email": email}

	//en la variable resultado voy a modelar un usuario
	var resultado models.Usuario

	//FindOne me devuelve un sólo registro que cumple con la condición
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
