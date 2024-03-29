package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/karasawa/gorm-crud.git/app/controllers"
	"github.com/karasawa/gorm-crud.git/app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	models.DbInit()
}

func main() {
  f, _ := os.OpenFile("apl.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  defer f.Close()
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

  r := gin.Default()
  r.Use(cors.New(cors.Config{
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
        "PUT",
        "DELETE",
    },
    AllowHeaders: []string{
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "X-CSRF-Token",
        "Authorization",
    },
    AllowOrigins: []string{
        "*",
    },
    MaxAge: 24 * time.Hour,
  }))
  r.LoadHTMLGlob("app/views/templates/*")
  r.GET("/", controllers.TodoGet)
  todo := r.Group("/todo")
	{
		todo.GET("/create", controllers.TodoCreate)
		todo.POST("/create", controllers.TodoCreate)
		todo.GET("/update/:ID", controllers.TodoUpdate)
		todo.POST("/update/:ID", controllers.TodoUpdate)
		todo.GET("/delete/:ID", controllers.TodoDelete)
	}
  r.GET("/api", func(ctx *gin.Context) {
    db := models.DbInit()

    todos := []models.Todo{}

    result := db.Find(&todos)
    if result.Error != nil {
      return
    }
    ctx.JSON(http.StatusOK, gin.H{
      "response": "wawawa",
      "todos": todos,
    })
  })
  r.Run(":8080")
}