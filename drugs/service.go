package drugs

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
)

type Service interface {
	GetAllDrugs() ([]entity.Drug, error)
	GetSpesialistByID(ID string) (entity.Drug, error)
	SaveNewDrug(input entity.DrugInput) (entity.Drug, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDrugs() ([]entity.Drug, error) {
	drugs, err := s.repository.FindAll()

	if err != nil {
		return drugs, err
	}

	return drugs, nil
}

func (s *service) GetSpesialistByID(ID string) (entity.Drug, error) {
	drug, err := s.repository.FindByDrugID(ID)

	if err != nil {
		return drug, err
	}

	if drug.ID == 0 {
		errStatus := fmt.Sprintf("drug id %s not created", ID)
		return drug, errors.New(errStatus)
	}

	return drug, nil
}

func (s *service) SaveNewDrug(input entity.DrugInput) (entity.Drug, error) {
	var NewSpesialist = entity.Drug{
		Name:  input.Name,
		Price: input.Price,
	}

	createSpesialist, err := s.repository.Create(NewSpesialist)

	if err != nil {
		return createSpesialist, err
	}

	return createSpesialist, nil
}
