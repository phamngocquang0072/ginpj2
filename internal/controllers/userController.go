package controllers

import (

	"github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/ginpj2/initializers"
	"github.com/phamngocquang0072/ginpj2/internal/models"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func init() {
	// Load .env file
	initializers.loadEnv()
	DB = initializers.ConnectDB()
}


func GetUser(c *gin.Context)  {
	users := []models.User{}
	DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}