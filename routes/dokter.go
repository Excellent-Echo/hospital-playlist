package routes

import (
	"hospital-playlist/dokter"
	"hospital-playlist/handler"

	"github.com/gin-gonic/gin"
)

var (
	dokterRepository = dokter.NewRepository(DB)
	dokterService    = dokter.NewService(dokterRepository)
	dokterHandler    = handler.NewDokterHandler(dokterService, authService)
)

func DokterRoute(r *gin.Engine) {
	r.GET("/dokters", dokterHandler.GetAllDokterHandler)
	r.GET("/dokter/:dokter_id", dokterHandler.GetDokterHandlerByID)
	r.POST("/dokter/register", dokterHandler.CreateDokterHandler)
	r.POST("/dokter/login", dokterHandler.LoginDokterHandler)
}
