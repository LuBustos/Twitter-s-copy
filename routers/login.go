package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/jwt"
	"github.com/LuBustos/Twitter-s-copy/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario/contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido"+err.Error(), 400)
		return
	}
	documento, existe := bd.TryLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario/contraseña invalida", 400)
		return
	}

	token, err := jwt.GenerateJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al generar el token"+err.Error(), 400)
		return
	}
	resp := models.AnswerLogin{
		Token: token,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Como guardar el token en la cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
}
