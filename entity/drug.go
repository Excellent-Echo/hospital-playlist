package entity

import "time"

type Drug struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateDrug struct {
	Name      string    `json:"name" binding:"required"`
	Price     int       `json:"price" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateDrug struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
