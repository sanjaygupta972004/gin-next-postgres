package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// GlobalConfig holds the entire configuration
type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig holds server-related settings
type ServerConfig struct {
	Mode               string
	Port               string
	Version            string
	StaticDir          string
	DocumentDir        string
	MaxMultipartMemory int64
	SecurityKey        string
}

// DatabaseConfig holds database settings
type DatabaseConfig struct {
	URL string
}

var (
	configInstance *Configuration
	isLoaded       bool
	mu             sync.RWMutex
	configOnce     sync.Once
)

// Load environment variables from the .env file
func LoadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No .env file found or failed to load: %v", err)
		return err
	}
	return nil
}

// Load server configuration from environment variables
func LoadServerConfig() (ServerConfig, error) {
	serverConfig := ServerConfig{
		Mode:               getEnvOrDefault("GIN_MODE", "debug"),
		Port:               getEnvOrDefault("PORT", "8000"),
		Version:            getEnvOrDefault("VERSION", "1.0.0"),
		StaticDir:          getEnvOrDefault("STATIC_DIR", "./public"),
		DocumentDir:        getEnvOrDefault("DOCUMENT_DIR", "./docs"),
		MaxMultipartMemory: 8 * 1024 * 1024, // 8 MB
		SecurityKey:        getEnvOrDefault("SECURITY_KEY", ""),
	}

	// Validate required fields
	if serverConfig.SecurityKey == "" {
		return ServerConfig{}, fmt.Errorf("SECURITY_KEY is required but missing")
	}

	return serverConfig, nil
}

// Utility function to fetch environment variable with a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Load database configuration from environment variables
func LoadDatabaseConfig() (DatabaseConfig, error) {
	dbURL := getEnvOrDefault("DATABASE_URL", "")

	if dbURL == "" {
		return DatabaseConfig{}, fmt.Errorf("DATABASE_URL is required but missing")
	}

	return DatabaseConfig{URL: dbURL}, nil
}

// Load all configurations into the global instance
func LoadGlobalConfig() error {
	var loadError error

	configOnce.Do(func() {
		if err := LoadEnvFile(); err != nil {
			loadError = err
			return
		}

		serverConfig, err := LoadServerConfig()
		if err != nil {
			loadError = err
			return
		}

		databaseConfig, err := LoadDatabaseConfig()
		if err != nil {
			loadError = err
			return
		}

		mu.Lock()
		configInstance = &Configuration{
			Server:   serverConfig,
			Database: databaseConfig,
		}
		isLoaded = true
		mu.Unlock()

		log.Println("Configuration loaded successfully")
	})

	return loadError
}

// Get the global configuration instance
func GetGlobalConfig() *Configuration {
	mu.RLock()
	if isLoaded {
		defer mu.RUnlock()
		return configInstance
	}
	mu.RUnlock()

	// If not loaded, load it first
	LoadGlobalConfig()
	mu.RLock()
	defer mu.RUnlock()
	return configInstance
}
