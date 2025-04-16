package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneaker_shop/config"
	"sneaker_shop/models"
)

// GetUserProfile возвращает профиль текущего пользователя
func GetUserProfile(c *gin.Context) {
	// Получаем user_id из контекста (добавляется middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить пользователя"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
