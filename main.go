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

    blogService services.BlogService = services.NewBlogService(database)
    blogController controllers.BlogController = controllers.NewBlogController(blogService)
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

    //career routes
    router.POST("/career", careerController.Save)
    router.GET("/careers", careerController.GetAllCareers)
    router.GET("/career/:id", careerController.GetCareer)
    router.DELETE("/career/:id", careerController.DeleteCareer)


    //blog routes
    router.POST("/blog", blogController.Save)
    router.DELETE("/blog/:id",blogController.Delete)
    router.GET("/blogs", blogController.GetAll)
    router.GET("/blog/:id", blogController.GetOne)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error occured", err)
	}
}
