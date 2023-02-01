package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/vendor", "./static/vendor") // Add ready-to-use compiled code for Bootstrap
	r.LoadHTMLGlob("template/**/**")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gin",
		})
	})
	log.Println("Server started")
	r.Run() // Default port 8080
}
