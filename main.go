package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"secure-lab/handlers"
	"secure-lab/middleware"
	"secure-lab/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Хешируем пароль для тестового пользователя
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	models.Users[0].Password = string(hashedPassword)

	// Создаем роутер
	router := gin.Default()

	// Маршруты
	// 1. Аутентификация
	router.POST("/auth/login", handlers.Login)

	// 2. Защищенные маршруты (требуют JWT токен)
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/data", handlers.GetData)  // Получить данные
		api.POST("/data", handlers.AddData) // Добавить данные
	}

	// 3. Публичный маршрут
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Secure API на Go",
			"version": "1.0",
		})
	})

	// Запускаем сервер
	log.Println("Сервер запущен на http://localhost:8080")
	router.Run(":8080")
}
