package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/AlanProgrammer93/twitter-go-react/middleware"
	"github.com/AlanProgrammer93/twitter-go-react/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.CheckDB(middleware.ValidJWT(routers.ViewProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
