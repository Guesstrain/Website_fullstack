package controllers

import (
	"my-personal-website/database"
	"my-personal-website/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var databaseService database.DatabaseService = database.NewDatabaseService()

func HomePage(c *gin.Context) {
	overview := databaseService.All()

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Title":       overview.Title,
		"Description": overview.Description,
		"Interests":   overview.Interest,
		"Contact":     overview.Contact,
	})
}

func ShowEditForm(c *gin.Context) {
	info := databaseService.All()
	articles := databaseService.FindBlog()
	c.HTML(http.StatusOK, "edit.html", gin.H{"info": info, "articles": articles})
}

func SubmitEditForm(c *gin.Context) {
	var info models.PersonalInfo
	if err := c.ShouldBind(&info); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	info.ID = 1
	if err := databaseService.Update(info); err != nil {
		c.String(http.StatusInternalServerError, "Failed to update info")
		return
	}
	c.Redirect(http.StatusFound, "/")
}
