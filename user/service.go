package user

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput, email string) (UserFormat, error)
	GetUserByID(userID string) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUser) (UserFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	// pengecekan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	// bisnis logic
	users, err := s.repository.FindAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput, email string) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	checkStatus, err := s.repository.FindByEmail(email)
	formatEmail := FormatUser(checkStatus)

	if err != nil {
		return formatEmail, err
	}

	if formatEmail.Email == email || len(formatEmail.Email) == 1 {
		// errorStatus := fmt.Sprintf("this email :%s has been created", email)
		// return formatEmail, errors.New(errorStatus)
		test := formatEmail.Email == email
		test1 := fmt.Sprintf("%t", test)
		return formatEmail, errors.New(test1)
	}

	var newUser = entity.User{
		NamaLengkap:  user.NamaLengkap,
		Email:        email,
		Password:     string(genPassword),
		Alamat:       user.Alamat,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Role:         "Pasien",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) GetUserByID(userID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	var userData = entity.User{
		ID:           user.ID,
		Email:        user.Email,
		NamaLengkap:  user.NamaLengkap,
		TanggalLahir: user.TanggalLahir,
		Alamat:       user.Alamat,
		JenisKelamin: user.JenisKelamin,
		Role:         user.Role,
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatPasien := FormatUser(userData)

	return formatPasien, nil

}

func (s *service) DeleteUserByID(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}

func (s *service) UpdateUserByID(userID string, dataInput entity.UpdateUser) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	patient, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if patient.ID == 0 {
		newError := fmt.Sprintf("patient id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.NamaLengkap != "" || len(dataInput.NamaLengkap) != 0 {
		dataUpdate["nama_pasien"] = dataInput.NamaLengkap
	}

	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	if dataInput.Alamat != "" || len(dataInput.Alamat) != 0 {
		dataUpdate["alamat"] = dataInput.Alamat
	}

	if dataInput.JenisKelamin != "" || len(dataInput.JenisKelamin) != 0 {
		dataUpdate["jenis_kelamin"] = dataInput.JenisKelamin
	}

	// if dataInput.TanggalLahir != time.Time(){

	// }

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	patientUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(patientUpdated)

	return formatUser, nil
}
