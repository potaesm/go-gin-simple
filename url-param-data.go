package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
	x, _ := strconv.ParseFloat(c.Param("x"), 64)
	y, _ := strconv.ParseFloat(c.Param("y"), 64)
	c.String(200, fmt.Sprintf("%f", x+y))
}

func main() {
	router := gin.Default()
	router.GET("/add/:x/:y", add)
	router.Run(":5000")
}
