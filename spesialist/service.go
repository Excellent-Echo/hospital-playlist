package spesialist

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"time"
)

type Service interface {
	GetAllSpesialist() ([]entity.Specialist, error)
	GetSpesialistByID(ID string) (entity.Specialist, error)
	GetSpecialistByNameDokter(spesialistID string) (entity.Specialist, error)
	SaveNewSpesialist(input entity.SpecialistInput) (entity.Specialist, error)
	UpdateSpesialistByID(ID string, dataInput entity.SpecialistInput) (entity.Specialist, error)
}

type service struct {
	repository Repository
	// dokterRepo dokter.Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllSpesialist() ([]entity.Specialist, error) {
	specialists, err := s.repository.FindAll()

	if err != nil {
		return specialists, err
	}

	return specialists, nil
}

func (s *service) GetSpesialistByID(ID string) (entity.Specialist, error) {
	specialist, err := s.repository.FindBySpesialistID(ID)

	if err != nil {
		return specialist, err
	}

	if specialist.ID == 0 {
		errStatus := fmt.Sprintf("userdetail for user id %s not created", ID)
		return specialist, errors.New(errStatus)
	}

	return specialist, nil
}

func (s *service) GetSpecialistByNameDokter(spesialistID string) (entity.Specialist, error) {
	specialist, err := s.repository.FindSpecialistByNameDokter(spesialistID)

	if err != nil {
		return specialist, err
	}

	if specialist.ID == 0 {
		errStatus := fmt.Sprintf("userdetail for user id %s not created", spesialistID)
		return specialist, errors.New(errStatus)
	}

	return specialist, nil

}

func (s *service) SaveNewSpesialist(input entity.SpecialistInput) (entity.Specialist, error) {
	var NewSpesialist = entity.Specialist{
		Name: input.Name,
	}

	createSpesialist, err := s.repository.Create(NewSpesialist)

	if err != nil {
		return createSpesialist, err
	}

	return createSpesialist, nil
}

func (s *service) UpdateSpesialistByID(ID string, dataInput entity.SpecialistInput) (entity.Specialist, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Specialist{}, err
	}

	specialist, err := s.repository.FindBySpesialistID(ID)

	if err != nil {
		return entity.Specialist{}, err
	}

	if specialist.ID == 0 {
		newError := fmt.Sprintf("spesialis id %s not found", ID)
		return entity.Specialist{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}

	dataUpdate["updated_at"] = time.Now()

	spesiaistUpdate, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return spesiaistUpdate, err
	}

	return spesiaistUpdate, nil

}
