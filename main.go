package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vikram761/backend/controllers"
	"github.com/vikram761/backend/db"
	"github.com/vikram761/backend/services"
)

var (
	database              *sql.DB
	careerService         services.CareerService
	careerController      controllers.CareerController
	blogService           services.BlogService
	blogController        controllers.BlogController
	authController        controllers.AuthController
	applicationService    services.JobApplicationService
	applicationController controllers.ApplicationController
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	username := os.Getenv("USERNAME")
	passwd := os.Getenv("PASSWD")
	hostDB := os.Getenv("HOSTDB")
	dbName := os.Getenv("DBNAME")
    port := os.Getenv("DBPORT")

	database = db.Connectdb(username, passwd, hostDB, dbName, port)
	careerService = services.NewCareerService(database)
	careerController = controllers.NewCareerController(careerService)
	blogService = services.NewBlogService(database)
	blogController = controllers.NewBlogController(blogService)
	authController = controllers.NewAuthController(database)
	applicationService = services.NewJobApplicationService(database)
	applicationController = controllers.NewApplicationController(applicationService)
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello world")
	})

	careerRoutes := router.Group("/career")
	{
		careerRoutes.POST("/", authController.CheckAuth, careerController.Save)
		careerRoutes.GET("/", careerController.GetAllCareers)
		careerRoutes.GET("/:id", careerController.GetCareer)
		careerRoutes.DELETE("/:id", authController.CheckAuth, careerController.DeleteCareer)
	}

	blogRoutes := router.Group("/blog")
	{
		blogRoutes.POST("/", authController.CheckAuth, blogController.Save)
		blogRoutes.GET("/", blogController.GetAll)
		blogRoutes.GET("/:id", blogController.GetOne)
		blogRoutes.DELETE("/:id", authController.CheckAuth, blogController.Delete)
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/validateUser", authController.CheckAuth, authController.ValidateUser)
        authRoutes.GET("/validate", authController.Validate)
	}

	router.POST("/application", applicationController.Save)

    appPort := os.Getenv("PORT")
    if appPort == "" {
      appPort = "8080"
    }

    if err := router.Run(":" + appPort); err != nil {
		log.Fatal("Error occurred", err)
	}
}
