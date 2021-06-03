package entity

import "time"

type UserDetail struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	NoHandphone uint      `json:"no_handphone"`
	BirthDate   time.Time `json:"birth_date"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   time.Time `gorm:"index"`
}

type CreateUserDetail struct {
	NoHandphone uint      `json:"no_handphone" binding:"required"`
	BirthDate   time.Time `json:"birth_date"`
	Gender      string    `json:"gender" binding:"required"`
	Address     string    `json:"address" binding:"required"`
}

type UpdateUserDetail struct {
	NoHandphone uint      `json:"no_handphone"`
	BirthDate   time.Time `json:"birth_date"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
}
