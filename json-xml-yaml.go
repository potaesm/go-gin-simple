package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id   int    `json:"id" xml:"Id" yaml:"id"`
	Name string `json:"name" xml:"Name" yaml:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/productJSON", func(c *gin.Context) {
		product := Product{1, "Apple"}
		c.JSON(200, product)
	})

	router.GET("/productXML", func(c *gin.Context) {
		product := Product{2, "Banana"}
		c.XML(200, product)
	})
	router.GET("/productYAML", func(c *gin.Context) {
		product := Product{3, "Mango"}
		c.YAML(200, product)
	})
	router.Run(":5000")
}
