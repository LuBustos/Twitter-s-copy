package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN recibe la conexion y es exportada a todos los miembros de bd
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbDev:a1b2c3d4@cluster0.offld.mongodb.net/test?authSource=admin&replicaSet=atlas-t9lae9-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

//ConectarBD es la funcion que me permite conectarme a la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client //Devuelvo el objeto por mas que este vacio
	}
	err = client.Ping(context.TODO(), nil) //Para ver si la base de datos esta activa
	if err != nil {
		log.Fatal(err.Error())
		return client //Devuelvo el objeto por mas que este vacio
	}
	log.Println("Conexion exitosa con la bd")
	return client
} //TODO conectate a la base de datos sin ninguna restriccion

//ConnectionCheck verifica si la conexion a la base de datos esta funcionando
func ConnectionCheck() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
