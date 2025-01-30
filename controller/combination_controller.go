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
