package dokter

import (
	"hospital-playlist/entity"
	"time"
)

type DokterFormat struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	FullName     string    `json:"full_name"`
	SpesialistID int       `json:"specialist_id "`
	CreatedAt    time.Time `json:"created_at"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDokter(dokter entity.Dokter) DokterFormat {
	var formatDokter = DokterFormat{
		ID:           dokter.ID,
		Email:        dokter.Email,
		FullName:     dokter.FullName,
		SpesialistID: dokter.SpecialistID,
		CreatedAt:    dokter.CreatedAt,
	}

	return formatDokter
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
