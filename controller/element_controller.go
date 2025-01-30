package controller

import (
	"encoding/json"
	"net/http"

	"github.com/PandaX185/computer-alchemy-api/service"
	"github.com/gorilla/mux"
)

// GetAllElements godoc
// @Summary Get all available elements
// @Description Returns a list of all elements in the game
// @Tags elements
// @Produce json
// @Success 200 {array} models.Element
// @Router /elements [get]
func GetAllElements(w http.ResponseWriter, r *http.Request) {
	elements := service.GetAllElements()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(elements)
}

// GetElementByName godoc
// @Summary Get element by name
// @Description Returns a specific element by its name
// @Tags elements
// @Produce json
// @Param name path string true "Element name"
// @Success 200 {object} models.Element
// @Failure 404 {object} string "Element not found"
// @Router /elements/{name} [get]
func GetElementByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	element, err := service.GetElementByName(name)
	if err != nil || element == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(element)
}
