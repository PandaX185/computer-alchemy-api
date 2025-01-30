package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PandaX185/computer-alchemy-api/controller"
	_ "github.com/PandaX185/computer-alchemy-api/docs"
	"github.com/PandaX185/computer-alchemy-api/seed"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := seed.SeedElements(); err != nil {
		log.Fatal(err)
	}

	log.Println("Elements seeded successfully")

	if err := seed.SeedCombinations(); err != nil {
		log.Fatal(err)
	}

	log.Println("Combinations seeded successfully")

}

// @title Computer Alchemy API
// @version 1.0
// @description This is the API server for the Computer Alchemy game.
// @host localhost:8080
// @BasePath /api
func main() {
	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}

	router := mux.NewRouter()

	// Swagger documentation endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	router.HandleFunc("/api/elements", controller.GetAllElements).Methods("GET")

	router.HandleFunc("/api/elements/{name}", controller.GetElementByName).Methods("GET")

	router.HandleFunc("/api/elements", controller.CombineElements).Methods("POST")

	server.Handler = router

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
