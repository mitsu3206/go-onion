package controller

import (
	"log"
	"net/http"
	"time"
	"yorozuya/internal/usecase"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

type CreateJson struct {
	Name     string
	Email    string
	Age      uint8
	Birthday time.Time
}

func NewUserController(uu usecase.UserUsecase) *userController {
	return &userController{uu}
}

func (uc userController) CreateUser(c *gin.Context) {
	var json CreateJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userUsecase.CreateUser(json.Name, json.Email, json.Age, json.Birthday)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
