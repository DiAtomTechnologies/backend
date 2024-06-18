package models

import "time"

type Career struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	WorkType     string    `json:"worktype"`
	Description  string    `json:"description,omitempty"`
	Duration     string     `json:"duration,omitempty"`
	DurationType string    `json:"durationType,omitempty"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
    ApplicationTime string   `json:"applicationTime"`
}
