package bd

import (
	"context"
	"time"

	"github.com/LuBustos/Twitter-s-copy/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserAlreadyExistCheck chequea si el usuario existe
func UserAlreadyExistCheck(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter-s-copy-des")
	col := db.Collection("user")

	condicion := bson.M{"email": email} //Tiene que ser formato JSON
	var result models.User
	err := col.FindOne(ctx, condicion).Decode(&result) //Busca un solo registro
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
