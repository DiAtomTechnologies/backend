package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vikram761/backend/models"
	"github.com/vikram761/backend/services"
)

type blogController struct {
	Blog services.BlogService
}

type BlogController interface {
	Save(*gin.Context)
	Delete(*gin.Context)
	GetAll(*gin.Context)
	GetOne(*gin.Context)
}

func NewBlogController(blog services.BlogService) BlogController {
	return &blogController{Blog: blog}
}

func (b *blogController) Save(ctx *gin.Context) {
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}
	user, _ := ctx.Get("user")
	userModel, _ := user.(models.User)
	blog.Author = userModel.Name

	if err := models.ValidateBlog(blog); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	if err := b.Blog.Save(blog); err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "Blog added successfully",
	})
	return
}

func (b *blogController) Delete(ctx *gin.Context) {
	var id string = ctx.Param("id")

	err := b.Blog.Delete(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":   "failed",
			"response": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status":   "success",
		"response": "Row deleted successfully",
	})
}

func (b *blogController) GetAll(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	limit := ctx.Query("limit")
	id := ctx.Query("id")
	tag := ctx.Query("tag")

	result, err := b.Blog.GetAll(limit, id, tag)

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

func (b *blogController) GetOne(ctx *gin.Context) {
	var id string = ctx.Param("id")

	result, err := b.Blog.GetOne(id)
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
