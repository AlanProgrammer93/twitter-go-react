package routers

import (
	"net/http"

	"github.com/AlanProgrammer93/twitter-go-react/db"
	"github.com/AlanProgrammer93/twitter-go-react/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar dejar de seguir."+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se puso dejar de seguir al usuario."+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
