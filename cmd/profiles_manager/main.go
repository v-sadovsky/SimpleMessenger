package profiles_manager

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
	http.HandleFunc("/health", handleLiveness)
	http.HandleFunc("/ready", handleReadiness)

	// Simulate readiness after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		setReady(true)
	}()

	log.Println("Starting Profiles manager service on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLiveness(w http.ResponseWriter, r *http.Request) {
	response := HealthStatus{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	if !isReady {
		http.Error(w, "profiles manager service is not ready", http.StatusServiceUnavailable)
		return
	}

	w.Write([]byte("Profiles manager service is ready!"))
}

func setReady(ready bool) {
	mu.Lock()
	defer mu.Unlock()
	isReady = ready
}
