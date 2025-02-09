# API Documentation

## Base URL
All endpoints are prefixed with `/api/server/`

## Authentication

### Overview
The application uses JWT (JSON Web Token) based authentication. Tokens are stored in HTTP-only cookies and validated on each request. The authentication flow is as follows:

1. User logs in via the login form
2. Server validates credentials and issues a JWT token
3. Token is stored in an HTTP-only cookie
4. Frontend includes this cookie automatically in subsequent requests
5. Server validates the token on each protected endpoint

### Endpoints

#### Login
- **URL**: `/api/server/login`
- **Method**: `POST`
- **Form Data**:
  ```
  email: string
  password: string
  ```
- **Response**:
  - Success: Redirects to `/client` with `Authorization` cookie set
  - Error: Returns error message with appropriate HTTP status
- **Code Location**: `server/main.go:authenticationsHandler`
- **Implementation Details**:
  ```go
  // server/middleware/auth.go:GenerateToken
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "sub": user.ID,
      "exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
  })
  ```

#### Token Validation
- **URL**: `/api/server/check`
- **Method**: `POST`
- **Headers**:
  ```
  Authorization: <token>
  ```
- **Response**:
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
- **Code Location**: `server/middleware/auth.go:ValidateToken`
- **Frontend Usage**: `client/src/App.vue:onMounted`

#### Logout
Implemented client-side by:
1. Clearing the Authorization cookie
2. Clearing the Vuex store state
3. Redirecting to login page

Code Location: `client/src/views/NavBar.vue:logout`

### Security Considerations
1. Tokens are stored in HTTP-only cookies to prevent XSS attacks
2. Token validation includes expiration check
3. Protected routes require valid token
4. Role-based access control implemented for admin routes

## Other Endpoints

### User Management

#### Get Users List
- **URL**: `/api/server/get_users`
- **Method**: `GET`
- **Headers**: 
  ```
  Authorization: <token>
  ```
- **Response**:
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
- **Code Location**: `server/middleware/admin.go:GetUsers`

#### Edit User
- **URL**: `/api/server/user_edit`
- **Method**: `PUT`
- **Headers**:
  ```
  Authorization: <token>
  Content-Type: application/json
  ```
- **Body**:
  ```json
  {
    "email": string,
    "firstname": string,
    "lastname": string,
    "office": string,
    "role": string
  }
  ```
- **Code Location**: `server/middleware/userEdit.go:EditUser`

#### User Registration
- **URL**: `/api/server/register`
- **Method**: `POST`
- **Form Data**:
  ```
  email: string
  firstname: string
  lastname: string
  birthdate: string (YYYY-MM-DD)
  office: string
  password: string
  ```
- **Code Location**: `server/middleware/register.go:RegisterUser`

## Frontend-Backend Communication

The frontend (`client/`) communicates with the backend (`server/`) through these main components:

1. **Authentication State Management**:
   - Location: `client/src/store/index.js`
   - Manages user session and dark mode preferences
   - Provides getters for user role and authentication status

2. **Route Guards**:
   - Location: `client/src/routes.js`
   - Protects admin routes based on user role
   - Redirects to login if token is invalid

3. **API Calls**:
   - All API calls include the Authorization token from cookies
   - Error handling with the ErrorBoundary component
   - Consistent error messages and loading states