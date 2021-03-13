package main

import (
	"log"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/handlers"
)

func main() {
	if !bd.ConnectionCheck() {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Handlers()
}
