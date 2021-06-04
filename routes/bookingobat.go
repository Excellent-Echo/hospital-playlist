package routes

import (
	"hospital-playlist/bookingobat"
	"hospital-playlist/handler"

	"github.com/gin-gonic/gin"
)

var (
	bookObatRepository = bookingobat.NewRepository(DB)
	bookObatService    = bookingobat.NewService(bookObatRepository)
	bookObatHandler    = handler.NewBookingObatHandler(bookObatService, authService)
)

func BookingObatRoute(r *gin.Engine) {
	r.GET("/bookingobat/:booking_id/:drug_id", bookObatHandler.GetBookObatByIDHandler)
}
