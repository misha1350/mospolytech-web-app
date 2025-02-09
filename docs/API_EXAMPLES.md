# Примеры API

Этот документ содержит примеры реальных запросов и ответов API для разработки и тестирования.

## Схема аутентификации

### 1. Запрос на вход
```http
POST /api/server/login
Content-Type: application/x-www-form-urlencoded

email=admin@example.com&password=yourpassword
```

### Ответ
```http
HTTP/1.1 301 Moved Permanently
Location: http://localhost:8087/client
Set-Cookie: Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...; Path=/; Domain=localhost; SameSite=Lax; HttpOnly
```

### 2. Проверка токена
```http
POST /api/server/check
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Успешный ответ
```json
{
  "userDetails": {
    "email": "admin@example.com",
    "firstname": "Admin",
    "lastname": "User",
    "office": 1,
    "role": "2"
  }
}
```

### Ответ с ошибкой
```json
{
  "error": "invalid token"
}
```

## Примеры управления пользователями

### Получить список пользователей
```http
GET /api/server/get_users
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Успешный ответ
```json
{
  "users": [
    {
      "id": 1,
      "role": "2",
      "email": "admin@example.com",
      "firstname": "Admin",
      "lastname": "User",
      "office": 1,
      "active": true
    },
    {
      "id": 2,
      "role": "1",
      "email": "user@example.com",
      "firstname": "Regular",
      "lastname": "User",
      "office": 2,
      "active": true
    }
  ]
}
```

### Редактировать пользователя
```http
PUT /api/server/user_edit
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "id": "2",
  "email": "user@example.com",
  "firstname": "Updated",
  "lastname": "User",
  "office": "2",
  "role": "1"
}
```

### Успешный ответ
```json
{
  "message": "User updated successfully"
}
```

### Ответ с ошибкой
```json
{
  "error": "Failed to update user"
}
```

## Общие коды состояния HTTP

- 200 OK: Успешный запрос
- 201 Created: Ресурс успешно создан
- 301 Moved Permanently: Перенаправление (используется после входа)
- 400 Bad Request: Неверные входные данные
- 401 Unauthorized: Отсутствует или недействительный токен
- 403 Forbidden: Недостаточно прав (например, не-админ пытается получить доступ к маршрутам администратора)
- 404 Not Found: Ресурс не найден
- 500 Internal Server Error: Ошибка на стороне сервера

## Формат ответа об ошибке
Все ответы об ошибках имеют следующий формат:
```json
{
  "error": "Читаемое сообщение об ошибке"
}
```

## Тестовые токены

Для целей разработки вы можете использовать эти примеры JWT:

### Пользователь-администратор
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcwMDAwMDAwMH0.signature
```

### Обычный пользователь
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjIsImV4cCI6MTcwMDAwMDAwMH0.signature
```

Примечание: это примеры токенов, и они не будут работать в production. Создавайте реальные токены через endpoint входа.