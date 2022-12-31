package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Schedule struct {
	Name       string
	Book       Book
	PickUpDate time.Time
}

type ScheduleRequest struct {
	Name          string   `json:"name"`
	Title         string   `json:"title"`
	Author        []Author `json:"authors"`
	EditionNumber int32    `json:"edition_count"`
}

var schedules []Schedule

func BookSchedule(w http.ResponseWriter, r *http.Request) {
	var request ScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	book := Book{
		Title:         request.Title,
		Author:        request.Author,
		EditionNumber: request.EditionNumber,
	}

	schedule := Schedule{
		Name:       request.Name,
		Book:       book,
		PickUpDate: time.Now(),
	}

	schedules = append(schedules, schedule)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Book success"))
}

func GetAllSchedule(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(schedules)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
