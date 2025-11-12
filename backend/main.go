package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"alchemy-system/backend/controllers"
	"alchemy-system/backend/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.ConnectDatabase()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/alchemists", controllers.GetAllAlchemists).Methods("GET")
	api.HandleFunc("/alchemists/{id}", controllers.GetAlchemistByID).Methods("GET")
	api.HandleFunc("/alchemists", controllers.CreateAlchemist).Methods("POST")
	api.HandleFunc("/alchemists/{id}", controllers.UpdateAlchemist).Methods("PUT")
	api.HandleFunc("/alchemists/{id}", controllers.DeleteAlchemist).Methods("DELETE")

	api.HandleFunc("/missions", controllers.GetAllMissions).Methods("GET")
	api.HandleFunc("/missions/{id}", controllers.GetMissionByID).Methods("GET")
	api.HandleFunc("/missions", controllers.CreateMission).Methods("POST")
	api.HandleFunc("/missions/{id}", controllers.UpdateMission).Methods("PUT")
	api.HandleFunc("/missions/{id}", controllers.DeleteMission).Methods("DELETE")

	api.HandleFunc("/transmutations", controllers.GetAllTransmutations).Methods("GET")
	api.HandleFunc("/transmutations/{id}", controllers.GetTransmutationByID).Methods("GET")
	api.HandleFunc("/transmutations", controllers.CreateTransmutation).Methods("POST")
	api.HandleFunc("/transmutations/{id}", controllers.UpdateTransmutation).Methods("PUT")
	api.HandleFunc("/transmutations/{id}", controllers.DeleteTransmutation).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
