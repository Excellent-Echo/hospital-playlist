package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"hospital-playlist/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// ShowUserHandler for handing show all user in database from route "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "success", users)
	c.JSON(200, response)
}

// CreateUserHandler for handing if user / external create new user from route "/users"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var createUser entity.CreateUser

	if err := c.ShouldBindJSON(&createUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(createUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new User", 201, "status Created", newUser)
	c.JSON(201, response)
}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userLoginID := c.MustGet("currentUser").(int)

	if userID, _ := strconv.Atoi(id); userID != userLoginID {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": "this user does not have the authority"})

		c.JSON(400, responseError)
		return
	}

	user, err := h.userService.GetUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) GetUserByRoleDocterHandler(c *gin.Context) {
	role := c.Params.ByName("role")
	if role == "" {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": "this user does not have the authority"})

		c.JSON(400, responseError)
		return
	}

	dockter, err := h.userService.GetUserByRoleDocter(role)

	if err != nil {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", dockter)
	c.JSON(200, response)

}

func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userLoginID := c.MustGet("currentUser").(int)

	if userID, _ := strconv.Atoi(id); userID != userLoginID {
		responseError := helper.APIResponse("error bad request user ID", 400, "error", gin.H{"error": "this user does not have the authority"})

		c.JSON(400, responseError)
		return
	}

	user, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updatePasienInput entity.UpdateUser

	if err := c.ShouldBindJSON(&updatePasienInput); err != nil {
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

	user, err := h.userService.UpdateUserByID(id, updatePasienInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", 200, "success", user)
	c.JSON(200, response)
}

// login = menangkap email dan password yang dikirim oleh user (POST)
// mengecek apakah email ada di database( service ke repository)
// pengecekan apakah password di database sama dengan password yang dikirim (bcrypt)
// kita menggunakan generate token ke handler (response)
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var loginUser entity.LoginUser

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(loginUser)

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
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token, "role": userData.Role})
	c.JSON(200, response)
}
