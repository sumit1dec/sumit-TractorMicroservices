package models

import (
	postgresql "Tractor_Details/db"

	"gorm.io/gorm"
)

type Tractor struct {
	gorm.Model
	ID                uint   `json:"id"  gorm:"autoIncrement"`
	TractorID         uint   `json:"tractor_id" gorm:"primaryKey"`
	YearOfManufacture uint   `json:"year_of_manufacture"`
	Color             string `json:"color"`
}

func init() {
	postgresql.DB.AutoMigrate(&Tractor{})
}

type Tractors []Tractor
