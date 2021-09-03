package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlanProgrammer93/twitter-go-react/db"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID.", http.StatusBadRequest)
		return
	}

	perfil, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el perfil.", 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
