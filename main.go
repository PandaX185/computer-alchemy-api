package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PandaX185/computer-alchemy-api/controller"
	"github.com/PandaX185/computer-alchemy-api/seed"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := seed.SeedElements(); err != nil {
		log.Fatal(err)
	}

	log.Println("Elements seeded successfully")

	if err := seed.SeedRelations(); err != nil {
		log.Fatal(err)
	}

	log.Println("Relations seeded successfully")

}

func main() {
	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}

	router := mux.NewRouter()
	router.HandleFunc("/elements", controller.GetAllElements).Methods("GET")
	router.HandleFunc("/elements/{name}", controller.GetElementByName).Methods("GET")
	router.HandleFunc("/elements", controller.GetCombinationResult).Methods("POST")

	server.Handler = router

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
