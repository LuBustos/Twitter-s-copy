package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertTweet se encarga de insertar un tweet en la base de datos
func InsertTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //Cancela el contexto y da de baja el timeout
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("tweet")

	registry := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, registry)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil
}
