package drugs

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Drug, error)
	FindByDrugID(ID string) (entity.Drug, error)
	Create(input entity.Drug) (entity.Drug, error)
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

func (r *repository) FindByDrugID(ID string) (entity.Drug, error) {
	var drug entity.Drug

	if err := r.db.Where("id = ?", ID).Find(&drug).Error; err != nil {
		return drug, err
	}

	return drug, nil
}

func (r *repository) Create(input entity.Drug) (entity.Drug, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}
	return input, nil
}
