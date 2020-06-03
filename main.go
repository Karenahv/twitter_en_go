package main

import (
	"log"

	"github.com/Karenahv/twitter_en_go/bd"
	"github.com/Karenahv/twitter_en_go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la BD")
	}
	handlers.Manejadores()

}
