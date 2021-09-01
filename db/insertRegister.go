package db

import (
	"context"
	"time"
	"github.com/AlanProgrammer93/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twitter-clone")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false , err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}