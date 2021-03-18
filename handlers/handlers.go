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

	router.HandleFunc("/registro", middleware.BDCheck(routers.Registry)).Methods("POST")
	router.HandleFunc("/login", middleware.BDCheck(routers.Login)).Methods("POST")

	router.HandleFunc("/perfil", middleware.BDCheck(middleware.JwtCheck(routers.WatchProfile))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleware.BDCheck(middleware.JwtCheck(routers.ModifiedProfile))).Methods("PUT")

	router.HandleFunc("/tweet", middleware.BDCheck(middleware.JwtCheck(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/subirBanner", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middleware.BDCheck(middleware.JwtCheck(routers.ReadTweets))).Methods("GET")

	router.HandleFunc("/altaRelacion", middleware.BDCheck(middleware.JwtCheck(routers.RegisterRelationship))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middleware.BDCheck(middleware.JwtCheck(routers.RemoveRelationship))).Methods("DELETE")
	router.HandleFunc("/getRelacion", middleware.BDCheck(middleware.JwtCheck(routers.GetAnswerRelationship))).Methods("GET")

	router.HandleFunc("/listUsuarios", middleware.BDCheck(middleware.JwtCheck(routers.UserList))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middleware.BDCheck(middleware.JwtCheck(routers.ReadRelationshipTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //Le da permisos a cualqueira
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
