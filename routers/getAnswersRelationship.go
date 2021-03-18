package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

//GetAnswerRelationship le devuelve el status de la relacion
func GetAnswerRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	var res models.AnswerRelatonship

	status, err := bd.GetRelationship(t)
	if err != nil || status == false {
		res.Status = false
	}

	res.Status = true

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
