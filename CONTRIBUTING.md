# Руководство по участию в разработке

## Настройка среды разработки

### Необходимые условия
- Go 1.x
- Node.js 16+
- MySQL 8.0+
- Docker и Docker Compose (опционально, для разработки в контейнерах)

### Настройка окружения

1. Клонируйте репозиторий
2. Скопируйте `.env.example` в `.env` и обновите значения:
   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=your_db_name
   JWT_SECRET=your_jwt_secret
   ```

### Настройка Backend

1. Установите Go зависимости:
   ```bash
   cd server
   go mod download
   ```

2. Настройте базу данных:
   - Импортируйте схему из `server/db/schema.sql`
   - Импортируйте начальные данные из `server/db/data_dump.sql`

3. Запустите сервер:
   ```bash
   go run main.go
   ```

Сервер будет доступен по адресу http://localhost:8086

### Настройка Frontend

1. Установите Node зависимости:
   ```bash
   cd client
   npm install
   ```

2. Запустите сервер разработки:
   ```bash
   npm run dev
   ```

Frontend будет доступен по адресу http://localhost:8087

## Структура проекта

### Backend (`/server`)
- `/config` - Загрузка и управление конфигурацией
- `/db` - Схемы базы данных и миграции
- `/middleware` - Промежуточное ПО для обработки запросов
  - `auth.go` - Аутентификация и авторизация
  - `admin.go` - Эндпоинты только для администраторов
  - `register.go` - Регистрация пользователей
  - `userEdit.go` - Управление пользователями

### Frontend (`/client`)
- `/src`
  - `/components` - Переиспользуемые Vue компоненты
  - `/views` - Компоненты страниц
  - `/store` - Управление состоянием Vuex
  - `/constants` - Общие константы и типы
  - `routes.js` - Определения маршрутов и guards

## Руководство по разработке

### Аутентификация
- Все защищенные эндпоинты должны проверять JWT токен
- Используйте компонент ErrorBoundary для обработки ошибок
- Управление токенами происходит в хранилище Vuex
- Всегда используйте HTTP-only cookies для токенов

### Обработка ошибок
1. Backend:
   - Используйте правильные коды состояния HTTP
   - Возвращайте согласованный формат ответа об ошибке
   - Логируйте ошибки с соответствующей детализацией

2. Frontend:
   - Используйте компонент ErrorBoundary
   - Обрабатывайте сетевые ошибки корректно
   - Отображайте понятные пользователю сообщения об ошибках

### Управление состоянием
1. Vuex Store:
   - Состояние аутентификации пользователя
   - Предпочтение темного режима
   - Глобальные состояния загрузки

2. Состояние компонента:
   - Данные формы
   - Состояния загрузки
   - Локальное состояние UI

### Стилизация
- Используйте утилитарные классы TailwindCSS
- Следуйте шаблону темного режима: `class="text-gray-900 dark:text-white"`
- Используйте предопределенные классы кнопок: `btn btn-primary`
- Используйте классы форм: `form-input`, `form-select`

### Тестирование
Перед отправкой PR:
1. Убедитесь, что все эндпоинты правильно аутентифицированы
2. Проверьте валидацию формы
3. Проверьте функциональность темного режима
4. Проверьте адаптивность для мобильных устройств
5. Проверьте обработку ошибок

## Распространенные задачи

### Добавление нового защищенного маршрута
1. Добавьте маршрут в `client/src/routes.js`
2. Добавьте мета-требования (например, `requiresAdmin`)
3. Создайте компонент представления
4. Обновите навигацию, если необходимо

### Добавление нового API Endpoint
1. Создайте обработчик в соответствующем файле middleware
2. Добавьте маршрут в `server/main.go`
3. Задокументируйте в `docs/API.md`
4. Добавьте обработку ошибок
5. Добавьте проверку аутентификации

### Изменение пользовательских данных
1. Обновите схему базы данных, если необходимо
2. Измените соответствующие обработчики
3. Обновите формы frontend
4. Добавьте валидацию
5. Обновите документацию