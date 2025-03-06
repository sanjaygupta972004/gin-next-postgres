package database

import (
	"log"

	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := config.Global.Database.URL

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})
	return db
}
