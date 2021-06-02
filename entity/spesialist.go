package entity

import "time"

type Specialist struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	NamaSpesialis string    `json:"nama_spesialis"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type SpecialistInput struct {
	NamaSpesialis string `json:"nama_spesialis"`
}
