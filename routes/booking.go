package routes

import (
	"hospital-playlist/booking"
	"hospital-playlist/handler"

	"github.com/gin-gonic/gin"
)

var (
	bookRepository = booking.NewRepository(DB)
	bookService    = booking.NewService(bookRepository)
	bookHandler    = handler.NewBookingHandler(bookService, authService)
)

func BookingRoute(r *gin.Engine) {
	r.GET("/booking", bookHandler.GetAllBookHandler)
	r.GET("/booking/:book_id", bookHandler.GetBookingByID)
	r.POST("/booking", bookHandler.SaveNewBookingHandler)
}
