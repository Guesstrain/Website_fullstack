package main

import (
	"my-personal-website/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	router.Static("/static", "./static/")

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controllers.HomePage)
	router.GET("/blog", controllers.BlogPage)
	router.DELETE("/article/:id", controllers.DeleteArticle)
	router.GET("/article/:id", controllers.ShowArticle)
	router.POST("/upload", controllers.UploadImage)
	router.POST("/edit", controllers.SubmitEditForm)
	router.GET("/edit", controllers.ShowEditForm)

	router.Run(":8080")
}
