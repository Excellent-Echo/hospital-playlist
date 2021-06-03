package handler

import (
	"hospital-playlist/auth"
	"hospital-playlist/booking"
	"hospital-playlist/entity"
	"hospital-playlist/helper"

	"github.com/gin-gonic/gin"
)

type bookingHandler struct {
	bookService booking.Service
	authService auth.Service
}

func NewBookingHandler(bookService booking.Service, authService auth.Service) *bookingHandler {
	return &bookingHandler{bookService, authService}
}

func (h *bookingHandler) GetAllBookHandler(c *gin.Context) {
	drugs, err := h.bookService.GetAllBooks()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all booking", 200, "success", drugs)
	c.JSON(200, response)
}

func (h *bookingHandler) GetBookingByID(c *gin.Context) {
	id := c.Params.ByName("book_id")

	book, err := h.bookService.GetBookingByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request booking ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", book)
	c.JSON(200, response)
}

func (h *bookingHandler) SaveNewBookingHandler(c *gin.Context) {
	var inputBooking entity.BookingInput

	if err := c.ShouldBindJSON(&inputBooking); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newBook, err := h.bookService.SaveNewBook(inputBooking)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new Spesialist", 201, "status Created", newBook)
	c.JSON(201, response)
}
