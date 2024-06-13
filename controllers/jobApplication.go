package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vikram761/backend/models"
	"github.com/vikram761/backend/services"
)

type applicationController struct {
	Application services.JobApplicationService
}

type ApplicationController interface{
	Save (*gin.Context)
}

func NewApplicationController(service services.JobApplicationService) ApplicationController {
  return &applicationController{Application: service}
} 

func (a *applicationController) Save(ctx *gin.Context) {
	var application models.JobApplication
	if err := ctx.BindJSON(&application); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return

	}
	if err := a.Application.Save(application); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "Applied for this post.",
	})
}
