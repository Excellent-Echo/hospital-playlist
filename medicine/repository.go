package medicine

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByMedicineID(medicineID string) (entity.Medicine, error)
	Create(input entity.Medicine) (entity.Medicine, error)
	FindAll() ([]entity.Medicine, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Medicine, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Medicine, error) {
	var medicine []entity.Medicine

	if err := r.db.Find(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (r *repository) FindByMedicineID(medicineID string) (entity.Medicine, error) {
	var medicine entity.Medicine

	if err := r.db.Where("medicine_id = ?", medicineID).Find(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (r *repository) Create(input entity.Medicine) (entity.Medicine, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Medicine, error) {
	var medicine entity.Medicine

	if err := r.db.Model(&medicine).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return medicine, err
	}

	if err := r.db.Where("id = ?", ID).Find(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}
