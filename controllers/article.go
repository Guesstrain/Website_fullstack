package controllers

import (
	"my-personal-website/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nextID int

func ShowArticle(c *gin.Context) {

	id := c.Param("id")
	number, _ := strconv.Atoi(id)
	article := databaseService.ShowArticle(number)
	c.HTML(http.StatusOK, "detail.html", article)
}

func UploadImage(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := strconv.Itoa(nextID) + "_" + filepath.Base(file.Filename)
	filepath := filepath.Join("static", filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.String(http.StatusInternalServerError, "Failed to upload file")
		return
	}

	article := models.Article{
		ID:       nextID,
		Title:    title,
		Content:  content,
		ImageURL: "/" + filepath,
	}

	err = databaseService.AddArticle(article)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to add in db")
		return
	}
	nextID++

	c.Redirect(http.StatusFound, "/")
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	number, _ := strconv.Atoi(id)
	article := databaseService.ShowArticle(number)

	_ = databaseService.DeleteArticle(article.ID)

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
