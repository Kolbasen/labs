package health

import (
	"encoding/json"
	"net/http"
	"time"
)

type healthResponse struct {
	Time string `json:"time"`
}

// Revert commit

// GetTime - simple func to get current time in RFC3339 format
func GetTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	response := healthResponse{Time: currentTime}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
