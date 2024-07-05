package database

import (
	"log"

	"github.com/YerzatCode/auth-service/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(url string) {
	var err error
	const op = "internal.database.InitDB"
	DB, err = gorm.Open(sqlite.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("%w: %s", op, err)
	}
	DB.AutoMigrate(&model.UserModel{})

}
