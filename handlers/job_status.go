package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/syrshax/invoice-go-v2/internal"
)

func JobStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Expected format: /jobs/{id}/status
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) < 2 || pathParts[0] != "jobs" {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	jobID := pathParts[1]

	job, ok := internal.GetJob(jobID)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Job not found",
			"message": "Job with this ID does not exist",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job) // Return the entire job object
}
