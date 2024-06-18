package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/vikram761/backend/models"
)

type careerService struct {
	Db *sql.DB
}

type CareerService interface {
	Save(models.Career) error
	GetCareer(id string) (models.Career, error)
	GetAllCareers() ([]models.Career, error)
	DeleteCareer(id string) error
}

func NewCareerService(db *sql.DB) CareerService {
	return &careerService{Db: db}
}

func (c *careerService) Save(career models.Career) error {
	var query string
	var args []interface{}
    fmt.Println(career)
	switch career.WorkType {
	case "job":
		query = "INSERT INTO CAREERS(TITLE,LOCATION,WORKTYPE, DESCRIPTION, START_DATE, END_DATE, APPLICATION_TIME) VALUES ($1, $2, $3, $4, $5, $6, $7)"
		args = []interface{}{career.Title, strings.ToLower(career.Location), strings.ToLower(career.WorkType), career.Description, career.StartDate, career.EndDate, career.ApplicationTime}
	case "internship", "event":
		query = "INSERT INTO CAREERS(TITLE, LOCATION, WORKTYPE, DESCRIPTION, DURATION, DURATIONTYPE, START_DATE, END_DATE, APPLICATION_TIME) VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9)"
		args = []interface{}{career.Title, strings.ToLower(career.Location), strings.ToLower(career.WorkType), career.Description, career.Duration, career.DurationType, career.StartDate, career.EndDate, career.ApplicationTime}
	default:
		return fmt.Errorf("Invalid worktype: %s", career.WorkType)
	}

	_, err := c.Db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *careerService) GetAllCareers() ([]models.Career, error) {
	query, err := c.Db.Query("SELECT ID, TITLE, LOCATION, WORKTYPE, DURATION, DURATIONTYPE, START_DATE, END_DATE, APPLICATION_TIME FROM CAREERS")
	if err != nil {
		return nil, err
	}
	var result []models.Career
	defer query.Close()
	for query.Next() {
		var career models.Career
		var duration sql.NullString
		var durationType sql.NullString
		err := query.Scan(&career.ID, &career.Title, &career.Location, &career.WorkType, &duration, &durationType, &career.StartDate, &career.EndDate, &career.ApplicationTime)
		if err != nil {
			return nil, err
		}

		if duration.Valid {
			career.Duration = duration.String
		} else {
			career.Duration = "0"
		}
		if durationType.Valid {
			career.DurationType = durationType.String
		} else {
			career.DurationType = ""
		}

		result = append(result, career)
	}

	if err := query.Err(); err != nil {
		return nil, err
	} 
     if len(result) == 0 {
      return []models.Career{} , nil
    }
	return result, nil

}

func (c *careerService) GetCareer(id string) (models.Career, error) {
	query := c.Db.QueryRow("SELECT * FROM CAREERS WHERE ID = $1", id)

	var career models.Career
	var duration sql.NullString
	var durationType sql.NullString
	err := query.Scan(&career.ID, &career.Title, &career.Location, &career.WorkType, &career.Description, &duration, &durationType, &career.StartDate, &career.EndDate, &career.ApplicationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Career{}, fmt.Errorf("user with ID %s not found", id)
		}
		return models.Career{}, err
	}

     if duration.Valid {
		career.Duration = duration.String
	}else{
      career.Duration = "0"
    }
	if durationType.Valid {
		career.DurationType = durationType.String
	}
	return career, nil
}

func (c* careerService) DeleteCareer(id string) error {
  _, err := c.Db.Exec("DELETE FROM CAREERS WHERE ID = $1",id);
  if err != nil {
    return err
  }
  return nil;
}

