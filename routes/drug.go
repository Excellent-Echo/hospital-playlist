package routes

import (
	"hospital-playlist/drugs"
	"hospital-playlist/handler"

	"github.com/gin-gonic/gin"
)

var (
	drugRepository = drugs.NewRepository(DB)
	drugService    = drugs.NewService(drugRepository)
	drugHandler    = handler.NewDrugHandler(drugService, authService)
)

func DrugRoute(r *gin.Engine) {
	r.GET("/drugs", drugHandler.GetAllDrugsHandler)
	r.GET("/drug/:drug_id", drugHandler.GetDrugByIDHandler)
	r.POST("/drug", drugHandler.SaveNewDrugHandler)
}
