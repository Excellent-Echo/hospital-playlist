package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"hospital-playlist/medicine"
	"strconv"

	"github.com/gin-gonic/gin"
)

type medicineHandler struct {
	medicineService medicine.Service
	authService     auth.Service
}

func NewMedicineHandler(medicineService medicine.Service, authService auth.Service) *medicineHandler {
	return &medicineHandler{medicineService, authService}
}

func (h *medicineHandler) GetAllMedicineHandler(c *gin.Context) {
	medicines, err := h.medicineService.GetAllMedicine()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all room", 200, "success", medicines)
	c.JSON(200, response)
}

func (h *medicineHandler) GetMedicineByIDHandler(c *gin.Context) {
	id := c.Params.ByName("medicine_id")

	medicine, err := h.medicineService.GetMedicineByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request room ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get medicine by ID", 200, "success", medicine)
	c.JSON(200, response)
}

func (h *medicineHandler) SaveNewMedicineHandler(c *gin.Context) {
	var inputMedicine entity.MedicineInput

	if err := c.ShouldBindJSON(&inputMedicine); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newMedicine, err := h.medicineService.SaveNewMedicine(inputMedicine)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Medicine", 201, "status Created", newMedicine)
	c.JSON(201, response)
}

func (h *medicineHandler) UpdateMedicineByIDHandler(c *gin.Context) {
	id := c.Params.ByName("medicine_id")

	var updateMedicineInput entity.MedicineInput

	if err := c.ShouldBindJSON(&updateMedicineInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	UserData := int(c.MustGet("currentUser").(int))

	if idParam != UserData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.medicineService.UpdateMedicineByID(id, updateMedicineInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update medicine by ID", 200, "success", user)
	c.JSON(200, response)
}
