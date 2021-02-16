package main

import (
	"log"

	"github.com/gonzacabrera2506/miniblogging/bd"
	"github.com/gonzacabrera2506/miniblogging/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a BD")
		return
	}
	handlers.Manejadores()
}
