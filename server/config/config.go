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
	Server    ServerConfig
	Database  DatabaseConfig
	AuthToken AuthToken
	AWS       AWSConfig
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

// AWS configuration veriable
type AWSConfig struct {
	Region         string
	BucketName     string
	SesSenderEmail string
}

// Load AWS configuration from environment variables
func LoadAWSConfig() (AWSConfig, error) {
	awsConfig := AWSConfig{
		BucketName:     getEnvOrDefault("AWS_BUCKET_NAME", ""),
		Region:         getEnvOrDefault("AWS_REGION", ""),
		SesSenderEmail: getEnvOrDefault("SES_SENDER_EMAIL", ""),
	}
	return awsConfig, nil
}

// DatabaseConfig holds database settings
type DatabaseConfig struct {
	URL string
}

// AuthToken holds authentication tokens
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

// load auth token secret from env
func LoadAuthToken() (AuthToken, error) {
	accessToken := getEnvOrDefault("ACCESS_TOKEN", "")
	refreshToken := getEnvOrDefault("REFRESH_TOKEN", "")

	if accessToken == "" || refreshToken == "" {
		return AuthToken{}, fmt.Errorf("ACCESS_TOKEN and REFRESH_TOKEN are required but missing")

	}
	return AuthToken{AccessToken: accessToken, RefreshToken: refreshToken}, nil
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

		awsConfig, err := LoadAWSConfig()
		if err != nil {
			loadError = err
			return
		}

		databaseConfig, err := LoadDatabaseConfig()
		if err != nil {
			loadError = err
			return
		}

		authToken, err := LoadAuthToken()
		if err != nil {
			loadError = err
			return
		}

		mu.Lock()
		configInstance = &Configuration{
			Server:    serverConfig,
			Database:  databaseConfig,
			AuthToken: authToken,
			AWS:       awsConfig,
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
