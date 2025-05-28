package services

import (
	"jobparser/models"
	"strings"
)

var Keywords = []string{
	"golang", "kubernetes", "linux", "kafka", "prometheus", "terraform",
	"sre", "devops", "cloud", "aws", "gcp", "azure",
}

func MatchesKeywords(job models.Job) bool {
	text := strings.ToLower(job.Title + " " + job.Description)

	for _, keyword := range Keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}

	return false
}

func FilterJobs(jobs []models.Job) []models.Job {
	var result []models.Job
	for _, job := range jobs {
		if MatchesKeywords(job) {
			result = append(result, job)
		}
	}
	return result
}
