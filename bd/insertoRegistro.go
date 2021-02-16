package bd

import (
	"context"
	"time"

	"github.com/gonzacabrera2506/miniblogging/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es el ultimo paso para insertar el registro en la BD
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("ususarios")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
