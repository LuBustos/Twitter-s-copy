package routers

import (
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

//RemoveRelationship llama al metodo deleteRelationship de lo que vino del front
func RemoveRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID
	status, err := bd.DeleteRelationship(t)
	if err != nil {
		http.Error(w, "Error al eliminar relacion", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado eliminar la relacion", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
