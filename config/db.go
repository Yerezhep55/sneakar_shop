package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sneaker_shop/models"
)

var DB *gorm.DB

// ConnectDatabase - подключение к базе данных
func ConnectDatabase() error {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env файл не найден. Используется строка подключения по умолчанию")
	}

	// Получаем строку подключения
	dsn := os.Getenv("DB_URI")
	if dsn == "" {
		// fallback — если .env нет, использовать хардкод
		dsn = "host=localhost user=postgres password=Topi2005 dbname=sneakers_shop port=5432 sslmode=disable TimeZone=Asia/Almaty"
	}

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Миграция моделей
	err = db.AutoMigrate(&models.Sneaker{}, &models.User{})
	if err != nil {
		return err
	}

	DB = db
	log.Println("✅ Успешно подключено к базе данных!")
	return nil
}
