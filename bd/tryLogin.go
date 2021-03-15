package bd

import (
	"github.com/LuBustos/Twitter-s-copy/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin valida si el email existe y si la contraseña ingresada son iguales
func TryLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := UserAlreadyExistCheck(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
