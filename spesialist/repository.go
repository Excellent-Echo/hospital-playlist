package spesialist

import (
	"hospital-playlist/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindBySpesialistID(spesialistID string) (entity.Specialist, error)
	Create(input entity.Specialist) (entity.Specialist, error)
	FindAll() ([]entity.Specialist, error)
	FindSpecialistByNameDokter(spesialistID string) (entity.Specialist, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Specialist, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Specialist, error) {
	var specialists []entity.Specialist

	if err := r.db.Preload("Dokters").Find(&specialists).Error; err != nil {
		return specialists, err
	}

	return specialists, nil
}

func (r *repository) FindBySpesialistID(spesialistID string) (entity.Specialist, error) {
	var spesialist entity.Specialist

	if err := r.db.Preload("Dokters").Where("id = ?", spesialistID).Find(&spesialist).Error; err != nil {
		return spesialist, err
	}

	return spesialist, nil
}

func (r *repository) FindSpecialistByNameDokter(spesialistID string) (entity.Specialist, error) {
	var spesialist entity.Specialist

	// if err := r.db.Where("id = ?", spesialistID).Find(&spesialist).Error; err != nil
	// if err := r.db.Table("specialists").Select("specialists.id, specialists.name, dokters.id, dokters.full_name").Joins("left join dokters on dokters.specialist_id = specialists.id").Where("specialists.id = ?", spesialistID).Scan(&spesialist).Error; err != nil
	if err := r.db.Joins("dokters").Find(&spesialist).Where("id = ?", spesialistID).Error; err != nil {
		// if err := r.db.Where("id = ?", spesialistID).Preload("dokters").Find(spesialist).Error; err != nil {
		return spesialist, err
	}

	return spesialist, nil
}

func (r *repository) Create(input entity.Specialist) (entity.Specialist, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Specialist, error) {
	var specialist entity.Specialist

	if err := r.db.Model(&specialist).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return specialist, err
	}

	if err := r.db.Where("id = ?", ID).Find(&specialist).Error; err != nil {
		return specialist, err
	}

	return specialist, nil
}
