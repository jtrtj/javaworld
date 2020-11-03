package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jtrtj/javaworld/models"
	"net/http"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()
	// comment the line below to seed/reseed database. Need to set up as a daily task in the future.
	// models.SeedDatabase()


	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/cafes", func(c *gin.Context) {
		var cafes []models.Cafe
		models.DB.Find(&cafes)

		c.JSON(http.StatusOK, gin.H{"data": cafes})
	})

	router.Run()
}
