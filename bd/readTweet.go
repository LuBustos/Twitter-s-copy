package bd

import (
	"context"
	"log"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTweet lee tweets y utiliza paginado en mongodb
func ReadTweet(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //Cancela el contexto y da de baja el timeout
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("tweet")

	var results []*models.ReturnTweet

	condition := bson.M{
		"userid": ID,
	}
	//Opciones en mongo
	opts := options.Find() //Propiedades que trae el find
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	//Puntero a una tabla de una base de datos donde se van a guardar los datos
	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registry models.ReturnTweet
		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}
	return results, true
}
