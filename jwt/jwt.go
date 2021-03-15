package jwt

import (
	"github.com/LuBustos/Twitter-s-copy/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.User) (string, error) {
	password := []byte("MasterdelDesarrollo_Grupo")
	payload := jwt.MapClaims{
		"email":    t.Email,
		"nombre":   t.Name,
		"apellido": t.Surname,
		"_id":      t.ID.Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(password)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
