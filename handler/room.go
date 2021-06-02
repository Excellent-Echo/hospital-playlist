package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"hospital-playlist/room"
	"strconv"

	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	roomService room.Service
	authService auth.Service
}

func NewRoomHandler(roomService room.Service, authService auth.Service) *roomHandler {
	return &roomHandler{roomService, authService}
}

func (h *roomHandler) GetAllRoomHandler(c *gin.Context) {
	rooms, err := h.roomService.GetAllRoom()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all room", 200, "success", rooms)
	c.JSON(200, response)
}

func (h *roomHandler) GetRoomByIDHandler(c *gin.Context) {
	id := c.Params.ByName("room_id")

	room, err := h.roomService.GetRoomByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request room ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", room)
	c.JSON(200, response)
}

func (h *roomHandler) SaveNewRoomHandler(c *gin.Context) {
	var inputRoom entity.RoomInput

	if err := c.ShouldBindJSON(&inputRoom); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newRoom, err := h.roomService.SaveNewRoom(inputRoom)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Room", 201, "status Created", newRoom)
	c.JSON(201, response)
}

func (h *roomHandler) UpdateRoomByIDHandler(c *gin.Context) {
	id := c.Params.ByName("room_id")

	var updateRoomInput entity.RoomInput

	if err := c.ShouldBindJSON(&updateRoomInput); err != nil {
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

	user, err := h.roomService.UpdateRoomByID(id, updateRoomInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update spesialist by ID", 200, "success", user)
	c.JSON(200, response)
}
