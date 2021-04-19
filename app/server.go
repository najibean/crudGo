package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB 		*gorm.DB
	Router	*mux.Router
}

// memanggil routes -- kenapa tidak dimaksudkan untuk para router saja ya?
func (server *Server) Initialize() {
	fmt.Println("Welcome to Raja Kado & Boutique")

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// menjalankan server port
func (server *Server) Run(port string) {
	fmt.Printf("Listening to port%s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

// menjalankan semua fungsi yang ada diatas
func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9000")
}