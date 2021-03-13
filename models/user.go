package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User table que se va a encontrar en mongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"` //omitempty si el campo viene vacio, lo omite
	Name      string             `bson:"name" json:"name"`
	Surname   string             `bson:"surname" json:"surname"`
	Birthday  time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}

//JSON es lo que viene en el navegador
//Bson es mongodb
