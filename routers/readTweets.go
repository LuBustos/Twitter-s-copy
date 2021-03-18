package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuBustos/Twitter-s-copy/bd"
)

//ReadTweets se encarga de traer los tweets mandarlos al front
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //string to integer
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)
	answer, correct := bd.ReadTweet(ID, pag)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
