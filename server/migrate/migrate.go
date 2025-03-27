package main

import (
	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/database"
	"github.com/savvy-bit/gin-react-postgres/models"
)

// Load environment variables & Connect DB
func init() {
	config.Init()
	database.Init()
}

func main() {
	database.DB().AutoMigrate(&models.User{})
}
