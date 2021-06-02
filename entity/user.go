package entity

import "time"

type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	NamaLengkap  string    `json:"nama_lengkap"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	JenisKelamin string    `json:"jenis_kelamin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Role         string    `json:"role"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	NamaLengkap  string    `json:"nama_lengkap"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Role         string    `json:"role"`
}

type UpdateUser struct {
	Email        string    `json:"email"`
	NamaLengkap  string    `json:"nama_lengkap"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	JenisKelamin string    `json:"jenis_kelamin"`
}

type UserDetailOutput struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	NamaLengkap  string    `json:"nama_lengkap"`
	Email        string    `json:"email"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	JenisKelamin string    `json:"jenis_kelamin"`
	// IdDokter     int       `json:"id_dokter"`
	// NamaDokter   string    `json:"nama_dokter"`
}
