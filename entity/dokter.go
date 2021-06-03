package entity

import "time"

type Dokter struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"-"`
	FullName     string `json:"full_name"`
	SpecialistID int
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `gorm:"index" json:"-"`
}

type CreateDokter struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	FullName     string `json:"full_name"`
	SpecialistID int    `json:"specialist_id"`
}

type LoginDokter struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateDokter struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
