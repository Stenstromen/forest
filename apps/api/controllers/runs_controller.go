package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/stenstromen/forest/api/models"
)

func CreateRun(w http.ResponseWriter, r *http.Request) {
	var run models.Run
	err := json.NewDecoder(r.Body).Decode(&run)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	runResponse := models.RunResponse{
		ID:         uuid.New().String(),
		OccurredAt: run.OccurredAt,
		Distance:   run.Distance,
		Duration:   run.Duration,
		Calories:   run.Calories,
	}
	err = json.NewEncoder(w).Encode(runResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
