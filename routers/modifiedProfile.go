package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

//ModifiedProfile se encarga de modificar el perfil
func ModifiedProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModifyRegistry(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, intente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No encontro el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
