package services

import (
    "encoding/json"
    "jobparser/models"
    "os"
    "log"
)

func SaveToJSON(jobs []models.Job, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatal("Trouble with creating file:", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(jobs)
    if err != nil {
        log.Fatal("JSON trouble with create JSON:", err)
    }
}
