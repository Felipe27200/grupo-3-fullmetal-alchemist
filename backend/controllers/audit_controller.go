package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"alchemy-system/models"
	"alchemy-system/services"
	"github.com/gorilla/mux"
)

func GetAllAudits(w http.ResponseWriter, r *http.Request) {
	audits, err := services.GetAllAudits()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(audits)
}

func GetAuditByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	audit, err := services.GetAuditByID(uint(id))
	if err != nil {
		http.Error(w, "audit not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(audit)
}

func CreateAudit(w http.ResponseWriter, r *http.Request) {
	var audit models.Audit

	if err := json.NewDecoder(r.Body).Decode(&audit); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	created, err := services.CreateAudit(audit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func DeleteAudit(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = services.DeleteAudit(uint(id))
	if err != nil {
		http.Error(w, "audit not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
