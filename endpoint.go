package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Printf(
			"Received request for %s from %s with Authorization header: \"%s\"",
			r.URL.Path,
			r.RemoteAddr,
			authHeader,
		)

		var payload struct {
			Message  string `json:"message"`
			Hostname string `json:"hostname"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Printf("error decoding request body: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		log.Printf("Received message: \"%s\" from hostname: \"%s\"", payload.Message, payload.Hostname)

		w.WriteHeader(http.StatusOK)

	})

	log.Println("Starting web server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
