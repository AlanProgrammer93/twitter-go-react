package main

import (
	"log"

	"github.com/AlanProgrammer93/twitter-go-react/db"
	"github.com/AlanProgrammer93/twitter-go-react/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
