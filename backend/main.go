package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"alchemy-system/controllers"
	"alchemy-system/database"
	"alchemy-system/middleware"
	"alchemy-system/queue"
	"alchemy-system/worker"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.ConnectDatabase()
	queue.InitRedis()

	go worker.StartTransmutationWorker()

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/auth/register", controllers.Register).Methods("POST")
	api.HandleFunc("/auth/login", controllers.Login).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.Handle("/test/auth", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(uint)
		role := r.Context().Value("role").(string)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"user_id": userID,
			"role":    role,
			"status":  "middleware OK",
		})
	})).Methods("GET")

	protected.HandleFunc("/alchemists", controllers.GetAllAlchemists).Methods("GET")
	protected.HandleFunc("/alchemists/{id}", controllers.GetAlchemistByID).Methods("GET")
	protected.HandleFunc("/alchemists", controllers.CreateAlchemist).Methods("POST")
	protected.HandleFunc("/alchemists/{id}", controllers.UpdateAlchemist).Methods("PUT")
	protected.HandleFunc("/alchemists/{id}", controllers.DeleteAlchemist).Methods("DELETE")

	protected.HandleFunc("/missions", controllers.GetAllMissions).Methods("GET")
	protected.HandleFunc("/missions/{id}", controllers.GetMissionByID).Methods("GET")
	protected.HandleFunc("/missions", controllers.CreateMission).Methods("POST")
	protected.HandleFunc("/missions/{id}", controllers.UpdateMission).Methods("PUT")
	protected.HandleFunc("/missions/{id}", controllers.DeleteMission).Methods("DELETE")

	protected.HandleFunc("/transmutations", controllers.GetAllTransmutations).Methods("GET")
	protected.HandleFunc("/transmutations/{id}", controllers.GetTransmutationByID).Methods("GET")
	protected.HandleFunc("/transmutations", controllers.CreateTransmutation).Methods("POST")
	protected.HandleFunc("/transmutations/{id}", controllers.UpdateTransmutation).Methods("PUT")
	protected.HandleFunc("/transmutations/{id}", controllers.DeleteTransmutation).Methods("DELETE")

	protected.HandleFunc("/materials", controllers.GetAllMaterials).Methods("GET")
	protected.HandleFunc("/materials/{id}", controllers.GetMaterialByID).Methods("GET")
	protected.HandleFunc("/materials", controllers.CreateMaterial).Methods("POST")
	protected.HandleFunc("/materials/{id}", controllers.UpdateMaterial).Methods("PUT")
	protected.HandleFunc("/materials/{id}", controllers.DeleteMaterial).Methods("DELETE")

	protected.HandleFunc("/audits", controllers.GetAllAudits).Methods("GET")
	protected.HandleFunc("/audits/{id}", controllers.GetAuditByID).Methods("GET")
	protected.HandleFunc("/audits", controllers.CreateAudit).Methods("POST")
	protected.HandleFunc("/audits/{id}", controllers.DeleteAudit).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
