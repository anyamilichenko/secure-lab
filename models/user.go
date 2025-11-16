package models

// User представляет структуру пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Здесь будет храниться хеш пароля
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// Для простоты будем хранить пользователей в памяти
var Users = []User{
	{
		ID:       1,
		Username: "admin",
		Password: "$2a$10$examplehashedpassword", // Мы заменим это реальным хешем
		Email:    "admin@example.com",
		FullName: "Администратор Системы",
	},
	{
		ID:       2,
		Username: "user1",
		Password: "$2a$10$examplehashedpassword2",
		Email:    "user1@example.com",
		FullName: "Тестовый Пользователь",
	},
}
