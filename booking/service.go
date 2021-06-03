package booking

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
)

type Service interface {
	GetAllBooks() ([]entity.Booking, error)
	GetBookingByID(ID string) (entity.Booking, error)
	SaveNewBook(input entity.BookingInput) (entity.Booking, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllBooks() ([]entity.Booking, error) {
	books, err := s.repository.FindAll()

	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetBookingByID(ID string) (entity.Booking, error) {
	book, err := s.repository.FindByBookingID(ID)

	if err != nil {
		return book, err
	}

	if book.ID == 0 {
		errStatus := fmt.Sprintf("booking id %s not created", ID)
		return book, errors.New(errStatus)
	}

	return book, nil
}

func (s *service) SaveNewBook(input entity.BookingInput) (entity.Booking, error) {
	var NewSpesialist = entity.Booking{
		Date:        input.Date,
		NameBooking: input.NameBooking,
		UserID:      input.UserID,
		DokterID:    input.DokterID,
	}

	createSpesialist, err := s.repository.Create(NewSpesialist)

	if err != nil {
		return createSpesialist, err
	}

	return createSpesialist, nil
}
