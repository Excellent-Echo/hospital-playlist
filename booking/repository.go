package booking

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Booking, error)
	FindByBookingID(ID string) (entity.Booking, error)
	Create(input entity.Booking) (entity.Booking, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Booking, error) {
	var books []entity.Booking

	if err := r.db.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindByBookingID(ID string) (entity.Booking, error) {
	var book entity.Booking

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Create(input entity.Booking) (entity.Booking, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}
	return input, nil
}
