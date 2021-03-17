package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LuBustos/Twitter-s-copy/bd"
	"github.com/LuBustos/Twitter-s-copy/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.TweetDecoded
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.Tweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertTweet(registry)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
