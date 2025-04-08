package migration

import (
	"errors"

	"github.com/savvy-bit/gin-react-postgres/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	var err error = nil
	// if err = db.Migrator().DropTable(&models.User{}); err != nil {
	// 	panic("Failed to drop tables: " + err.Error())
	// }
	if err := models.CreateEnumUserRole(db); err != nil {
		return err
	}

	if db == nil {
		return errors.New("db instance is nil; ensure it is properly initialized")
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate User table: " + err.Error())
	}

	println("Database migration completed successfully")
	return nil
}
