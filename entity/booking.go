package entity

import "time"

type Booking struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Date        time.Time `json:"date"`
	NameBooking string    `json:"name_booking"`
	UserID      int       `json:"user_id"`
	DokterID    int       `json:"dokter_id"`
}

type BookingInput struct {
	Date        time.Time `json:"date"`
	NameBooking string    `json:"name_booking"`
	UserID      int       `json:"user_id"`
	DokterID    int       `json:"dokter_id"`
}
