package parsers

import (
	"jobparser/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseRemoteOK() []models.Job {
	url := "https://remoteok.com/remote-dev+golang-jobs"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Fail riquest:", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Fail parsing:", err)
	}

	var jobs []models.Job

	doc.Find("tr.job").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h2").Text())
		company := strings.TrimSpace(s.Find(".companyLink h3").Text())
		location := strings.TrimSpace(s.Find(".location").Text())

		link, exists := s.Attr("data-href")
		if !exists || title == "" {
			return
		}

		fullLink := "https://remoteok.com" + link

		// 👉 Загружаем описание страницы вакансии
		description := fetchJobDescription(fullLink)

		jobs = append(jobs, models.Job{
			Title:       title,
			Company:     company,
			Location:    location,
			Link:        fullLink,
			Source:      "RemoteOK",
			Description: description,
		})

		time.Sleep(500 * time.Millisecond) // защита от бана
	})

	return jobs
}

func fetchJobDescription(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Println("Fail dawnload description:", url, err)
		return ""
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Fail parsing description:", url, err)
		return ""
	}

	description := ""
	selectors := []string{".description", ".job-desc", ".markdown"}

	for _, sel := range selectors {
		text := strings.TrimSpace(doc.Find(sel).Text())
		if len(text) > len(description) {
			description = text
		}
	}

	return description
}
