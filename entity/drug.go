package entity

import "time"

type Drug struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	// Bookings  []Booking `gorm:"many2many:booking_obat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type DrugInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
