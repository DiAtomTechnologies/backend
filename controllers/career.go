package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikram761/backend/models"
	"github.com/vikram761/backend/services"
)

type careerController struct {
	Career services.CareerService
}

type CareerController interface {
	Save(ctx *gin.Context)
	GetAllCareers(ctx *gin.Context)
	GetCareer(ctx *gin.Context)
	DeleteCareer(ctx *gin.Context)
}

func NewCareerController(career services.CareerService) CareerController {
	return &careerController{Career: career}
}

func (c *careerController) Save(ctx *gin.Context) {
	var career models.Career
	career.StartDate = time.Now()
	if err := ctx.BindJSON(&career); err != nil {
		ctx.JSON(400, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	career.EndDate = career.StartDate.AddDate(0, 0, career.ApplicationTime)

    err := c.Career.Save(career)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "Data added successfully",
	})
}

func (c *careerController) GetAllCareers(ctx *gin.Context) {
	result, err := c.Career.GetAllCareers()
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "failed",
			"reponse": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": result,
	})
}

func (c *careerController) GetCareer(ctx *gin.Context) {
	var id string = ctx.Param("id")
	result, err := c.Career.GetCareer(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": result,
	})
}

func (c *careerController) DeleteCareer(ctx *gin.Context) {
	var id string = ctx.Param("id")
	err := c.Career.DeleteCareer(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "Record was deleted sucessfully.",
	})
}
