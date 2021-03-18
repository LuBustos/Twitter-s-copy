package routers

import (
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
)

//Delete llama al metodo deleteTweet de lo que vino del front
func Delete(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "No se pudo eliminar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
