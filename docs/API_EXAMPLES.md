# API Examples

This document provides examples of actual API requests and responses for development and testing purposes.

## Authentication Flow

### 1. Login Request
```http
POST /api/server/login
Content-Type: application/x-www-form-urlencoded

email=admin@example.com&password=yourpassword
```

### Response
```http
HTTP/1.1 301 Moved Permanently
Location: http://localhost:8087/client
Set-Cookie: Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...; Path=/; Domain=localhost; SameSite=Lax; HttpOnly
```

### 2. Token Validation
```http
POST /api/server/check
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Success Response
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

### Error Response
```json
{
  "error": "invalid token"
}
```

## User Management Examples

### Get Users List
```http
GET /api/server/get_users
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Success Response
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

### Edit User
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

### Success Response
```json
{
  "message": "User updated successfully"
}
```

### Error Response
```json
{
  "error": "Failed to update user"
}
```

## Common HTTP Status Codes

- 200 OK: Successful request
- 201 Created: Resource successfully created
- 301 Moved Permanently: Redirect (used after login)
- 400 Bad Request: Invalid input data
- 401 Unauthorized: Missing or invalid token
- 403 Forbidden: Insufficient permissions (e.g., non-admin accessing admin routes)
- 404 Not Found: Resource not found
- 500 Internal Server Error: Server-side error

## Error Response Format
All error responses follow this format:
```json
{
  "error": "Human readable error message"
}
```

## Testing Tokens

For development purposes, you can use these example JWTs:

### Admin User
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTcwMDAwMDAwMH0.signature
```

### Regular User
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjIsImV4cCI6MTcwMDAwMDAwMH0.signature
```

Note: These are example tokens and won't work in production. Generate real tokens through the login endpoint.