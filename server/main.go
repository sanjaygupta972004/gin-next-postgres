package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/database"
	"github.com/savvy-bit/gin-react-postgres/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Load environment variables & Connect DB
func init() {
	config.Init()
	database.Init()
}

// @title Gin + Postgres Back-end Swagger Documentation
// @version 1.0
// @description Testing Swagger APIs.
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

func main() {
	// Defaulting to the port specified in the global configuration
	addr := flag.String("addr", config.Global.Server.Port, "Address to listen and serve")
	flag.Parse()

	// Swagger URL
	url := ginSwagger.URL("http://localhost:8080/swagger.json")

	// Set Gin mode
	gin.SetMode(config.Global.Server.Mode)

	// access database for global variable
	db := database.DB()
	if db == nil {
		log.Fatal("Database connection failed")
		os.Exit(1)
	}
	defer database.CloseDB()

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	app.StaticFile("/swagger.json", filepath.Join(config.Global.Server.DocumentDir, "swagger.json"))
	app.Static("/images", filepath.Join(config.Global.Server.StaticDir, "img"))
	app.StaticFile("/favicon.ico", filepath.Join(config.Global.Server.StaticDir, "img/favicon.ico"))
	app.MaxMultipartMemory = config.Global.Server.MaxMultipartMemory << 20

	routers.SetupRouters(app, db)

	// Listen and Serve
	if err := app.Run(*addr); err != nil {
		log.Fatal(err.Error())
	}
}
