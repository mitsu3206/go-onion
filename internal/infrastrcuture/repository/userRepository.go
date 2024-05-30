package repository

import (
	"log"
	"yorozuya/internal/domain/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db}
}

func (ur userRepository) CreateUser(user model.User) (model.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, nil
}

func (ur userRepository) GetUserById(id uint) (model.User, error) {
	user := model.User{}
	result := ur.db.First(&user, id)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, nil
}

func (ur userRepository) UpdateUser(user model.User) (model.User, error) {
	result := ur.db.Updates(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, nil
}

func (ur userRepository) DeleteUser(user model.User) error {
	result := ur.db.Delete(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return nil
}
