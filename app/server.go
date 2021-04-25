package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/najibean/crudGo/database/seeders"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.Name)

	// server.initializeDB(dbConfig)
	server.initializeRoutes()
	// seeders.DBSeed(server.DB)
}

// menjalankan server port
func (server *Server) Run(port string) {
	fmt.Printf("Listening to port%s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database server")
	} else {
		fmt.Println("Connected to database postgres")
	}

}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("database migrated successfully")
}

func (server *Server) initCommands(appConfig AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
	var dbConfig = DBConfig{}

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

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "admin")
	dbConfig.DBName = getEnv("DB_NAME", "rajaKado")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.Port)
	}
}
