package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlanProgrammer93/twitter-go-react/db"
	"github.com/AlanProgrammer93/twitter-go-react/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateRegister(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al modificar el registro.", 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
