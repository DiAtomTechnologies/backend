package services

import (
	"database/sql"
	"fmt"
	"time"

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
	query := j.Db.QueryRow("SELECT END_DATE FROM CAREERS WHERE ID = $1", application.JobId)

	var endDate time.Time
	err := query.Scan(&endDate)
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

	stmt := "INSERT INTO JOB_APPLICATIONS(FIRST_NAME, LAST_NAME, EMAIL, PHONE, ADDRESS, WORK_EXPERIENCE, JOB_ID, NOTES) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = j.Db.Exec(stmt, application.FirstName, application.LastName, application.Email, application.Phone, application.Address, application.WorkExperience, application.JobId, application.Notes)
	if err != nil {
		return err
	}
	return nil
}
