package db

import (
	"dating-sim/dotenv"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() {
	port, err := strconv.Atoi(dotenv.Env("DB_PORT"))
	if err != nil {
		panic("failed parse db port")
	}
	dsn := fmt.Sprintf(
		"host=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"port=%d "+
			"sslmode=disable "+
			"TimeZone=%s",
		dotenv.Env("DB_HOST"),
		dotenv.Env("DB_USERNAME"),
		dotenv.Env("DB_PASSWORD"),
		dotenv.Env("DB_NAME"),
		port,
		dotenv.Env("DB_TZ"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	log.Print("db connection open")
}
