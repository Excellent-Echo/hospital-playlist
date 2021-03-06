package migration

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	// Medicines []Drug    `gorm:"foreignKey:UserID" json:"medicine"`
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

type Room struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	RoomType  string    `json:"room_type"`
	Rates     int       `json:"rates"`
	User      User      `gorm:"foreignKey:RommID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Specialist struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	// User      []User    `gorm:"foreignKey:MedicalSpecialistID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

type Drug struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}
