package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Deepakpmk007/movie-stream-go/database"
	"github.com/Deepakpmk007/movie-stream-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetMovies() gin.HandlerFunc {
	movieCollection := database.Collection("movies")
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies []models.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Faild to fecth data..",
			})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Faild to fecth data..",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   movies,
		})
	}
}

func GetMovie() gin.HandlerFunc {
	movieCollection := database.Collection("movies")
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		movieID := c.Param("imdb_id")

		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Movies Id is required",
			})
			return
		}

		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Movie not found",
			})
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}
