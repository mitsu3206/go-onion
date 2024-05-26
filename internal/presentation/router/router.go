package router

import (
	"yorozuya/internal/infrastrcuture/repository"
	"yorozuya/internal/presentation/controller"
	"yorozuya/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	userController := controller.NewUserController(usecase.NewUserUsecase(repository.NewUserRepository(db)))
	r := gin.Default()

	user := r.Group("user")
	{
		user.POST("/create", userController.CreateUser)
		user.GET("/:id", userController.GetUserById)
	}
	return r
}
