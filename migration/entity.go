package migration

import "time"

type Room struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	JenisKamar string    `json:"jenis_kamar"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Specialist struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	NamaSpesialis string    `json:"nama_spesialis"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Medicine struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	NamaObat  string    `json:"nama_obat"`
	Dosis     string    `json:"dosis"`
	HargaObat int       `json:"harga_obat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	NoHandphone uint
	Address     string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

type User struct {
	ID           int          `gorm:"primaryKey" json:"id"`
	NamaLengkap  string       `json:"nama_lengkap"`
	Email        string       `json:"email"`
	Password     string       `json:"-"`
	TanggalLahir time.Time    `json:"tanggal_lahir"`
	Alamat       string       `json:"alamat"`
	JenisKelamin string       `json:"jenis_kelamin"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    time.Time    `gorm:"index" json:"-"`
	Specialist   []Specialist `gorm:"foreignKey:UserID"`
	Medicine     []Medicine   `gorm:"foreignKey:UserID"`
	Room         []Room       `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"` // kita tangkap dari file (foto) , path / dir file foto
	UserID      int    `json:"user_id"`
}
