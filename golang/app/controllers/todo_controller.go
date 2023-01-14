package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/karasawa/gorm-crud.git/app/models"
)

func TodoGet(ctx *gin.Context) {
	db := models.DbInit()

	todos := []models.Todo{}
	result := db.Find(&todos)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	log.Println("aaa")
	log.Println(result)
	log.Println(todos)
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"todos": todos,
	})
}

func TodoCreate(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "create_todo.html", nil)
		return
	}
	db := models.DbInit()

	postTask := ctx.PostForm("task")
	todo := models.Todo{Task: postTask}
	result := db.Create(&todo)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	ctx.Redirect(http.StatusFound, "/")
}

func TodoDelete(ctx *gin.Context) {
	db := models.DbInit()

	todo := models.Todo{}

	id := ctx.Param("ID")
	result := db.Where("id = ?", id).Delete(&todo)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	ctx.Redirect(http.StatusFound, "/")
}

func TodoUpdate(ctx *gin.Context) {
	ID := ctx.Param("ID")

	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "update_todo.html", gin.H{
			"ID": ID,
		})
		return
	}
	db := models.DbInit()

	postTask := ctx.PostForm("task")
	result := db.Where("id = ?", ID).Updates(models.Todo{Task: postTask})
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	ctx.Redirect(http.StatusFound, "/")
}