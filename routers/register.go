package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AlanProgrammer93/twitter-go-react/db"
	"github.com/AlanProgrammer93/twitter-go-react/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido.", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres.", 400)
		return
	}

	_, encontrado, _ := db.CheckExistUser(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email.", 400)
		return
	}

	_, status, err := db.createUser(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro.", 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo realizar el registro.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
