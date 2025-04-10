package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-build/config"
	"go-blog-build/models"
	"net/http"
	"strconv"
)

// CreateArticle 处理文章创建请求
func CreateArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = uint(userID.(uint))
	// 数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	if db == nil {
		return
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": post})
}

// GetArticles 处理获取所有文章请求
func GetArticles(c *gin.Context) {
	// 数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var posts []models.Post
	//Preload 预加载User
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": posts})
}

// GetArticle 处理获取单个文章请求
func GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	// 数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var post models.Post
	//不展示User set User = nil
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": post})
}

// UpdateArticle 处理文章更新请求
func UpdateArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	// 数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	if post.UserID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this article"})
		return
	}
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": post})
}

// DeleteArticle 处理文章删除请求
func DeleteArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	// 数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	if post.UserID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this article"})
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
