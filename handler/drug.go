package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/drugs"
	"hospital-playlist/entity"
	"hospital-playlist/helper"

	"github.com/gin-gonic/gin"
)

type drugHandler struct {
	drugService drugs.Service
	authService auth.Service
}

func NewDrugHandler(drugService drugs.Service, authService auth.Service) *drugHandler {
	return &drugHandler{drugService, authService}
}

func (h *drugHandler) GetAllDrugsHandler(c *gin.Context) {
	drugs, err := h.drugService.GetAllDrugs()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all drugs", 200, "success", drugs)
	c.JSON(200, response)
}

func (h *drugHandler) GetDrugByIDHandler(c *gin.Context) {
	id := c.Params.ByName("drug_id")

	drug, err := h.drugService.GetSpesialistByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request drug ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", drug)
	c.JSON(200, response)
}

func (h *drugHandler) SaveNewDrugHandler(c *gin.Context) {
	var inputDrug entity.DrugInput

	if err := c.ShouldBindJSON(&inputDrug); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newSpesialis, err := h.drugService.SaveNewDrug(inputDrug)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Spesialist", 201, "status Created", newSpesialis)
	c.JSON(201, response)
}
