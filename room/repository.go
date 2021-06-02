package room

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByRoomID(roomID string) (entity.Room, error)
	Create(input entity.Room) (entity.Room, error)
	FindAll() ([]entity.Room, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Room, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Room, error) {
	var rooms []entity.Room

	if err := r.db.Find(&rooms).Error; err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (r *repository) FindByRoomID(roomID string) (entity.Room, error) {
	var room entity.Room

	if err := r.db.Where("spesialist_id = ?", roomID).Find(&room).Error; err != nil {
		return room, err
	}

	return room, nil
}

func (r *repository) Create(input entity.Room) (entity.Room, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Room, error) {
	var room entity.Room

	if err := r.db.Model(&room).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return room, err
	}

	if err := r.db.Where("id = ?", ID).Find(&room).Error; err != nil {
		return room, err
	}

	return room, nil
}
