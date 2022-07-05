package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/vikas-gautam/golang-jwt-project/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})

	})

	router.Run(":" + port)
}
