package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Ok Deepak",
		})
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("server is not run on port 8000")
	}
}
