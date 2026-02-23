package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/stenstromen/forest/api/models"
)

var runs = make(map[string]models.RunResponse)

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
	runs[runId] = runResponse
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(runResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetRuns(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(runs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRun(w http.ResponseWriter, r *http.Request) {
	runId := r.PathValue("id")
	run, ok := runs[runId]
	if !ok {
		http.Error(w, "Run not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(run)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateRun(w http.ResponseWriter, r *http.Request) {
	var run models.Run
	var duration float64
	var distance float64
	var PaceMinutesPerKm float64
	var PaceMinutesPerMile float64
	runId := r.PathValue("id")

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
	runs[runId] = runResponse
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(runResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PatchRun(w http.ResponseWriter, r *http.Request) {
	runId := r.PathValue("id")
	existing, ok := runs[runId]
	if !ok {
		http.Error(w, "Run not found", http.StatusNotFound)
		return
	}

	var patch models.RunPatch
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if patch.OccurredAt != nil {
		existing.OccurredAt = *patch.OccurredAt
	}
	if patch.Distance != nil {
		if patch.Distance.Unit == "km" {
			existing.Distance = patch.Distance.Value * 1000
		} else {
			existing.Distance = patch.Distance.Value * 1609.34
		}
	}
	if patch.Duration != nil {
		if patch.Duration.Unit == "min" {
			existing.Duration = patch.Duration.Value * 60
		} else {
			existing.Duration = patch.Duration.Value
		}
	}
	if patch.Calories != nil {
		existing.Calories = *patch.Calories
	}

	durationMinutes := existing.Duration / 60
	existing.Derived.PaceMinutesPerKm = durationMinutes / (existing.Distance / 1000)
	existing.Derived.PaceMinutesPerMile = durationMinutes / (existing.Distance / 1609.34)

	runs[runId] = existing
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(existing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteRun(w http.ResponseWriter, r *http.Request) {
	runId := r.PathValue("id")
	delete(runs, runId)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(runs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
