package middleware

import (
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/routers"
)

func JwtCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
	}
}
