package models

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDataBase() {
	dsn := "host=postgres user=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Cafe{})

	lula := Cafe{Name: "Lula Rose", 
							 Address: "123 Street Ave.",
							 City: "Denver",
							 State: "Colorado",
							 Zip: "12345",
							 Rating: 90,
							}
	db.Create(&lula)

	DB = db
}