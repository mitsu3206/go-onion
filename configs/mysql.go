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
		log.Println(err)
	}
	c := &mysql.Config{
		User:      findEnv("DB_USER"),
		Passwd:    findEnv("DB_PASS"),
		Net:       "tcp",
		Addr:      findEnv("DB_HOST") + ":" + findEnv("DB_PORT"),
		DBName:    findEnv("DB_NAME"),
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
		Loc:       jst,
	}
	return c
}

func findEnv(key string) string {
	env, err := os.LookupEnv(key)
	if !err {
		log.Printf("error: Failed to locate the 'env' for key = %s", key)
	}
	return env
}
