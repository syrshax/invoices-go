package internal

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type StatusType string

const (
	Pending    StatusType = "pending"
	Processing StatusType = "processing"
	Canceled   StatusType = "canceled"
	Finished   StatusType = "finished"
)

type Job struct {
	ID        string     `json:"jobId"`
	Status    StatusType `json:"status"`
	MsgError  string     `json:"msgError"`
	EndPath   string     `json:"endPath"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

var JobStorage sync.Map

func CreateJob() Job {
	job := Job{
		ID:        uuid.New().String(),
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	JobStorage.Store(job.ID, job)
	return job
}

func GetJob(jobID string) (Job, bool) {
	if value, ok := JobStorage.Load(jobID); ok {
		return value.(Job), true
	}
	return Job{}, false
}

func UpdateJobStatus(jobID string, status StatusType, errorMsg string) {
	if value, ok := JobStorage.Load(jobID); ok {
		job := value.(Job)
		job.Status = status
		job.MsgError = errorMsg
		job.UpdatedAt = time.Now()
		JobStorage.Store(jobID, job)
	}
}

func UpdateJobPath(jobID string, filePath string) {
	if value, ok := JobStorage.Load(jobID); ok {
		job := value.(Job)
		job.EndPath = filePath
		job.UpdatedAt = time.Now()
		JobStorage.Store(jobID, job)
	}
}
