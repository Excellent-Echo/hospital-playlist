package entity

import "time"

type Medicine struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	NamaObat  string    `json:"nama_obat"`
	Dosis     string    `json:"dosis"`
	HargaObat int       `json:"harga_obat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
}

type MedicineInput struct {
	NamaObat  string `json:"nama_obat"`
	Dosis     string `json:"dosis"`
	HargaObat int    `json:"harga_obat"`
}
