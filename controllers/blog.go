package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogPage(c *gin.Context) {
	articles := databaseService.FindBlog()
	c.HTML(http.StatusOK, "list.html", gin.H{
		"articles": articles,
	})
}
