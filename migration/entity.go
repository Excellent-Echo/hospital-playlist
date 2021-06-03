package migration

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	Bookings  []Booking `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Dokter struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"unique" json:"email"`
	Password     string    `json:"-"`
	FullName     string    `json:"full_name"`
	Bookings     []Booking `gorm:"foreignKey:DokterID"`
	SpecialistID int
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `gorm:"index" json:"-"`
}

type UserDetail struct {
	ID          int       `gorm:"primaryKey"`
	NoHandphone uint      `json:"no_handphone"`
	BirthDate   time.Time `json:"birth_date"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"` // kita tangkap dari file (foto) , path / dir file foto
	UserID      int    `json:"user_id"`
}

type Specialist struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Dokters   []Dokter  `gorm:"foreignKey:SpecialistID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Drug struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Bookings  []Booking `gorm:"many2many:booking_obat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Booking struct {
	ID          int `gorm:"primaryKey" json:"id"`
	Date        time.Time
	NameBooking string `json:"name_booking"`
	UserID      int
	DokterID    int
	Drugs       []Drug `gorm:"many2many:booking_obat"`
}
