package bd

import "golang.org/x/crypto/bcrypt"

//Encrypt nos permite encriptar passwords
func Encrypt(pass string) (string, error) {
	cost := 8 //Es el nivel de seguridad
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
