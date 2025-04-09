package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sneaker_shop/models"
)

var DB *gorm.DB // 👈 эта переменная должна быть видимой и глобальной

func ConnectDB() {
	dsn := "host=localhost user=postgres password=Topi2005 dbname=sneakers_shop port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Не удалось подключиться к базе данных:", err)
	}

	err = db.AutoMigrate(&models.Sneaker{})
	if err != nil {
		log.Fatal("❌ Миграция не удалась:", err)
	}

	DB = db // 👈 эта строка обязательно должна быть
	log.Println("✅ Успешно подключено к базе данных!")
}
