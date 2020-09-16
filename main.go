package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jtrtj/javaworld/models"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()


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
