package drug

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
)

type Service interface {
	GetAllDrug() ([]entity.Drug, error)
	SaveNewDrug(input entity.CreateDrug) (entity.Drug, error)
	UpdateDrugByID(drugID string, dataInput entity.UpdateDrug) (entity.Drug, error)
	// DeleteCategoryByID(categoryID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDrug() ([]entity.Drug, error) {
	drugs, err := s.repository.FindAll()

	if err != nil {
		return drugs, err
	}

	return drugs, nil
}

func (s *service) SaveNewDrug(input entity.CreateDrug) (entity.Drug, error) {

	var newDrug = entity.Drug{
		Name:  input.Name,
		Price: input.Price,
	}

	createDrug, err := s.repository.Create(newDrug)

	if err != nil {
		return createDrug, err
	}

	return createDrug, nil
}

func (s *service) UpdateDrugByID(drugID string, dataInput entity.UpdateDrug) (entity.Drug, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(drugID); err != nil {
		return entity.Drug{}, err
	}

	// todo, err := s.repository.FindByID(categoryID)

	drug, err := s.repository.FindByID(drugID)

	if err != nil {
		return entity.Drug{}, err
	}

	// if todo.ID == 0 {
	// 	newError := fmt.Sprintf("todo id %s not found", categoryID)
	// 	return entity.Category{}, errors.New(newError)
	// }

	if drug.ID == 0 {
		newError := fmt.Sprintf("drug id %s not found", drugID)
		return entity.Drug{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Price != 0 {
		dataUpdate["price"] = dataInput.Price
	}
	// fmt.Println(dataUpdate)

	drugUpdated, err := s.repository.UpdateByID(drugID, dataUpdate)

	if err != nil {
		return drugUpdated, err
	}

	return drugUpdated, nil
}

// func (s *service) DeleteCategoryByID(categoryID string) (interface{}, error) {
// 	if err := helper.ValidateIDNumber(categoryID); err != nil {
// 		return nil, err
// 	}

// 	cat, err := s.repository.FindByID(categoryID)

// 	if err != nil {
// 		return nil, err
// 	}
// 	if cat.ID == 0 {
// 		newError := fmt.Sprintf("category id %s not found", categoryID)
// 		return nil, errors.New(newError)
// 	}

// 	status, err := s.repository.DeleteByID(categoryID)

// 	if status == "error" {
// 		return nil, errors.New("error delete in internal server")
// 	}

// 	msg := fmt.Sprintf("success delete category ID : %s", categoryID)

// 	formatDelete := FormatDeleteCategory(msg)

// 	return formatDelete, nil
// }
