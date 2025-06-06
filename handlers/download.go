package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/syrshax/invoice-go-v2/internal"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 || pathParts[0] != "download" {
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

	defer internal.CleaningFiles(job.ID)

	fmt.Println("THIS IS THE JOB ID!!!")
	if job.EndPath == "" && (job.Status == internal.Pending || job.Status == internal.Processing) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "File not ready",
			"message": "Zip file is not yet available for this job",
		})
		return
	}

	if _, err := os.Stat(job.EndPath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "File not found",
			"message": "Zip file no longer exists on server",
		})
		return
	}

	file, err := os.Open(job.EndPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Server error",
			"message": "Could not open file for download",
		})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Server error",
			"message": "Could not get file information",
		})
		return
	}

	fileName := filepath.Base(job.EndPath)
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Length", string(rune(fileInfo.Size())))

	http.ServeContent(w, r, fileName, fileInfo.ModTime(), file)

	fmt.Println("THIS IS THE jobENDPATH to delete the ZIP", job.EndPath+".zip")
	defer func() {
		filedir := filepath.Join(job.EndPath, jobID+".zip")
		err := os.Remove(filedir)
		if err != nil {
			log.Printf("Couldnt delete file: %v \n", filedir)
		}
		fmt.Println("DELETED ======++++++++++++++++++++++++++++++++++++++=================")
	}()
}
