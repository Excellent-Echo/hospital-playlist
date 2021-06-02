package room

import (
	"errors"
	"fmt"
	"hospital-playlist/entity"
	"hospital-playlist/helper"
	"time"
)

type Service interface {
	GetAllRoom() ([]entity.Room, error)
	GetRoomByID(ID string) (entity.Room, error)
	SaveNewRoom(input entity.RoomInput) (entity.Room, error)
	UpdateRoomByID(ID string, dataInput entity.RoomInput) (entity.Room, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllRoom() ([]entity.Room, error) {
	rooms, err := s.repository.FindAll()

	if err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (s *service) GetRoomByID(ID string) (entity.Room, error) {
	room, err := s.repository.FindByRoomID(ID)

	if err != nil {
		return room, err
	}

	if room.ID == 0 {
		errStatus := fmt.Sprintf("room id %s not created", ID)
		return room, errors.New(errStatus)
	}

	return room, nil
}

func (s *service) SaveNewRoom(input entity.RoomInput) (entity.Room, error) {
	var NewRoom = entity.Room{
		JenisKamar: input.JenisKamar,
	}

	createRoom, err := s.repository.Create(NewRoom)

	if err != nil {
		return createRoom, err
	}

	return createRoom, nil
}

func (s *service) UpdateRoomByID(ID string, dataInput entity.RoomInput) (entity.Room, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Room{}, err
	}

	room, err := s.repository.FindByRoomID(ID)

	if err != nil {
		return entity.Room{}, err
	}

	if room.ID == 0 {
		newError := fmt.Sprintf("room id %s not found", ID)
		return entity.Room{}, errors.New(newError)
	}

	if dataInput.JenisKamar != "" || len(dataInput.JenisKamar) != 0 {
		dataUpdate["jenis_kamar"] = dataInput.JenisKamar
	}

	dataUpdate["updated_at"] = time.Now()

	roomUpdate, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return roomUpdate, err
	}

	return roomUpdate, nil

}
