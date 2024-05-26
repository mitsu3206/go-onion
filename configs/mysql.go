package configs

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetMysqlConfig() *mysql.Config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	c := &mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
		Loc:       jst,
	}
	return c
}
