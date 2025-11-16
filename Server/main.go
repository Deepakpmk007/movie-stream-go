package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Deepakpmk007/movie-stream-go/controllers"
	"github.com/Deepakpmk007/movie-stream-go/database"
)

func main() {
	database.InitDB()
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Deepak")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movies/:imdb_id", controllers.GetMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
