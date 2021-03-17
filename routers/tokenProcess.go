package routers

import (
	"errors"
	"strings"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email del usuario
var Email string

//IDUser
var IDUser string

//TokenProcess procesa el token a enviar
func TokenProcess(tk string) (*models.Claim, bool, string, error) { //El error va al final
	pass := []byte("MasterdelDesarrollo_Grupo")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return pass, nil
	})
	if err == nil {
		_, encontrado, _ := bd.UserAlreadyExistCheck(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, encontrado, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}
