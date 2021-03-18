package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadFollowTweets leo los tweets de los seguidores
func ReadFollowTweets(ID string, pagina int) ([]models.ReturnFollowTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("relationship")

	skip := (pagina - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":        "tweet",
			"localField":  "userrelationshipid",
			"foreingFiel": "userid",
			"as":          "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnFollowTweets
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
