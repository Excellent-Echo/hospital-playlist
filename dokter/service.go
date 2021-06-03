package dokter

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllDokter() ([]DokterFormat, error)
	SaveNewDokter(dokter entity.CreateDokter) (DokterFormat, error)
	GetDokterByID(ID string) (DokterFormat, error)
	UpdateDokterByID(userID string, dataInput entity.UpdateDokter) (DokterFormat, error)
	LoginDokter(input entity.LoginDokter) (entity.Dokter, error)
}

type service struct {
	repository Repository
	// spesialistRepo spesialist.Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) LoginDokter(input entity.LoginDokter) (entity.Dokter, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("dokter id %v not found", user.ID)
		return user, errors.New(newError)
	}

	// pengecekan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}

func (s *service) GetAllDokter() ([]DokterFormat, error) {
	// bisns logic
	Dokters, err := s.repository.FindAll()
	var formatDokters []DokterFormat
	for _, Dokter := range Dokters {
		formatDokter := FormatDokter(Dokter)
		formatDokters = append(formatDokters, formatDokter)
	}

	if err != nil {
		return formatDokters, err
	}

	return formatDokters, nil
}

func (s *service) SaveNewDokter(dokter entity.CreateDokter) (DokterFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(dokter.Password), bcrypt.MinCost)

	if err != nil {
		return DokterFormat{}, err
	}

	// IDSpecialist, _ := strconv.Atoi(SpecialistID)

	var newDokter = entity.Dokter{
		Email:        dokter.Email,
		Password:     string(genPassword),
		FullName:     dokter.FullName,
		SpecialistID: dokter.SpecialistID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createUser, err := s.repository.Create(newDokter)
	formatUser := FormatDokter(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) GetDokterByID(ID string) (DokterFormat, error) {
	specialist, err := s.repository.FindByID(ID)
	formatDokter := FormatDokter(specialist)

	if err != nil {
		return formatDokter, err
	}

	if specialist.ID == 0 {
		errStatus := fmt.Sprintf("userdetail for user id %s not created", ID)
		return formatDokter, errors.New(errStatus)
	}

	return formatDokter, nil

}

func (s *service) UpdateDokterByID(userID string, dataInput entity.UpdateDokter) (DokterFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return DokterFormat{}, err
	}

	dokter, err := s.repository.FindByID(userID)

	if err != nil {
		return DokterFormat{}, err
	}

	if dokter.ID == 0 {
		newError := fmt.Sprintf("dokter id %s not found", userID)
		return DokterFormat{}, errors.New(newError)
	}

	// if dataInput.NamaLengkap != "" || len(dataInput.NamaLengkap) != 0 {
	// 	dataUpdate["nama_pasien"] = dataInput.NamaLengkap
	// }

	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	if dataInput.FullName != "" || len(dataInput.FullName) != 0 {
		dataUpdate["full_name"] = dataInput.FullName
	}

	// if dataInput.JenisKelamin != "" || len(dataInput.JenisKelamin) != 0 {
	// 	dataUpdate["jenis_kelamin"] = dataInput.JenisKelamin
	// }

	// if dataInput.TanggalLahir != "00/00/000" {

	// }

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return DokterFormat{}, err
	}

	formatUser := FormatDokter(userUpdated)

	return formatUser, nil
}
