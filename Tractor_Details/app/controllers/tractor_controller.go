package controllers

import (
	"Tractor_Details/app/models"
	"Tractor_Details/app/services"
	"fmt"

	"github.com/revel/revel"
)

type TractorController struct {
	*revel.Controller
}

func (c TractorController) Create() revel.Result {
	var tractor models.Tractor

	if err := c.Params.BindJSON(&tractor); err != nil {
		fmt.Println(err)
	}

	result := services.TractorService.CreateTractor(tractor)

	// if result != nil {
	// 	return c.RenderJSON(saveErr)
	// }
	return c.RenderJSON(result)
}

func (c TractorController) Get(tractorID uint) revel.Result {
	result := services.TractorService.GetTractor(tractorID)

	// if getErr != nil {
	// 	c.RenderJSON(getErr)
	// }
	return c.RenderJSON(result)
}
