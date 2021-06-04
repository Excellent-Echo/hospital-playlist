package bookingobat

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
)

type Service interface {
	GetBookingByID(BookingID string, DrugID string) (entity.BookingObat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBookingByID(BookingID string, DrugID string) (entity.BookingObat, error) {
	book, err := s.repository.FindByBookingObatID(BookingID, DrugID)

	if err != nil {
		return book, err
	}

	if book.BookingID == 0 {
		errStatus := fmt.Sprintf("booking id and drug id %s not created", BookingID)
		return book, errors.New(errStatus)
	}
	if book.DrugID == 0 {
		errStatus := fmt.Sprintf("booking id and drug id %s not created", DrugID)
		return book, errors.New(errStatus)
	}

	return book, nil
}
