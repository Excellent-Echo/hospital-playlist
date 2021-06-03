package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/drug"
	"hospital-playlist/entity"
	"hospital-playlist/helper"

	"github.com/gin-gonic/gin"
)

type drughandler struct {
	drugService drug.Service
	authService auth.Service
}

func NewDrugHandler(drugService drug.Service, authService auth.Service) *drughandler {
	return &drughandler{drugService, authService}
}

func (h *drughandler) GetAllDrugHandler(c *gin.Context) {
	drug, err := h.drugService.GetAllDrug()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all drug", 200, "success", drug)
	c.JSON(200, response)

}

func (h *drughandler) SaveNewDrugHandler(c *gin.Context) {

	var inputDrug entity.CreateDrug

	if err := c.ShouldBindJSON(&inputDrug); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newDrug, err := h.drugService.SaveNewDrug(inputDrug)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new drug", 201, "success", newDrug)
	c.JSON(201, response)
}

func (h *drughandler) UpdateDrugByIDHandler(c *gin.Context) {

	id := c.Params.ByName("drug_id")

	var updateInputDrug entity.UpdateDrug

	if err := c.ShouldBindJSON(&updateInputDrug); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	updateDrug, err := h.drugService.UpdateDrugByID(id, updateInputDrug)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update category by ID", 200, "success", updateDrug)
	c.JSON(200, response)
}
