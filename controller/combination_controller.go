package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PandaX185/computer-alchemy-api/dto"
	"github.com/PandaX185/computer-alchemy-api/service"
)

// CombineElements godoc
// @Summary Combine elements
// @Description Combines two elements and returns the resulting element
// @Tags combinations
// @Accept json
// @Produce json
// @Param combination body dto.CombinationRequest true "Elements to combine"
// @Success 200 {array} models.Element
// @Failure 400 {object} string "Invalid combination"
// @Router /combinations [post]
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

	result, err := service.CombineElements(request.FirstElement, request.SecondElement)
	if err != nil {
		http.Error(w, "Invalid combination", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetAllCombinations godoc
// @Summary Get all combinations
// @Description Returns a list of all combinations
// @Tags combinations
// @Produce json
// @Success 200 {array} dto.CombinationResponse
// @Param element query string false "Element name"
// @Router /combinations [get]
func GetAllCombinations(w http.ResponseWriter, r *http.Request) {
	element := r.URL.Query().Get("element")
	var combinations []*dto.CombinationResponse
	if element == "" {
		combinations = service.GetAllCombinations()
	} else {
		combinations = service.GetAllElementCombinations(element)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(combinations)
}

// GetAllResultCombinations godoc
// @Summary Get all combinations for a resulting element
// @Description Returns a list of all combinations for a resulting element
// @Tags combinations
// @Produce json
// @Param resultingElement query string true "Resulting element name"
// @Success 200 {array} dto.CombinationResponse
// @Router /combinations/result [get]
func GetAllResultCombinations(w http.ResponseWriter, r *http.Request) {
	resultingElement := r.URL.Query().Get("resultingElement")
	combinations := service.GetAllResultCombinations(resultingElement)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(combinations)
}
