package database

import (
	"log"
	"yorozuya/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	dsn := configs.GetMysqlConfig().FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
