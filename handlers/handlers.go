package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/LuBustos/Twitter-s-copy/middleware"
	"github.com/LuBustos/Twitter-s-copy/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers seteo mi handler y a escuchar el servidor
func Handlers() {
	router := mux.NewRouter() //

	router.HandleFunc("/registro", middleware.BDCheck(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //Le da permisos a cualqueira
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
