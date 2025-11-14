package controllers

import (
	 "alchemy-system/models"
    "alchemy-system/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllMissions handles GET /missions
func GetAllMissions(w http.ResponseWriter, r *http.Request) {
	missions, err := services.GetAllMissions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missions)
}

// GetMissionByID handles GET /missions/{id}
func GetMissionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mission, err := services.GetMissionByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}

// CreateMission handles POST /missions
func CreateMission(w http.ResponseWriter, r *http.Request) {
	var mission models.Mission
	if err := json.NewDecoder(r.Body).Decode(&mission); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	created, err := services.CreateMission(mission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

// UpdateMission handles PUT /missions/{id}
func UpdateMission(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated models.Mission
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mission, err := services.UpdateMission(uint(id), updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mission)
}

// DeleteMission handles DELETE /missions/{id}
func DeleteMission(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteMission(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
	