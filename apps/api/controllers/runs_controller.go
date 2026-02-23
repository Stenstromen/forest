package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/stenstromen/forest/api/models"
)

func CreateRun(w http.ResponseWriter, r *http.Request) {
	var run models.Run
	var duration float64
	var distance float64
	var PaceMinutesPerKm float64
	var PaceMinutesPerMile float64
	runId := "run_" + uuid.New().String()

	err := json.NewDecoder(r.Body).Decode(&run)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if run.Distance.Unit == "km" {
		distance = run.Distance.Value * 1000
	} else {
		distance = run.Distance.Value * 1609.34
	}

	if run.Duration.Unit == "min" {
		duration = run.Duration.Value * 60
	} else {
		duration = run.Duration.Value
	}

	durationMinutes := duration / 60
	PaceMinutesPerKm = durationMinutes / (distance / 1000)
	PaceMinutesPerMile = durationMinutes / (distance / 1609.34)

	runResponse := models.RunResponse{
		ID:         runId,
		OccurredAt: run.OccurredAt,
		Distance:   distance,
		Duration:   duration,
		Calories:   run.Calories,
		Derived: models.Derived{
			PaceMinutesPerKm:   PaceMinutesPerKm,
			PaceMinutesPerMile: PaceMinutesPerMile,
		},
	}
	err = json.NewEncoder(w).Encode(runResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
