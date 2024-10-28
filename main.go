package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/hardikkum444/goRSS/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL does not exist in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to databse: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", checkServerHealth)
	v1router.Get("/err", handlerErr)
	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("server will now run on -> " + portString)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
