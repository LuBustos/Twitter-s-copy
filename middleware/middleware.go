package middleware

import (
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
)

//BDCheck es el middleware que permite conocer el estado de la base de datos
func BDCheck(next http.HandlerFunc) http.HandlerFunc {
	//Middleware recibe y retorna lo mismo

	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.ConnectionCheck() {
			http.Error(w, "Conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w, r) //Le pasa todos los valores que recibe al siguiente eslabon
	}
}
