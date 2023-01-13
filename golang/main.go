package main

import (
	"github.com/karasawa/gorm-crud.git/app/models"
	"github.com/karasawa/gorm-crud.git/app/controllers"
    "github.com/karasawa/gorm-crud.git/utils"
    "net/http"
	"os"
	"io"

    "github.com/gin-gonic/gin"
)

func init() {
	utils.LogginSettings("exp.log")
	models.DbInit()
}

func main() {
  f, _ := os.Create("apl.log")
  gin.DefaultWriter = io.MultiWriter(f)

  r := gin.Default()
  r.LoadHTMLGlob("app/views/templates/*")
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
      "message": "pong",
    })
  })
  r.POST("/create", controllers.PostCreate)
  r.Run(":8080")
}