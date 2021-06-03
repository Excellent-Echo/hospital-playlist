package drug

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Drug, error)
	FindByID(drugID string) (entity.Drug, error)
	Create(input entity.Drug) (entity.Drug, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Drug, error)
	// DeleteByID(categoryID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Drug, error) {
	var drugs []entity.Drug

	if err := r.db.Find(&drugs).Error; err != nil {
		return drugs, err
	}

	return drugs, nil
}

func (r *repository) Create(input entity.Drug) (entity.Drug, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Drug, error) {
	var drug entity.Drug

	if err := r.db.Model(&drug).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return drug, err
	}

	if err := r.db.Where("id = ?", ID).Find(&drug).Error; err != nil {
		return drug, err
	}

	return drug, nil

}

// func (r *repository) DeleteByID(categoryID string) (string, error) {
// 	if err := r.db.Where("id = ?", categoryID).Delete(&entity.Category{}).Error; err != nil {
// 		return "error", err
// 	}

// 	status := fmt.Sprintf("category id %v success deleted", categoryID)

// 	return status, nil
// }

func (r *repository) FindByID(drugID string) (entity.Drug, error) {
	var drug entity.Drug

	if err := r.db.Where("user_id = ?", drugID).Find(&drug).Error; err != nil {
		return drug, err
	}

	return drug, nil
}
