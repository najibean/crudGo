package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB 		*gorm.DB
	Router	*mux.Router
}

type AppConfig struct {
	Name	string
	Env		string
	Port 	string
}

// memanggil routes -- kenapa tidak dimaksudkan untuk para router saja ya?
func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.Name)

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// menjalankan server port
func (server *Server) Run(port string) {
	fmt.Printf("Listening to port%s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

// membuat func nilai default .env jika data didalamnya kosong. Karena di Go jika data hilang, server port masih jalan
func getEnv(key, defaultnya string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultnya
}

// menjalankan semua fungsi yang ada diatas
func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	// appConfig.Name = os.Getenv("APP_NAME")
	// appConfig.Env = os.Getenv("APP_ENV")
	// appConfig.Port = os.Getenv("APP_PORT")

	// tes fungsi jika data di .env tidak ada
	appConfig.Name = getEnv("APP_NAME", "waroeng boutique")
	appConfig.Env = getEnv("APP_ENV", "environment")
	appConfig.Port = getEnv("APP_PORT", "9989")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.Port)
}