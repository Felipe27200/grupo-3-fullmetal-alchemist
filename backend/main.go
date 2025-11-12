package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"alchemy-system/backend/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	// Connect to database (creates tables automatically)
	database.ConnectDatabase()

	// Basic route
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Alchemy API is running and database is ready!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	http.ListenAndServe(":"+port, r)
}
