package models

type Cafe struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Rating int `json:"rating"`
}

type CafeCreateInput struct {
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}

type CafeUpdateInput struct {
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}
