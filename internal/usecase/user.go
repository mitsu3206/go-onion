package usecase

import (
	"log"
	"time"
	"yorozuya/internal/domain/model"
)

type UserUsecase interface {
	CreateUser(name, email string, age uint8, birthday time.Time) (model.User, error)
}

type userUsecase struct {
	userRepository model.UserRepository
}

func NewUserUsecase(ur model.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu userUsecase) CreateUser(name, email string, age uint8, birthday time.Time) (model.User, error) {
	user, err := uu.userRepository.CreateUser(name, email, age, birthday)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}
