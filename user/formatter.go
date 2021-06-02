package user

import (
	"hospital-playlist/entity"
	"time"
)

type UserFormat struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	NamaLengkap string `json:"nama_lengkap"`
	// TanggalLahir time.Time `json:"tanggal_lahir"`
	// Alamat       string    `json:"alamat"`
	// JenisKelamin string    `json:"jenis_kelamin"`
	Role string `json:"role"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:          user.ID,
		Email:       user.Email,
		NamaLengkap: user.NamaLengkap,
		// TanggalLahir: user.TanggalLahir,
		// Alamat:       user.Alamat,
		// JenisKelamin: user.JenisKelamin,
		Role: user.Role,
	}

	return formatUser
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
