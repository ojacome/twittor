package routers

import (
	"encoding/json"
	"github.com/ojacome/twittor/db"
	"github.com/ojacome/twittor/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Cuerpo de la petición inválido. "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email es requerido. ", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password es requerido. ", 400)
		return
	}

	_, encontrado, _ := db.ExisteUsuario(user.Email)
	if encontrado {
		http.Error(w, "El usuario ya existe en la base. ", 400)
		return
	}
	_, status, err := db.InsertarRegistro(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if !status {
		http.Error(w, "Error al insertar registro. ", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
