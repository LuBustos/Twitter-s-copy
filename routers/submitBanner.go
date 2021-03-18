package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

//SubmitBanner nos permite subir el banner al server
func SubmitBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1] //Separar el nombre del archivo con el punto, nos quedamos con la extension
	var arch string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(arch, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = bd.ModifyRegistry(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la base de datos"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
