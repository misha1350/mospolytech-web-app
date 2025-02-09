# Документация API

## Базовый URL
Все endpoints имеют префикс `/api/server/`

## Аутентификация

### Обзор
В приложении используется аутентификация на основе JWT (JSON Web Token). Токены хранятся в HTTP-only cookies и проверяются при каждом запросе. Схема аутентификации выглядит следующим образом:

1. Пользователь входит в систему через форму входа
2. Сервер проверяет учетные данные и выдает JWT токен
3. Токен сохраняется в HTTP-only cookie
4. Frontend автоматически включает этот cookie в последующие запросы
5. Сервер проверяет токен для каждого защищенного endpoint

### Endpoints

#### Login
- **URL**: `/api/server/login`
- **Метод**: `POST`
- **Данные формы**:
  ```
  email: string
  password: string
  ```
- **Ответ**:
  - Успех: Перенаправляет на `/client` с установленным cookie `Authorization`
  - Ошибка: Возвращает сообщение об ошибке с соответствующим HTTP статусом
- **Расположение кода**: `server/main.go:authenticationsHandler`
- **Детали реализации**:
  ```go
  // server/middleware/auth.go:GenerateToken
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "sub": user.ID,
      "exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
  })
  ```

#### Token Validation
- **URL**: `/api/server/check`
- **Метод**: `POST`
- **Заголовки**:
  ```
  Authorization: <token>
  ```
- **Ответ**:
  ```json
  {
    "userDetails": {
      "email": string,
      "firstname": string,
      "lastname": string,
      "office": number,
      "role": string
    }
  }
  ```
- **Расположение кода**: `server/middleware/auth.go:ValidateToken`
- **Использование во Frontend**: `client/src/App.vue:onMounted`

#### Logout
Реализован на стороне клиента путем:
1. Очистки cookie Authorization
2. Очистки состояния хранилища Vuex
3. Перенаправления на страницу входа

Расположение кода: `client/src/views/NavBar.vue:logout`

### Вопросы безопасности
1. Токены хранятся в HTTP-only cookies для предотвращения XSS атак
2. Проверка токена включает проверку срока действия
3. Защищенные маршруты требуют валидный токен
4. Реализован контроль доступа на основе ролей для admin маршрутов

## Другие Endpoints

### Управление пользователями

#### Get Users List
- **URL**: `/api/server/get_users`
- **Метод**: `GET`
- **Заголовки**: 
  ```
  Authorization: <token>
  ```
- **Ответ**:
  ```json
  {
    "users": [
      {
        "id": number,
        "role": string,
        "email": string,
        "firstname": string,
        "lastname": string,
        "office": number,
        "active": boolean
      }
    ]
  }
  ```
- **Расположение кода**: `server/middleware/admin.go:GetUsers`

#### Edit User
- **URL**: `/api/server/user_edit`
- **Метод**: `PUT`
- **Заголовки**:
  ```
  Authorization: <token>
  Content-Type: application/json
  ```
- **Тело запроса**:
  ```json
  {
    "email": string,
    "firstname": string,
    "lastname": string,
    "office": string,
    "role": string
  }
  ```
- **Расположение кода**: `server/middleware/userEdit.go:EditUser`

#### User Registration
- **URL**: `/api/server/register`
- **Метод**: `POST`
- **Данные формы**:
  ```
  email: string
  firstname: string
  lastname: string
  birthdate: string (YYYY-MM-DD)
  office: string
  password: string
  ```
- **Расположение кода**: `server/middleware/register.go:RegisterUser`

## Frontend-Backend Communication

Frontend (`client/`) взаимодействует с backend (`server/`) через следующие основные компоненты:

1. **Управление состоянием аутентификации**:
   - Расположение: `client/src/store/index.js`
   - Управляет пользовательской сессией и настройками темной темы
   - Предоставляет getters для роли пользователя и статуса аутентификации

2. **Route Guards**:
   - Расположение: `client/src/routes.js`
   - Защищает admin маршруты на основе роли пользователя
   - Перенаправляет на страницу входа, если токен недействителен

3. **API Calls**:
   - Все API calls включают токен Authorization из cookies
   - Обработка ошибок с помощью компонента ErrorBoundary
   - Единообразные сообщения об ошибках и состояния загрузки