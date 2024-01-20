package main 

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()

	database.StartDB()
	router.Use(gin.Logger())

	router.GET("/api", func(c *gin.Context){
		c.JSON(200, gin.H{
			"success": "Welcome to movies API"})
	})

	router.Run(":" + port)
}