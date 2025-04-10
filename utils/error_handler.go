package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleDBError 处理数据库连接错误
func HandleDBError(c *gin.Context, err error) {
	log.Println("Database connection error:", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
}

// HandleAuthError 处理用户认证失败错误
func HandleAuthError(c *gin.Context) {
	log.Println("User authentication failed")
	c.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
}

// HandleResourceNotFoundError 处理文章或评论不存在的错误
func HandleResourceNotFoundError(c *gin.Context, resource string) {
	log.Printf("%s not found\n", resource)
	c.JSON(http.StatusNotFound, gin.H{"error": resource + " not found"})
}
