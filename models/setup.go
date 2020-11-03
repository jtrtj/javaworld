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
		if ( input.Name == "Starbucks" || input.Name == "Peet's Coffee" ) {
			continue
		}
		cafe := Cafe{
								 Name: input.Name, 
								 Address: input.Vicinity,
								 GoogleRating: input.Rating,
								 Lat: input.Geometry.Location.Lat,
								 Lng: input.Geometry.Location.Lng,
								}
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
			if ( input.Name == "Starbucks" || input.Name == "Peet's Coffee" ) {
				continue
			}
			cafe := Cafe{
								 Name: input.Name, 
								 Address: input.Vicinity,
								 GoogleRating: input.Rating,
								 Lat: input.Geometry.Location.Lat,
								 Lng: input.Geometry.Location.Lng,
								}
			DB.Create(&cafe)
		}
		check(api_call)
	}
}


// {
//             "business_status": "OPERATIONAL",
//             "geometry": {
//                 "location": {
//                     "lat": 39.7718444,
//                     "lng": -105.0441417
//                 },
//                 "viewport": {
//                     "northeast": {
//                         "lat": 39.77319467989272,
//                         "lng": -105.0427253201073
//                     },
//                     "southwest": {
//                         "lat": 39.77049502010728,
//                         "lng": -105.0454249798927
//                     }
//                 }
//             },
//             "icon": "https://maps.gstatic.com/mapfiles/place_api/icons/v1/png_71/cafe-71.png",
//             "name": "Downpours Coffee",
//             "opening_hours": {
//                 "open_now": true
//             },
//             "photos": [
//                 {
//                     "height": 3036,
//                     "html_attributions": [
//                         "<a href=\"https://maps.google.com/maps/contrib/101175743725300481004\">Anna Arenz</a>"
//                     ],
//                     "photo_reference": "CmRaAAAA9IQd8ePhdFmog8r-cVzm6gkzpxwBHnUYd0BnNiZKIaIEzvt_q2grZ-kkrJI3VjR-aRhm5d6nAgEXqwCLmi_WZA2QfbNQxJrD6TSsgE_wmBE2tFcIEB54t6ABVsA7VNoKEhCfqZuBHDabW_PuVbMzDySNGhRc8UG8fMF7HRJ4HPoQTOyL9OeSYQ",
//                     "width": 4048
//                 }
//             ],
//             "place_id": "ChIJHXrrTJ2Ha4cRfXJ6JkGZoXM",
//             "plus_code": {
//                 "compound_code": "QXC4+P8 Denver, Colorado",
//                 "global_code": "85FPQXC4+P8"
//             },
//             "rating": 4.6,
//             "reference": "ChIJHXrrTJ2Ha4cRfXJ6JkGZoXM",
//             "scope": "GOOGLE",
//             "types": [
//                 "cafe",
//                 "food",
//                 "point_of_interest",
//                 "store",
//                 "establishment"
//             ],
//             "user_ratings_total": 173,
//             "vicinity": "3937 Tennyson St, Denver"
// },