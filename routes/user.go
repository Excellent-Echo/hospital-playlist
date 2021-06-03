package routes

import (
	"hospital-playlist/auth"
	"hospital-playlist/config"
	"hospital-playlist/handler"
	"hospital-playlist/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	r.POST("/user/register", userHandler.CreateUserHandler)
	r.POST("/user/login", userHandler.LoginUserHandler)
	r.GET("/user/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
	r.GET("/dokter/:role", userHandler.GetUserByRoleDocterHandler)
	r.DELETE("/user/:user_id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
	r.PUT("/user/:user_id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
}
