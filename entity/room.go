package entity

import "time"

type Room struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	JenisKamar string    `json:"jenis_kamar"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RoomInput struct {
	JenisKamar string `json:"jenis_kamar"`
}
