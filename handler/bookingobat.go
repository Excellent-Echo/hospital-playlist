package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/bookingobat"
	"hospital-playlist/helper"

	"github.com/gin-gonic/gin"
)

type bookingObatHandler struct {
	bookObatService bookingobat.Service
	authObatService auth.Service
}

func NewBookingObatHandler(bookObatService bookingobat.Service, authObatService auth.Service) *bookingObatHandler {
	return &bookingObatHandler{bookObatService, authObatService}
}

func (h *bookingObatHandler) GetBookObatByIDHandler(c *gin.Context) {
	idBoking := c.Params.ByName("booking_id")
	idDrug := c.Params.ByName("drug_id")

	books, err := h.bookObatService.GetBookingByID(idBoking, idDrug)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all booking", 200, "success", books)
	c.JSON(200, response)
}
