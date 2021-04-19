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

func (server *Server) Initialize() {
	fmt.Println("Welcome to Raja Kado & Boutique")

	server.Router = mux.NewRouter()
}

func (server *Server) Run(port string) {
	fmt.Printf("Listening to port%s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9000")
}