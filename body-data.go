package main

import (
	"github.com/gin-gonic/gin"
)

type AddParams struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func add(c *gin.Context) {
	var ap AddParams
	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(400, gin.H{"error": "Calculator error"})
		return
	}

	c.JSON(200, gin.H{"answer": ap.X + ap.Y})
}

func main() {
	router := gin.Default()
	router.POST("/add", add)
	router.Run(":5000")
}
