package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
)

//DeleteRelationship elimina la relacion
func DeleteRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("relationship")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
