package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vikram761/backend/controllers"
	"github.com/vikram761/backend/db"
	"github.com/vikram761/backend/services"
)

var (
	USERNAME string  = os.Getenv("USERNAME")
	PASSWD   string  = os.Getenv("PASSWD")
	HOSTDB   string  = os.Getenv("HOSTDB")
	DBNAME   string  = os.Getenv("DBNAME")
	database *sql.DB = db.Connectdb(USERNAME, PASSWD, HOSTDB, DBNAME)
    careerService services.CareerService = services.NewCareerService(database);
    careerController controllers.CareerController = controllers.NewCareerController(careerService)
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello world")
	})
    router.POST("/career", careerController.Save)
    router.GET("/careers", careerController.GetAllCareers)
    router.GET("/career/:id", careerController.GetCareer)
    router.DELETE("/career/:id", careerController.DeleteCareer)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error occured", err)
	}
}
