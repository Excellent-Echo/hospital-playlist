package routes

import (
	"hospital-playlist/drug"
	"hospital-playlist/handler"

	"github.com/gin-gonic/gin"
)

var (
	drugRepository = drug.NewRepository(DB)
	drugService    = drug.NewService(drugRepository)
	drugHandler    = handler.NewDrugHandler(drugService, authService)
)

func DrugRoute(r *gin.Engine) {
	r.GET("/drugs", drugHandler.GetAllDrugHandler)
	r.GET("/drug/:drug_id", drugHandler.GetDrugByIDHandler)
	r.POST("/drug", drugHandler.SaveNewDrugHandler)
	r.PUT("/drug/:drug_id", drugHandler.UpdateDrugByIDHandler)
}
