package routers

import (
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

func RegisterRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID
	status, err := bd.InsertRelationship(t)
	if err != nil {
		http.Error(w, "Error al insertar relacion", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
