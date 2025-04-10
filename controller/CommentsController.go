package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog-build/config"
	"go-blog-build/models"
	"net/http"
	"strconv"
)

// CreateComment 处理评论创建请求
func CreateComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	postIDStr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.UserID = uint(userID.(uint))
	comment.PostID = uint(postID)

	// 假设这里有数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetComments 处理获取某篇文章所有评论请求
func GetComments(c *gin.Context) {
	postIDStr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 假设这里有数据库连接，在实际应用中从 config.ConnectDB() 获取
	db := config.GetDB()
	var comments []models.Comment
	if err := db.Where("post_id =?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
