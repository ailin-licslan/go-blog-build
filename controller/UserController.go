package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-build/config"
	"go-blog-build/models"
	"net/http"
)

// RegisterUser 处理用户注册请求
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 假设这里有数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser 处理用户登录请求
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 假设这里有数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var storedUser models.User
	if err := db.Where("username =? AND password =?", user.Username, user.Password).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := config.GenerateJWT(storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
