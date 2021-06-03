package routes

import (
	"hospital-playlist/handler"
	"hospital-playlist/spesialist"

	"github.com/gin-gonic/gin"
)

var (
	spesialistRepository = spesialist.NewRepository(DB)
	spesialistService    = spesialist.NewService(spesialistRepository)
	spesialistHandler    = handler.NewSpesialistHandler(spesialistService, authService)
)

func SpesialistRoute(r *gin.Engine) {
	r.GET("/spesialist", spesialistHandler.GetAllSpesialistHandler)
	r.GET("/spesialist:spesialist_id", spesialistHandler.GetSpesialistByIDHandler)
	r.POST("/spesialist", spesialistHandler.SaveNewSpesialistHandler)
	r.PUT("/spesialist/:spesialist_id", spesialistHandler.UpdateSpesialistByIDHandler)
}
