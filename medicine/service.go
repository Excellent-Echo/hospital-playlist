package medicine

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"time"
)

type Service interface {
	GetAllMedicine() ([]entity.Medicine, error)
	GetMedicineByID(ID string) (entity.Medicine, error)
	SaveNewMedicine(input entity.MedicineInput) (entity.Medicine, error)
	UpdateMedicineByID(ID string, dataInput entity.MedicineInput) (entity.Medicine, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllMedicine() ([]entity.Medicine, error) {
	medicines, err := s.repository.FindAll()

	if err != nil {
		return medicines, err
	}

	return medicines, nil
}

func (s *service) GetMedicineByID(ID string) (entity.Medicine, error) {
	room, err := s.repository.FindByMedicineID(ID)

	if err != nil {
		return room, err
	}

	if room.ID == 0 {
		errStatus := fmt.Sprintf("medicine %s not created", ID)
		return room, errors.New(errStatus)
	}

	return room, nil
}

func (s *service) SaveNewMedicine(input entity.MedicineInput) (entity.Medicine, error) {
	var NewMedicine = entity.Medicine{
		NamaObat:  input.NamaObat,
		Dosis:     input.Dosis,
		HargaObat: input.HargaObat,
	}

	createMedicine, err := s.repository.Create(NewMedicine)

	if err != nil {
		return createMedicine, err
	}

	return createMedicine, nil
}

func (s *service) UpdateMedicineByID(ID string, dataInput entity.MedicineInput) (entity.Medicine, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Medicine{}, err
	}

	room, err := s.repository.FindByMedicineID(ID)

	if err != nil {
		return entity.Medicine{}, err
	}

	if room.ID == 0 {
		newError := fmt.Sprintf("room id %s not found", ID)
		return entity.Medicine{}, errors.New(newError)
	}

	if dataInput.NamaObat != "" || len(dataInput.NamaObat) != 0 {
		dataUpdate["nama_obat"] = dataInput.NamaObat
	}

	if dataInput.Dosis != "" || len(dataInput.Dosis) != 0 {
		dataUpdate["dosis"] = dataInput.NamaObat
	}

	if dataInput.HargaObat != 0 {
		dataUpdate["harga_obat"] = dataInput.HargaObat
	}

	dataUpdate["updated_at"] = time.Now()

	medicineUpdate, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return medicineUpdate, err
	}

	return medicineUpdate, nil

}
