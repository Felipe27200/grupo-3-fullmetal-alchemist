package controllers

import (
	"alchemy-system/backend/models"
	"alchemy-system/backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllTransmutations handles GET /transmutations
func GetAllTransmutations(w http.ResponseWriter, r *http.Request) {
	transmutations, err := services.GetAllTransmutations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transmutations)
}

// GetTransmutationByID handles GET /transmutations/{id}
func GetTransmutationByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transmutation, err := services.GetTransmutationByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transmutation)
}

// CreateTransmutation handles POST /transmutations
func CreateTransmutation(w http.ResponseWriter, r *http.Request) {
	var transmutation models.Transmutation
	if err := json.NewDecoder(r.Body).Decode(&transmutation); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	created, err := services.CreateTransmutation(transmutation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

// UpdateTransmutation handles PUT /transmutations/{id}
func UpdateTransmutation(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated models.Transmutation
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	transmutation, err := services.UpdateTransmutation(uint(id), updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transmutation)
}

// DeleteTransmutation handles DELETE /transmutations/{id}
func DeleteTransmutation(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteTransmutation(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
