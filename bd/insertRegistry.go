package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegistry(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //Cancela el contexto y da de baja el timeout
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("users")

	u.Password, _ = Encrypt(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
