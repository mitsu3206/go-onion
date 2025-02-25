package usecase

import (
	"log"
	"time"
	"yorozuya/internal/domain/model"
)

type UserUsecase interface {
	CreateUser(name, email string, age uint8, birthday time.Time) (model.User, error)
	GetUserById(id uint) (model.User, error)
	IncrementAge(id uint) (model.User, error)
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepository model.UserRepository
}

func NewUserUsecase(ur model.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu userUsecase) CreateUser(name, email string, age uint8, birthday time.Time) (model.User, error) {
	user := model.NewUser(name, email, age)
	user, err := uu.userRepository.CreateUser(user)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func (uu userUsecase) GetUserById(id uint) (model.User, error) {
	user, err := uu.userRepository.GetUserById(id)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func (uu userUsecase) IncrementAge(id uint) (model.User, error) {
	user, err := uu.userRepository.GetUserById(id)
	if err != nil {
		log.Println(err)
	}
	user.Age++
	user, err = uu.userRepository.UpdateUser(user)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func (uu userUsecase) DeleteUser(id uint) error {
	user, err := uu.userRepository.GetUserById(id)
	if err != nil {
		log.Println(err)
	}
	err = uu.userRepository.DeleteUser(user)
	if err != nil {
		log.Println(err)
	}
	return nil
}
