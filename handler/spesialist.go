package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"hospital-playlist/spesialist"
	"strconv"

	"github.com/gin-gonic/gin"
)

type specialistHandler struct {
	specialistService spesialist.Service
	authService       auth.Service
}

func NewSpesialistHandler(specialistService spesialist.Service, auth auth.Service) *specialistHandler {
	return &specialistHandler{specialistService, auth}
}

func (h *specialistHandler) GetAllSpesialistHandler(c *gin.Context) {
	specialists, err := h.specialistService.GetAllSpesialist()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "success", specialists)
	c.JSON(200, response)
}

func (h *specialistHandler) GetSpesialistByIDHandler(c *gin.Context) {
	id := c.Params.ByName("spesialist_id")

	spesialist, err := h.specialistService.GetSpesialistByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request specialist ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", spesialist)
	c.JSON(200, response)
}

func (h *specialistHandler) SaveNewSpesialistHandler(c *gin.Context) {
	var inputSpesialis entity.SpecialistInput

	if err := c.ShouldBindJSON(&inputSpesialis); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newSpesialis, err := h.specialistService.SaveNewSpesialist(inputSpesialis)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Spesialist", 201, "status Created", newSpesialis)
	c.JSON(201, response)
}

func (h *specialistHandler) UpdateSpesialistByIDHandler(c *gin.Context) {
	id := c.Params.ByName("spesialist_id")

	var updateSpesialistInput entity.SpecialistInput

	if err := c.ShouldBindJSON(&updateSpesialistInput); err != nil {
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

	user, err := h.specialistService.UpdateSpesialistByID(id, updateSpesialistInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update spesialist by ID", 200, "success", user)
	c.JSON(200, response)
}
