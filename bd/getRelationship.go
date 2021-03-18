package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetRelationship trae la lista de relaciones
func GetRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("relationship")

	condition := bson.M{
		"userid":             t.UserID,
		"userrelationshipid": t.UserRelationshipID,
	}
	var result models.Relationship
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
