package main

import (
	"fmt"
	"gin-skeleton/config"
	"os"
)

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}

func main() {
	fmt.Printf("🚀 Server is running on port: %v\n", config.Global.Server.Port)
}
