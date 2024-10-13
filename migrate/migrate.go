package main

import (
	"github.com/phamngocquang0072/ginpj1/initializers"
	"github.com/phamngocquang0072/ginpj1/internal/models"
	"log"
	"gorm.io/gorm"
)


var DB *gorm.DB

func init() {
	// Load .env file
	initializers.loadEnv()
	DB = initializers.ConnectDB()
}

func main() {
	DB.AutoMigrate(&models.Employee{})
	log.Fatal("Migrated!")
}
