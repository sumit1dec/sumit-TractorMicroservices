package services

import (
	"Tractor_Details/app/models"
	"fmt"
)

var (
	//TripService exports ticket Service
	TractorService tractorServiceInterface = &tractorService{}
)

type tractorService struct{}

type tractorServiceInterface interface {
	CreateTractor(tractor models.Tractor) *models.Tractor
	GetTractor(tractorID uint) *models.Tractor
}

func (s *tractorService) CreateTractor(tractor models.Tractor) *models.Tractor {
	if err := tractor.Save(); err != nil {
		fmt.Println("Error is Save :", err)
		return nil
	}
	return &tractor
}

func (s *tractorService) GetTractor(TractorID uint) *models.Tractor {
	if TractorID >= 0 {
		tractor := models.Tractor{}
		tractor.ID = TractorID

		if err := tractor.Get(); err != nil {
			fmt.Println("Error is GET :", err)
			return nil
		}
		return &tractor
	}
	return nil
}
