package bookingobat

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByBookingObatID(BookingID string, DrugID string) (entity.BookingObat, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByBookingObatID(BookingID string, DrugID string) (entity.BookingObat, error) {
	var book entity.BookingObat

	if err := r.db.Where("booking_id = ?", BookingID).Where("drug_id = ?", DrugID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
