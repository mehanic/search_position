package models

type Job struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Link        string `json:"link"`
	Source      string `json:"source"`
	Description string `json:"description"`
}
