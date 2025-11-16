🔒 Secure REST API на Go

Добро пожаловать в документацию по защищённому REST API, разработанному в рамках лабораторной работы по безопасности веб-приложений.
API реализует аутентификацию с использованием JWT, защиту данных и меры против наиболее распространённых атак.

📌 Эндпоинты API
1. 🔑 Аутентификация

POST /auth/login — вход в систему

Пример запроса (PowerShell):

$loginData = @{ username="admin"; password="admin123" } | ConvertTo-Json

Invoke-RestMethod "http://localhost:8080/auth/login" -Method Post `
-Body $loginData -ContentType "application/json"


Ответ:

{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
"expires": "2024-01-20T12:00:00Z",
"user": "admin"
}

2. 🔐 Защищённые эндпоинты (требуется JWT)
   GET /api/data
   $headers = @{ Authorization = "Bearer ваш_jwt_токен" }

Invoke-RestMethod "http://localhost:8080/api/data" -Headers $headers


Ответ:

{
"message": "Добро пожаловать в защищенную зону!",
"user": "admin",
"data": "Это ваши защищенные данные"
}

POST /api/data — защита от XSS
$newData = @{ data = "<script>alert('XSS')</script>Ваш текст" } | ConvertTo-Json

Invoke-RestMethod "http://localhost:8080/api/data" -Method Post `
-Headers $headers -Body $newData -ContentType "application/json"


Ответ:

{
"message": "Данные успешно добавлены",
"original": "<script>alert('XSS')</script>Ваш текст",
"escaped_data": "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;Ваш текст"
}

3. 🌍 Публичные эндпоинты
   GET /
   Invoke-RestMethod "http://localhost:8080/" -Method Get

🛡 Меры защиты
1. 🚫 SQL-инъекции

Используются параметризованные запросы:

stmt, _ := db.Prepare("SELECT * FROM users WHERE username = ? AND password = ?")
stmt.Query(username, hashedPassword)

2. 🧼 Защита от XSS
   escaped := html.EscapeString(request.Data)


Ввод:

<script>alert('attack')</script>


Вывод:

&lt;script&gt;alert(&#39;attack&#39;)&lt;/script&gt;

3. 🔐 Безопасная аутентификация (bcrypt + JWT)

✔ bcrypt
bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))

✔ JWT-токены
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString(jwtKey)

4. 🧩 Middleware аутентификации

if authHeader == "" {
c.JSON(401, gin.H{"error": "Требуется авторизация"})
c.Abort()
}

⚙️ CI/CD и Security-сканеры

Проект использует GitHub Actions и следующие security-инструменты:

🧪 SAST — Gosec
- name: Run Gosec Security Scanner
  uses: securego/gosec@master
  with:
  args: -fmt=html -out=gosec-report.html ./...

🔎 Go Vulnerability Scanner — govulncheck
govulncheck -show=traces ./... > govulncheck-report.txt

🧬 Snyk (SCA)
snyk test --file=go.mod --package-manager=gomodules --json > snyk-report.json

🔍 Quality — golangci-lint
golangci-lint run ./...

📊 Отчёты CI/CD

(будут добавлены позже)

🧪 Тестирование API
# 1. Информация о API
Invoke-RestMethod "http://localhost:8080/" -Method Get

# 2. Аутентификация
$loginData = @{username="admin";password="admin123"} | ConvertTo-Json
$auth = Invoke-RestMethod "http://localhost:8080/auth/login" -Method Post `
-Body $loginData -ContentType "application/json"

# 3. Доступ к защищённым данным
$headers = @{Authorization="Bearer $($auth.token)"}
Invoke-RestMethod "http://localhost:8080/api/data" -Headers $headers