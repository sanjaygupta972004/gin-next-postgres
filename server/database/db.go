package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/savvy-bit/gin-react-postgres/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() any
	instance    any
	once        sync.Once
}

var (
	dbInstance *DBInstance
)

// Instance gets the singleton instance
func (i *DBInstance) Instance() any {
	i.once.Do(func() {
		i.instance = i.initializer()
	})
	return i.instance
}

func ConnectDB() any {
	dbURL := config.Global.Database.URL

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func CloseDB() {
	dbInstance = nil
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func Init() {
	fmt.Println("⌛             Loading database...             ⌛")
	fmt.Println("=================================================")
	dbInstance = &DBInstance{initializer: ConnectDB}
}
