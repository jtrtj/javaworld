package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jtrtj/javaworld/models"
	"github.com/jtrtj/javaworld/services"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()


	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api_test", func(c *gin.Context) {
		base, err := url.Parse("https://maps.googleapis.com/maps/api/place/")

		if err != nil {
			log.Fatalln(err)
		}

		// Path params
		base.Path += "nearbysearch/json"

		// Query params
		params := url.Values{}
		params.Add("key", os.Getenv("GOOGLE_KEY"))
		params.Add("location", "39.7392,-104.9903" )
		params.Add("radius", "33000" )
		params.Add("type", "cafe" )
		base.RawQuery = params.Encode()

		resp, err := http.Get(base.String())

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		bodyString := string(bodyBytes)

		var placesStruct services.PlacesResponse
		json.Unmarshal(bodyBytes, &placesStruct)

		c.JSON(200, gin.H{
			"string_version": bodyString,
			"json_version": placesStruct,
		})
	})

	router.GET("/cafes", func(c *gin.Context) {
		var cafes []models.Cafe
		models.DB.Find(&cafes)

		c.JSON(http.StatusOK, gin.H{"data": cafes})
	})

	router.Run()
}
