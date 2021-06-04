package entity

import "time"

type Specialist struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Dokters   []Dokter
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SpecialistInput struct {
	Name string `json:"name"`
}

type SpecialistByNameDocter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IDDokter   int    `json:"id_dokter"`
	NameDokter string `json:"name_dockter"`
}
