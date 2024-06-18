package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/vikram761/backend/models"
)

type jobApplicationService struct {
	Db *sql.DB
}

type JobApplicationService interface {
	Save(models.JobApplication) error
}

func NewJobApplicationService(db *sql.DB) JobApplicationService {
	return &jobApplicationService{
		Db: db,
	}
}

func (j *jobApplicationService) Save(application models.JobApplication) error {
	query := j.Db.QueryRow("SELECT WORKTYPE, END_DATE FROM CAREERS WHERE ID = $1", application.JobId)

	var endDate time.Time
	var worktype string
	err := query.Scan(&worktype, &endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("career application with ID %s not found", application.JobId)
		}
		return err
	}

	currTime := time.Now()
	diff := endDate.Sub(currTime)
	days := int(diff.Hours() / 24)

	if days < 0 {
		return fmt.Errorf("The Application time has ended.")
	}

	var stmt string
	var args []interface{}

	if worktype == "event" {
		stmt = "INSERT INTO JOB_APPLICATIONS(FIRST_NAME, LAST_NAME, EMAIL, PHONE, ADDRESS, JOB_ID) VALUES ($1, $2, $3, $4, $5, $6)"
		args = []interface{}{application.FirstName, application.LastName, application.Email, application.Phone, application.Address, application.JobId}
	} else {
		stmt = "INSERT INTO JOB_APPLICATIONS(ID, FIRST_NAME, LAST_NAME, EMAIL, PHONE, ADDRESS, WORK_EXPERIENCE, JOB_ID, NOTES, SKILLS) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
		args = []interface{}{application.ID ,application.FirstName, application.LastName, application.Email, application.Phone, application.Address, application.WorkExperience, application.JobId, application.Notes, pq.Array(application.Skills)}
	}
	_, err = j.Db.Exec(stmt, args...)
	if err != nil {
		return err
	}
	return nil
}
