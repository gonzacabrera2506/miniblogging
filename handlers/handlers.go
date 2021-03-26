package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gonzacabrera2506/miniblogging/middlew"
	"github.com/gonzacabrera2506/miniblogging/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores setea el puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
