package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
)

//InsertRelationship graba en la BD la relacion
func InsertRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //Cancela el contexto y da de baja el timeout
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("relationship")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
