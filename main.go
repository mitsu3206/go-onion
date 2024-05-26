package main

import (
	"yorozuya/internal/domain/model"
	"yorozuya/internal/infrastrcuture/database"
	"yorozuya/internal/presentation/router"
)

func main() {
	db := database.NewGorm()
	db.AutoMigrate(&model.User{})
	r := router.NewRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
