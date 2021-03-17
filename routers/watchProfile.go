package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
)

//WatchProfile recibe el ID de la url, lo busca en la bd y lo trae
func WatchProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
