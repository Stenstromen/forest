package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/stenstromen/forest/api/controllers"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	v1 := http.NewServeMux()

	v1.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	v1.HandleFunc("GET /ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Create a new run
	v1.HandleFunc("POST /runs", controllers.CreateRun)
	// Get all runs
	v1.HandleFunc("GET /runs", controllers.GetRuns)
	// Get a specific run
	v1.HandleFunc("GET /runs/{id}", controllers.GetRun)
	// Update a specific run
	v1.HandleFunc("PUT /runs/{id}", controllers.UpdateRun)
	// Patch a specific run
	v1.HandleFunc("PATCH /runs/{id}", controllers.PatchRun)
	// Delete a specific run
	v1.HandleFunc("DELETE /runs/{id}", controllers.DeleteRun)

	mux := http.NewServeMux()
	mux.Handle("/v1/", jsonMiddleware(http.StripPrefix("/v1", v1)))

	tcpPort := os.Getenv("TCP_PORT")
	if tcpPort == "" {
		tcpPort = "8080"
	}

	server := &http.Server{
		Addr:         ":" + tcpPort,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
