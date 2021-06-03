package routes

import (
	"hospital-playlist/handler"
	"hospital-playlist/userDetail"

	"github.com/gin-gonic/gin"
)

var (
	userDetailRepository = userDetail.NewRepository(DB)
	userDetailService    = userDetail.NewService(userDetailRepository)
	userDetailHandler    = handler.NewUserDetailHandler(userDetailService, authService)
)

func UserDetailRoute(r *gin.Engine) {
	r.GET("/user_details", handler.Middleware(userService, authService), userDetailHandler.GetUserDetailByUserIDHandler)
	r.POST("/user_details", handler.Middleware(userService, authService), userDetailHandler.SaveNewUserDetailHandler)
	r.PUT("/user_details", handler.Middleware(userService, authService), userDetailHandler.UpdateUserDetailByIDHandler)
}
