package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTSecret 是用于签名 JWT 的密钥，应该在生产环境中妥善保管
const JWTSecret = "LIN-PLAY-GO-WEB-JWT"

// GenerateJWT 生成 JWT
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token 有效期为 24 小时
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// VerifyJWT 验证 JWT
func VerifyJWT(tokenString string) (uint, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}
	userID := uint(claims["user_id"].(float64))
	return userID, true
}
