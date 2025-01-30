package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/PandaX185/computer-alchemy-api/docs" // Add this line for swagger docs

	"github.com/PandaX185/computer-alchemy-api/controller"
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
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Group element-related routes
	elementRouter := router.PathPrefix("/api/elements").Subrouter()
	elementRouter.HandleFunc("", controller.GetAllElements).Methods("GET")
	elementRouter.HandleFunc("/{name}", controller.GetElementByName).Methods("GET")

	// Group combination-related routes
	combinationRouter := router.PathPrefix("/api/combinations").Subrouter()
	combinationRouter.HandleFunc("", controller.CombineElements).Methods("POST")

	server.Handler = router

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
