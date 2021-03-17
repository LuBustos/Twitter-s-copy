package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegistry(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("users")

	registry := make(map[string]interface{})
	if len(u.Name) > 0 {
		registry["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		registry["surname"] = u.Surname
	}
	if len(u.Password) > 0 {
		registry["password"] = u.Password
	}
	if len(u.Avatar) > 0 {
		registry["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registry["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		registry["biography"] = u.Biography
	}
	if len(u.Website) > 0 {
		registry["website"] = u.Website
	}

	updString := bson.M{
		"$set": registry,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := col.UpdateOne(ctx, filter, updString)
	if err != nil {
		return false, err
	}
	return true, nil
}
