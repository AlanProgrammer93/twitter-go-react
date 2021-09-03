package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AlanProgrammer93/twitter-go-react/db"
	"github.com/AlanProgrammer93/twitter-go-react/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar publicar el Tweet.", 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo publicar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
