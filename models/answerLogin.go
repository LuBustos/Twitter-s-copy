package models

//AnswerLogin es la clase donde va almacenar la respuesta del login
type AnswerLogin struct {
	Token string `json:"website,omitempty"`
}
