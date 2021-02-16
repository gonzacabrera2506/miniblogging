package middlew

import (
	"net/http"

	"github.com/gonzacabrera2506/miniblogging/bd"
)

//ChequeoBD es el middleware que me permite conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
