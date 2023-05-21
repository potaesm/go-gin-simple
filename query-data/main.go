package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
	x, _ := strconv.ParseFloat(c.DefaultQuery("x", "0"), 64)
	y, _ := strconv.ParseFloat(c.DefaultQuery("y", "0"), 64)
	c.String(200, fmt.Sprintf("%f", x+y))
}

func main() {
	router := gin.Default()
	router.GET("/add", add)
	router.Run(":5000")
}
