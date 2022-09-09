package main

import (
	"github.com/gin-gonic/gin"

	"alpha-indo-soft/controllers"

	"alpha-indo-soft/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/articles", controllers.FindArticles)
	r.GET("/articles/:author", controllers.ArticleAuthorHandler)
	r.GET("/articles/query", controllers.ArticleQueryHandler)
	r.POST("/articles", controllers.CreateArticle)

	r.Run()
}