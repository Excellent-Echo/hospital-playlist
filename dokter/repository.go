package dokter

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Dokter, error)
	Create(user entity.Dokter) (entity.Dokter, error)
	FindByID(ID string) (entity.Dokter, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Dokter, error)
	FindByEmail(email string) (entity.Dokter, error)
	// FindBYRoleDocter(role string) (entity.Dokter, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Dokter, error) {
	var users []entity.Dokter

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entity.Dokter) (entity.Dokter, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID string) (entity.Dokter, error) {
	var user entity.Dokter

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// func (r *repository) FindBYRoleDocter(role string) (entity.Dokter, error) {
// 	var user entity.Dokter

// 	if err := r.db.Where("role = ?", role).Find(&user).Error; err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Dokter{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Dokter, error) {

	var user entity.Dokter

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.Dokter, error) {
	var user entity.Dokter

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
