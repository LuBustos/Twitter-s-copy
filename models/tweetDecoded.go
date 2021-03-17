package models

//TweetDecoded es el que va a ser decodificado cuando venga del body
type TweetDecoded struct {
	Message string `bson:"message" json:"message"`
}
