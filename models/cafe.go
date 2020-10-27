package models

type Cafe struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Address string `json:"address"`
	Rating int `json:"rating"`
}

type CafeCreateInput struct {
	Name string `json:"name"`
	Address string `json:"vicinity"`
}

type CafeUpdateInput struct {
	Name string `json:"name"`
	Address string `json:"vicinity"`
}
