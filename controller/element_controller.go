package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PandaX185/computer-alchemy-api/dto"
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
	element := service.GetElementByName(name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(element)
}

// CombineElements godoc
// @Summary Combine elements
// @Description Combines two elements and returns the resulting element
// @Tags elements
// @Accept json
// @Produce json
// @Param combination body dto.CombinationRequest true "Elements to combine"
// @Success 200 {array} models.Element
// @Failure 400 {object} string "Invalid combination"
// @Router /elements [post]
func CombineElements(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var request dto.CombinationRequest
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := service.CombineElements(request.FirstElement, request.SecondElement)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
