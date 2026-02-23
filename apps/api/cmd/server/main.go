package main

import (
	"log"
	"net/http"
	"time"

	"github.com/stenstromen/forest/api/controllers"
)

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

	v1.HandleFunc("POST /runs", controllers.CreateRun)

	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
