# Contributing Guide

## Development Setup

### Prerequisites
- Go 1.x
- Node.js 16+
- MySQL 8.0+
- Docker and Docker Compose (optional, for containerized development)

### Environment Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and update the values:
   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=your_db_name
   JWT_SECRET=your_jwt_secret
   ```

### Backend Setup

1. Install Go dependencies:
   ```bash
   cd server
   go mod download
   ```

2. Set up the database:
   - Import schema from `server/db/schema.sql`
   - Import initial data from `server/db/data_dump.sql`

3. Start the server:
   ```bash
   go run main.go
   ```

Server will be available at http://localhost:8086

### Frontend Setup

1. Install Node dependencies:
   ```bash
   cd client
   npm install
   ```

2. Start the development server:
   ```bash
   npm run dev
   ```

Frontend will be available at http://localhost:8087

## Project Structure

### Backend (`/server`)
- `/config` - Configuration loading and management
- `/db` - Database schemas and migrations
- `/middleware` - Request handling middleware
  - `auth.go` - Authentication and authorization
  - `admin.go` - Admin-only endpoints
  - `register.go` - User registration
  - `userEdit.go` - User management

### Frontend (`/client`)
- `/src`
  - `/components` - Reusable Vue components
  - `/views` - Page components
  - `/store` - Vuex state management
  - `/constants` - Shared constants and types
  - `routes.js` - Route definitions and guards

## Development Guidelines

### Authentication
- All protected endpoints must validate the JWT token
- Use the ErrorBoundary component for error handling
- Token management happens in the Vuex store
- Always use HTTP-only cookies for tokens

### Error Handling
1. Backend:
   - Use proper HTTP status codes
   - Return consistent error response format
   - Log errors with appropriate detail

2. Frontend:
   - Use the ErrorBoundary component
   - Handle network errors gracefully
   - Show user-friendly error messages

### State Management
1. Vuex Store:
   - User authentication state
   - Dark mode preference
   - Global loading states

2. Component State:
   - Form data
   - Loading states
   - Local UI state

### Styling
- Use TailwindCSS utility classes
- Follow dark mode pattern: `class="text-gray-900 dark:text-white"`
- Use predefined button classes: `btn btn-primary`
- Use form classes: `form-input`, `form-select`

### Testing
Before submitting a PR:
1. Ensure all endpoints are properly authenticated
2. Test form validation
3. Verify dark mode functionality
4. Check mobile responsiveness
5. Validate error handling

## Common Tasks

### Adding a New Protected Route
1. Add route in `client/src/routes.js`
2. Add meta requirements (e.g., `requiresAdmin`)
3. Create view component
4. Update navigation if needed

### Adding a New API Endpoint
1. Create handler in appropriate middleware file
2. Add route in `server/main.go`
3. Document in `docs/API.md`
4. Add error handling
5. Add authentication check

### Modifying User Data
1. Update database schema if needed
2. Modify relevant handlers
3. Update frontend forms
4. Add validation
5. Update documentation