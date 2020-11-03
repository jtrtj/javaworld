package models

type Cafe struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Address string `json:"address"`
	GoogleRating float64 `json:"google_rating"`
	Rating int `json:"rating"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Verified bool `json:"verified"`
}

type CafeCreateInput struct {
	Name string `json:"name"`
	Address string `json:"vicinity"`
}

type CafeUpdateInput struct {
	Name string `json:"name"`
	Address string `json:"vicinity"`
}

