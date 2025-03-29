package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/database"
	"github.com/savvy-bit/gin-react-postgres/database/migration"
	"github.com/savvy-bit/gin-react-postgres/middlewares"
	"github.com/savvy-bit/gin-react-postgres/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Load environment variables & Connect DB
func init() {
	if err := config.LoadGlobalConfig(); err != nil {
		log.Fatal(err)
	}
}

// @title Gin + Postgres Back-end Swagger Documentation
// @version 1.0
// @description Testing Swagger APIs.
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

func main() {
	// Defaulting to the port specified in the global configuration
	config := config.GetGlobalConfig()
	addr := flag.String("addr", config.Server.Port, "Address to listen and serve")
	flag.Parse()

	// Swagger URL
	url := ginSwagger.URL("http://localhost:8080/swagger.json")

	// Set Gin mode
	gin.SetMode(config.Server.Mode)

	// access database for global variable
	if err := database.ConnectDB(&database.DBConfig{Database_URL: config.Database.URL}); err != nil {
		log.Fatal("Failed to connect to Postgres DB:", err)
		log.Fatal("Critical Error: Shutting down application due to database connection failure.")
		os.Exit(1)

	}
	defer database.DisConnectDB()

	if database.DB == nil {
		log.Fatal("Failed to connect to Postgres DB:")
		log.Fatal("Critical Error: Shutting down application due to database connection failure.")
		os.Exit(1)
	}

	// miggrate models
	if err := migration.MigrateModels(database.DB); err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	app.Use(middlewares.CorsMiddleWare())
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	app.StaticFile("/swagger.json", filepath.Join(config.Server.DocumentDir, "swagger.json"))
	app.Static("/images", filepath.Join(config.Server.StaticDir, "img"))
	app.StaticFile("/favicon.ico", filepath.Join(config.Server.StaticDir, "img/favicon.ico"))
	app.MaxMultipartMemory = config.Server.MaxMultipartMemory << 20

	routers.SetupRouters(app, database.DB)

	// Listen and Serve
	if err := app.Run(*addr); err != nil {
		log.Fatal(err.Error())
	}
}
