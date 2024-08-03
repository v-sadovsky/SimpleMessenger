package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type HealthStatus struct {
	Status string `json:"status"`
}

var (
	isReady bool
	mu      sync.RWMutex
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleLiveness)
	http.HandleFunc("/ready", handleReadiness)

	// Simulate readiness after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		setReady(true)
	}()

	log.Println("Starting Friends manager service v1.0.0 on port :82")
	log.Fatal(http.ListenAndServe(":82", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Friends manager Service v1.0.0!\n"))
}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	response := HealthStatus{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	if !isReady {
		http.Error(w, "friends manager service is not ready", http.StatusServiceUnavailable)
		return
	}

	w.Write([]byte("Friends manager service is ready!"))
}

func setReady(ready bool) {
	mu.Lock()
	defer mu.Unlock()
	isReady = ready
}
