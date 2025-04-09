package models

import "gorm.io/gorm"

type Sneaker struct {
	gorm.Model
	Brand  string  `json:"brand"`
	Models string  `json:"models"`
	Size   int     `json:"size"`
	Price  float64 `json:"price"`
}
