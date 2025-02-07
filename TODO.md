# Project Improvements TODO List

## Bonus
- [ ] Replace hard-coded values with variables according to the DevOps good practices
  - `server/middleware/userEdit.go`: The code uses hardcoded values for `Roleid`, `Officeid` and `ID` in the `UpdateUserParams`. These should be replaced with variables.

## S-Tier (Critical Security & Stability Issues)
- [x] Fix database connection pooling - current implementation uses a global DB variable and doesn't properly manage connections
  - `server/middleware/dbConnect.go`: The code uses `sql.Open` to open a database connection and sets connection pool parameters using `SetMaxOpenConns`, `SetMaxIdleConns`, and `SetConnMaxLifetime`. This addresses the database connection pooling issue.
- [x] Add proper error handling for database connection failures instead of `log.Fatal`
  - `server/middleware/dbConnect.go`: The code includes error handling for database connection failures using `db.PingContext` and returns an error if the connection fails. The `initDB` function returns an error, which is then handled by `GetDB`. This addresses the error handling issue.
- [ ] Implement secure session management - current JWT implementation lacks proper expiration handling
  - `server/main.go`: The JWT implementation lacks proper expiration handling. The cookie is set to expire in 30 days, but the token itself doesn't seem to have an expiration. This is a security vulnerability.
- [ ] Add input validation and sanitization for all user inputs to prevent SQL injection
  - `server/middleware/register.go`: The code does not have explicit input validation and sanitization for all user inputs. It only parses the office ID and birthdate. This is a critical security vulnerability.
  - `server/middleware/userEdit.go`: The code binds the JSON data to the `editedUserInfo` struct, but it doesn't have any explicit input validation and sanitization. This is a security vulnerability. The `Office` and `Role` fields are received as strings but are not validated before being used.
- [ ] Add CSRF protection for forms
  - `server/main.go`: The code does not have CSRF protection for forms. This is a security vulnerability.
- [ ] Move environment variables to a secure configuration management system
  - `server/.env`: The code uses a `.env` file to store environment variables. This is not a secure configuration management system. In a production environment, it's recommended to use a more secure system like HashiCorp Vault or AWS Secrets Manager.
- [ ] Add rate limiting for authentication endpoints
  - `server/middleware/register.go`: The code does not have rate limiting for the registration endpoint. This can lead to abuse.
- [ ] Implement proper password validation rules
  - `server/middleware/register.go`: The code uses bcrypt to hash the password, which is good. However, it doesn't have any password validation rules (e.g., minimum length, complexity).
- [ ] Add database migration system for schema changes
  - `server/config/database.go`: The code does not have a database migration system. This is important for managing schema changes in a controlled and repeatable way.

## A-Tier (Important Functionality & Performance)
- [ ] Implement proper logging system with levels (error, info, debug) instead of fmt.Println
- [ ] Add database transaction handling for critical operations
- [ ] Implement proper API versioning
- [ ] Add request context timeout handling
- [ ] Implement proper error responses with consistent format
- [ ] Add database query optimization and indexing
- [ ] Implement caching layer for frequently accessed data
- [ ] Add proper metrics collection for monitoring
- [ ] Implement proper cleanup for expired sessions/tokens

## B-Tier (Code Quality & Maintainability)
- [ ] Restructure project to follow standard Go project layout
- [ ] Add comprehensive API documentation using Swagger/OpenAPI
- [ ] Implement proper dependency injection instead of global variables
- [ ] Add unit tests for critical components
- [ ] Add integration tests for API endpoints
- [ ] Implement proper interface abstractions for database operations
- [ ] Add code documentation for complex business logic
- [ ] Create proper development and production configurations

## C-Tier (Frontend Improvements)
- [ ] Implement proper state management in Vue components
- [ ] Add proper loading states for async operations
- [ ] Implement proper error handling and display
- [ ] Add form validation on the client side
- [ ] Improve responsive design for mobile devices
- [ ] Add proper TypeScript types for API responses
- [ ] Implement proper route guards for protected routes
- [ ] Add progressive loading for large data sets

## D-Tier (User Experience)
- [ ] Add proper feedback messages for user actions
- [ ] Implement proper form auto-completion
- [ ] Add remember-me functionality for login
- [ ] Implement proper navigation breadcrumbs
- [ ] Add proper form validation feedback
- [ ] Implement proper loading indicators
- [ ] Add proper error recovery flows
- [ ] Implement proper success/error notifications

## E-Tier (Nice-to-Have Features)
- [ ] Add dark mode support
- [ ] Implement multi-language support
- [ ] Add export functionality for data
- [ ] Implement user preferences storage
- [ ] Add keyboard shortcuts for common actions
- [ ] Implement proper print layouts
- [ ] Add accessibility features
- [ ] Implement data backup functionality

## F-Tier (Future Considerations)
- [ ] Add support for OAuth providers
- [ ] Implement real-time updates using WebSocket
- [ ] Add support for file attachments
- [ ] Implement audit logging
- [ ] Add support for different grading systems
- [ ] Implement automated reports generation
- [ ] Add support for academic calendar integration
- [ ] Implement student attendance tracking

## Notes for Student Marks System Adaptation
- Current "Office" concept can be repurposed for "Department" or "Faculty"
- "Users" table needs to be extended with student-specific fields
- New tables needed for:
  - Courses
  - Assignments
  - Grades
  - Academic Terms
  - Student Groups
  - Attendance Records