package main

import (
	"fmt"
	"jobparser/parsers"
	"jobparser/services"
)

func main() {
	fmt.Println(" Parcing position  RemoteOK...")
	jobs := parsers.ParseRemoteOK()
	fmt.Printf("find %d vacancies\n", len(jobs))

	services.SaveToJSON(jobs, "data/jobs.json")
	fmt.Println("Preserve in  data/jobs.json")

	filtered := services.FilterJobs(jobs)
	fmt.Printf("Filtered by key skills of your profile: %d vacancies\n", len(filtered))

	services.SaveToJSON(filtered, "data/filtered_jobs.json")
	fmt.Println("Preserve filtered in data/filtered_jobs.json")
}
