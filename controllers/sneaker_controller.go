package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneaker_shop/config"
	"sneaker_shop/models"
	"strconv"
)

func GetSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)
	offset := (pageInt - 1) * limitInt

	model := c.DefaultQuery("models", "")
	brand := c.DefaultQuery("brand", "")
	price := c.DefaultQuery("price", "")

	query := config.DB.Model(&models.Sneaker{})

	if model != "" {
		query = query.Where("models = ?", model)
	}
	if brand != "" {
		query = query.Where("brand = ?", brand)
	}
	if price != "" {
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err == nil {
			query = query.Where("price = ?", priceFloat)
		}
	}

	query.Limit(limitInt).Offset(offset).Find(&sneakers)

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

func CreateSneaker(c *gin.Context) {
	var sneaker models.Sneaker
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&sneaker)
	c.JSON(http.StatusCreated, sneaker)
}

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

func GetBrands(c *gin.Context) {
	var brands []string
	config.DB.Model(&models.Sneaker{}).Distinct().Pluck("brand", &brands)
	c.JSON(http.StatusOK, brands)
}

func GetModelsByBrand(c *gin.Context) {
	brand := c.Param("brand")
	var modelsList []string
	config.DB.Model(&models.Sneaker{}).Where("brand = ?", brand).Distinct().Pluck("models", &modelsList)
	c.JSON(http.StatusOK, modelsList)
}
