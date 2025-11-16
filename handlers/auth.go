package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"secure-lab/models"
)

// Секретный ключ для подписи JWT, в реальном приложении мы бы хранили его в безопасном месте
var jwtKey = []byte("my_secret_key")

// LoginRequest структура для входа
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Claims структура для JWT токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login обрабатывает аутентификацию пользователя
func Login(c *gin.Context) {
	var loginReq LoginRequest

	//Получаем данные из запроса
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	//Ищем пользователя, в реальном приложении это была бы база данных
	var user models.User
	found := false
	for _, u := range models.Users {
		if u.Username == loginReq.Username {
			user = u
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	//Проверяем пароль
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}

	//Создаем JWT токен
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: loginReq.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания токена"})
		return
	}

	//Возвращаем токен
	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"expires": expirationTime,
		"user":    loginReq.Username,
	})
}
