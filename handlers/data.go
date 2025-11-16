package handlers

import (
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DataResponse структура для ответа с данными
type DataResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

// GetData возвращает защищенные данные
func GetData(c *gin.Context) {
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"message": "Добро пожаловать в защищенную зону!",
		"user":    username,
		"data":    "Это ваши защищенные данные",
	})
}

// AddData добавляет новые данные (защита от XSS)
func AddData(c *gin.Context) {
	var request struct {
		Data string `json:"data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// ЗАЩИТА ОТ XSS: экранируем HTML-символы
	escapedData := html.EscapeString(request.Data)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Данные успешно добавлены",
		"original":     request.Data,
		"escaped_data": escapedData, // Это безопасная версия
	})
}
