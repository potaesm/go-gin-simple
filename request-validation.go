package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PrintJob struct {
	JobId int `json:"jobId" binding:"required,gte=10000"`
	Pages int `json:"pages" binding:"required,gte=1,lte=100"`
}

func main() {
	router := gin.Default()
	router.POST("/print", func(c *gin.Context) {
		var p PrintJob
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!"})
			return
		}
		c.JSON(200, gin.H{"message": fmt.Sprintf("PrintJob #%v started!", p.JobId)})
	})
	router.Run(":5000")
}
