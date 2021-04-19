package app

import (
	"github.com/najibean/crudGo/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}