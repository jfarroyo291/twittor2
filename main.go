package main

import (
	"log"

	"github.com/jfarroyo291/twittor2/bd"
	"github.com/jfarroyo291/twittor2/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
