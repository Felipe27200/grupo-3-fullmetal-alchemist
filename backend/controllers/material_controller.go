package controllers

import (
	"encoding/json"
	"net/http"

	"alchemy-system/backend/models"
	"alchemy-system/backend/services"

	"github.com/gorilla/mux"
)

func GetAllMaterials(w http.ResponseWriter, r *http.Request) {
	materials, err := services.GetAllMaterials()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(materials)
}

func GetMaterialByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	material, err := services.GetMaterialByID(id)
	if err != nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(material)
}

func CreateMaterial(w http.ResponseWriter, r *http.Request) {
	var material models.Material
	json.NewDecoder(r.Body).Decode(&material)

	newMaterial, err := services.CreateMaterial(material)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newMaterial)
}

func UpdateMaterial(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var data models.Material
	json.NewDecoder(r.Body).Decode(&data)

	updated, err := services.UpdateMaterial(id, data)
	if err != nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func DeleteMaterial(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := services.DeleteMaterial(id); err != nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
