ЭНДПОИНТЫ API

1. АУТЕНТИФИКАЦИЯ
   POST /auth/login — вход в систему

Пример запроса:

powershell
$loginData = @{ username="admin"; password="admin123" } | ConvertTo-Json
Invoke-RestMethod "http://localhost:8080/auth/login" -Method Post `
-Body $loginData -ContentType "application/json"
Ответ:

json
{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
"expires": "2024-01-20T12:00:00Z",
"user": "admin"
}
2. ЗАЩИЩЁННЫЕ ЭНДПОИНТЫ (требуется токен)
   GET /api/data

powershell
$headers = @{ Authorization = "Bearer ваш_jwt_токен" }
Invoke-RestMethod "http://localhost:8080/api/data" -Headers $headers
Ответ:

json
{
"message": "Добро пожаловать в защищенную зону!",
"user": "admin",
"data": "Это ваши защищенные данные"
}
POST /api/data — защита от XSS

powershell
$newData = @{ data = "<script>alert('XSS')</script>Ваш текст" } | ConvertTo-Json
Invoke-RestMethod "http://localhost:8080/api/data" -Method Post `
-Headers $headers -Body $newData -ContentType "application/json"
Ответ:

json
{
"message": "Данные успешно добавлены",
"original": "<script>alert('XSS')</script>Ваш текст",
"escaped_data": "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;Ваш текст"
}
3. ПУБЛИЧНЫЕ ЭНДПОИНТЫ
   GET / — информация о сервисе

powershell
Invoke-RestMethod "http://localhost:8080/" -Method Get
МЕРЫ ЗАЩИТЫ
1. ЗАЩИТА ОТ SQL-ИНЪЕКЦИЙ
   Используются параметризованные запросы:

go
stmt, _ := db.Prepare("SELECT * FROM users WHERE username = ? AND password = ?")
stmt.Query(username, hashedPassword)
2. ЗАЩИТА ОТ XSS
   go
   escaped := html.EscapeString(request.Data)
   Ввод:

text
<script>alert('attack')</script>
Вывод:

text
&lt;script&gt;alert(&#39;attack&#39;)&lt;/script&gt;
3. БЕЗОПАСНАЯ АУТЕНТИФИКАЦИЯ (bcrypt + JWT)
   ✔ bcrypt

go
bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
✔ JWT токены

go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString(jwtKey)
4. MIDDLEWARE АУТЕНТИФИКАЦИИ
   go
   if authHeader == "" {
   c.JSON(401, gin.H{"error": "Требуется авторизация"})
   c.Abort()
   }
   CI/CD PIPELINE И SECURITY-СКАНЕРЫ
   Используется GitHub Actions

🛠 Используемые security-инструменты:
SAST — Gosec

yaml
- name: Run Gosec Security Scanner
  uses: securego/gosec@master
  with:
  args: -fmt=html -out=gosec-report.html ./...
  Go Vulnerability Scanner — govulncheck

bash
govulncheck -show=traces ./... > govulncheck-report.txt
Snyk (SCA)

bash
snyk test --file=go.mod --package-manager=gomodules --json > snyk-report.json
Quality — golangci-lint

bash
golangci-lint run ./...

📊 Отчёты CI/CD

тут будут

ТЕСТИРОВАНИЕ API
powershell
# 1. Информация о API
Invoke-RestMethod "http://localhost:8080/" -Method Get

# 2. Аутентификация
$loginData = @{username="admin";password="admin123"} | ConvertTo-Json
$auth = Invoke-RestMethod "http://localhost:8080/auth/login" -Method Post `
-Body $loginData -ContentType "application/json"

# 3. Доступ к защищенным данным
$headers = @{Authorization="Bearer $($auth.token)"}
Invoke-RestMethod "http://localhost:8080/api/data" -Headers $headers