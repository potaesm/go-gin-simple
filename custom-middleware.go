package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.GetHeader("User-Agent"))
		// Before calling handler
		c.Next()
		// After calling handler
	}
}
func main() {
	router := gin.Default()
	router.Use(FindUserAgent())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Middleware works!"})
	})
	router.Run(":5000")
}
