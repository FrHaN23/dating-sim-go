package migrate

import (
	"dating-sim/db"
	"dating-sim/models"
)

func Db() {
	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Preferance{},
		&models.Picture{},
		&models.Swipe{},
	)
	if err != nil {
		panic(err)
	}
}
