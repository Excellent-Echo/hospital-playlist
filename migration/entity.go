package migration

import "time"

type User struct {
	ID          int          `gorm:"primaryKey" json:"id"`
	NamaLengkap string       `json:"nama_lengkap"`
	Email       string       `json:"email"`
	Password    string       `json:"-"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   time.Time    `gorm:"index" json:"-"`
	Specialists []Specialist `gorm:"foreignKey:UserID" json:"spesialis"`
	Medicines   []Medicine   `gorm:"foreignKey:UserID" json:"medicine"`
	Rooms       []Room       `gorm:"foreignKey:UserID" json:"room"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"` // kita tangkap dari file (foto) , path / dir file foto
	UserID      int    `json:"user_id"`
}

type Room struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	JenisKamar string    `json:"jenis_kamar"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `gorm:"index" json:"-"`
}

type Specialist struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	NamaSpesialis string `json:"nama_spesialis"`
}

type Medicine struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	NamaObat  string    `json:"nama_obat"`
	HargaObat int       `json:"harga_obat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}
