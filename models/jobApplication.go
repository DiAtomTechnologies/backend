package models

import "time"

type JobApplication struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	WorkExperience int       `json:"work_experienc"`
	JobId          string    `json:"jobId"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
}
