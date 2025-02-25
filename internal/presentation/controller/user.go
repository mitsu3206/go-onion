package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"yorozuya/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userUsecase.CreateUser(json.Name, json.Email, json.Age, json.Birthday)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uc userController) GetUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userUsecase.GetUserById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc userController) IncrementAge(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userUsecase.IncrementAge(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc userController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = uc.userUsecase.DeleteUser(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
