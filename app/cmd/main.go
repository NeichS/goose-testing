package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8080"
	}

	err := router.Run(":" + port) // listen and serve on 0.0.0.0:8080

	if err != nil {
		panic(err)
	}
}
