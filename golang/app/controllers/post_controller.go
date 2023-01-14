package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/karasawa/gorm-crud.git/app/models"
)

func PostCreate(c *gin.Context) {
	db := models.DbInit()

	PostContent := c.PostForm("content")
	content := models.Post{Content: PostContent}
	result := db.Create(&content)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": content,
	})
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}