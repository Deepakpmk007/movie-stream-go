package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Deepakpmk007/movie-stream-go/database"
	"github.com/Deepakpmk007/movie-stream-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(HashPassword), nil
}

func Registeruser() gin.HandlerFunc {
	userCollection := database.Collection("users")

	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()

		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

		hashPassword, err := HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to hash to password"})
		}
		user.Password = hashPassword

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Faild to check existiog user"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists!"})
			return
		}
		user.UserId = bson.NewObjectID().Hex()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		result, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to create"})
			return
		}

		c.JSON(http.StatusCreated, result)

	}
}

func LoginUser() gin.HandlerFunc {
	userCollection := database.Collection("users")
	return func(c *gin.Context) {
		var userLogin models.UserLogin

		if err := c.ShouldBindJSON(&userLogin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var foundUser models.User

		err := userCollection.FindOne(ctx, bson.M{"email": userLogin.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid email or password",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userLogin.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "pssword no match"})
			return
		}

	}
}
