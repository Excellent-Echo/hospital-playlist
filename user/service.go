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
	SaveNewUser(user entity.CreateUser) (UserFormat, error)
	GetUserByID(userID string) (UserFormat, error)
	GetUserByRoleDocter(role string) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUser) (UserFormat, error)
	LoginUser(input entity.LoginUser) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) LoginUser(input entity.LoginUser) (entity.User, error) {
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
	// bisns logic
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

func (s *service) SaveNewUser(user entity.CreateUser) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		Email:     user.Email,
		Password:  string(genPassword),
		FullName:  user.FullName,
		Role:      "Pasien",
		CreatedAt: time.Now(),
		// UpdatedAt:    time.Now(),

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
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatGetUser := FormatUser(userData)

	return formatGetUser, nil

}

func (s *service) GetUserByRoleDocter(role string) (UserFormat, error) {

	if err := helper.ValidateString(role); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindBYRoleDocter(role)

	if err != nil {
		return UserFormat{}, err
	}

	var userData = entity.User{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	}

	if user.Role == "" {
		newError := fmt.Sprintf("user role %s not found", role)
		return UserFormat{}, errors.New(newError)
	}

	formatGetUser := FormatUser(userData)

	return formatGetUser, nil
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

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	// if dataInput.NamaLengkap != "" || len(dataInput.NamaLengkap) != 0 {
	// 	dataUpdate["nama_pasien"] = dataInput.NamaLengkap
	// }

	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	// if dataInput.Alamat != "" || len(dataInput.Alamat) != 0 {
	// 	dataUpdate["alamat"] = dataInput.Alamat
	// }

	// if dataInput.JenisKelamin != "" || len(dataInput.JenisKelamin) != 0 {
	// 	dataUpdate["jenis_kelamin"] = dataInput.JenisKelamin
	// }

	// if dataInput.TanggalLahir != "00/00/000" {

	// }

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdated)

	return formatUser, nil
}
