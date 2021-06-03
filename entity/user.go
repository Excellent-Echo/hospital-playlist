package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	// Role     string `json:"role"`
	// TanggalLahir time.Time `json:"tanggal_lahir"`
	// Alamat       string    `json:"alamat"`
	// JenisKelamin string    `json:"jenis_kelamin"`
	// Role         string    `json:"role"`
}

type UpdateUser struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	// TanggalLahir time.Time `json:"tanggal_lahir"`
	// Alamat       string    `json:"alamat"`
	// JenisKelamin string    `json:"jenis_kelamin"`
}

type UserDetailOutput struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	Email       string      `json:"email"`
	Password    string      `json:"-"`
	FullName    string      `json:"full_name"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	UserProfile UserProfile `json:"user_profile"`
	UserDetail  UserDetail  `json:"user_detail"`
}
