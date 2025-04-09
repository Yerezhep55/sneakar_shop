package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneaker_shop/config"
	"sneaker_shop/models"
)

func GetSneakers(c *gin.Context) {
	var sneakers []models.Sneaker
	config.DB.Find(&sneakers)
	c.JSON(http.StatusOK, sneakers)
}

func GetSneakerByID(c *gin.Context) {
	id := c.Param("id")
	var sneaker models.Sneaker
	if err := config.DB.First(&sneaker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кроссовки не найдены"})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

// Создать новую пару
func CreateSneaker(c *gin.Context) {
	var sneaker models.Sneaker
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&sneaker)
	c.JSON(http.StatusCreated, sneaker)
}

// Обновить кроссовки по ID
func UpdateSneaker(c *gin.Context) {
	id := c.Param("id")
	var sneaker models.Sneaker
	if err := config.DB.First(&sneaker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кроссовки не найдены"})
		return
	}

	var updatedSneaker models.Sneaker
	if err := c.ShouldBindJSON(&updatedSneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sneaker.Brand = updatedSneaker.Brand
	sneaker.Models = updatedSneaker.Models
	sneaker.Size = updatedSneaker.Size
	sneaker.Price = updatedSneaker.Price

	config.DB.Save(&sneaker)
	c.JSON(http.StatusOK, sneaker)
}

// Удалить кроссовки по ID
func DeleteSneaker(c *gin.Context) {
	id := c.Param("id")
	var sneaker models.Sneaker
	if err := config.DB.First(&sneaker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кроссовки не найдены"})
		return
	}
	config.DB.Delete(&sneaker)
	c.JSON(http.StatusOK, gin.H{"message": "Кроссовки удалены"})
}
