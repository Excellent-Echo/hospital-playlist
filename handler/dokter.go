package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/dokter"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dokterHandler struct {
	dokterService dokter.Service
	authService   auth.Service
}

func NewDokterHandler(dokterService dokter.Service, authService auth.Service) *dokterHandler {
	return &dokterHandler{dokterService, authService}
}

func (h *dokterHandler) GetAllDokterHandler(c *gin.Context) {
	dokters, err := h.dokterService.GetAllDokter()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "success", dokters)
	c.JSON(200, response)
}

func (h *dokterHandler) GetDokterHandlerByID(c *gin.Context) {
	id := c.Params.ByName("dokter_id")

	spesialist, err := h.dokterService.GetDokterByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request dokter ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get dokter by ID", 200, "success", spesialist)
	c.JSON(200, response)
}

func (h *dokterHandler) LoginDokterHandler(c *gin.Context) {
	var loginDokter entity.LoginDokter

	if err := c.ShouldBindJSON(&loginDokter); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.dokterService.LoginDokter(loginDokter)

	if err != nil {
		responseError := helper.APIResponse("input data error", 401, "error", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err})

		c.JSON(401, responseError)
		return
	}
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token})
	c.JSON(200, response)
}

func (h *dokterHandler) CreateDokterHandler(c *gin.Context) {
	// var userData int

	// dokterData, _ := strconv.Itoa(userData)

	var createDokter entity.CreateDokter

	if err := c.ShouldBindJSON(&createDokter); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.dokterService.SaveNewDokter(createDokter)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new User", 201, "status Created", newUser)
	c.JSON(201, response)
}

func (h *userHandler) UpdateDokterByIDHandler(c *gin.Context) {
	id := c.Params.ByName("dokter_id")

	var updateDokterInput entity.UpdateDokter

	if err := c.ShouldBindJSON(&updateDokterInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	userData := int(c.MustGet("currentUser").(int)) // perlu diperbaiki

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.userService.UpdateUserByID(id, entity.UpdateUser(updateDokterInput))
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", 200, "success", user)
	c.JSON(200, response)
}
