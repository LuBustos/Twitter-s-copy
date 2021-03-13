package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

func Registry(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t) //Es un stream de solo lectura, una vez que se lee, se destruye
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.AlreadyUserCheck(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertRegistry(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al guardar el registro del usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
