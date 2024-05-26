package repository

import (
	"log"
	"time"
	"yorozuya/internal/domain/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db}
}

func (ur userRepository) CreateUser(name, email string, age uint8, birthday time.Time) (model.User, error) {
	user := model.User{Name: name, Email: &email, Age: age, Birthday: &birthday}
	result := ur.db.Create(&user)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	return model.User{ID: user.ID, Name: user.Name, Email: user.Email, Age: user.Age, Birthday: user.Birthday}, nil
}
