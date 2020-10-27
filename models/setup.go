package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"github.com/jtrtj/javaworld/services"
	"time"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := fmt.Sprintf("host=%s user=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Cafe{})
	DB = db
}

func SeedDatabase() {
	api_call := services.Call()
	results := api_call.Results
	for _, input := range results {
		cafe := Cafe{Name: input.Name, Address: input.Vicinity}
		DB.Create(&cafe)
	}
	check(api_call)
}

func check(results services.PlacesResponse) {
	time.Sleep(3000 * time.Millisecond)
	if len(results.NextPageToken) > 0 {
		api_call := services.CallNextPage(results.NextPageToken)
		results := api_call.Results
		for _, input := range results {
			cafe := Cafe{Name: input.Name, Address: input.Vicinity}
			DB.Create(&cafe)
		}
		check(api_call)
	}
}
