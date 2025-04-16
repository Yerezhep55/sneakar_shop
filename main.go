package main

import (
	"github.com/gin-gonic/gin"
	"sneaker_shop/config"
	"sneaker_shop/controllers"
	"sneaker_shop/middleware"
	"sneaker_shop/routes"
)

func main() {
	// Инициализация конфигурации и подключения к базе данных
	if err := config.ConnectDatabase(); err != nil {
		panic("Не удалось подключиться к базе данных: " + err.Error())
	}

	// Создание нового экземпляра Gin
	r := gin.Default()

	// Регистрация маршрутов для работы с кроссовками
	routes.RegisterSneakerRoutes(r)

	// Регистрация маршрутов для аутентификации
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Защищенные маршруты с JWT аутентификацией
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		// Эти маршруты доступны только авторизованным пользователям
		authorized.GET("/profile", controllers.GetUserProfile)
	}

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
