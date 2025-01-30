package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PandaX185/computer-alchemy-api/service"
	"github.com/gorilla/mux"
)

func GetAllElements(w http.ResponseWriter, r *http.Request) {
	elements := service.GetAllElements()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(elements)
}

func GetElementByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	element := service.GetElementByName(name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(element)
}

func GetCombinationResult(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var request struct {
		FirstElement  string `json:"firstElement"`
		SecondElement string `json:"secondElement"`
	}
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result := service.GetCombinationResult(request.FirstElement, request.SecondElement)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
