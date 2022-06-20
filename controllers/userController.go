package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vikas-gautam/golang-jwt-project/database"
	"go.mongodb.org/mongo-driver/mongo"
	helper "github.com/vikas-gautam/golang-jwt-project/helpers"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Signup() {}

func Login() {}

func HashPassword() {}

func VerifyPassword() {}

func GetUsers() {}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Params("user_Id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
