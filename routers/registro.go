package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gonzacabrera2506/miniblogging/bd"
	"github.com/gonzacabrera2506/miniblogging/models"
)

//Registro crea en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}
	//si no hay error en el body hago validaciones
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres ", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese Email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario", 400)
		return
	}
	/*Si llegó hasta acá todo anduvo bien*/
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
