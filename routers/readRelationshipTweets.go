package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuBustos/Twitter-s-copy/bd"
)

//ReadRelationshipTweets Le envia al front los tweets
func ReadRelationshipTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //string to integer
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	answer, correct := bd.ReadFollowTweets(IDUser, pagina)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
