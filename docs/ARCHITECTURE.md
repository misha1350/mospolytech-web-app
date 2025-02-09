# System Architecture Overview

```mermaid
graph TD
    subgraph Frontend [Frontend - Vue.js]
        VR[Vue Router]
        VX[Vuex Store]
        VC[Vue Components]
        VC -->|State Management| VX
        VC -->|Navigation| VR
        VX -->|Auth State| VC
    end

    subgraph Backend [Backend - Go]
        GR[Gin Router]
        MW[Middleware]
        DB[(MySQL)]
        GR -->|Auth/Validation| MW
        MW -->|Queries| DB
    end

    subgraph Security [Security Layer]
        JWT[JWT Tokens]
        CORS[CORS]
        XSS[XSS Protection]
        CSRF[CSRF Protection]
    end

    Frontend -->|HTTP Requests| Backend
    Backend -->|JWT in HTTPOnly Cookie| Frontend
    Security -->|Protects| Frontend
    Security -->|Protects| Backend

```

## Component Communication Flow

1. **Authentication Flow**
   ```mermaid
   sequenceDiagram
       participant User
       participant Frontend
       participant Backend
       participant Database
       
       User->>Frontend: Login Form
       Frontend->>Backend: POST /api/server/login
       Backend->>Database: Validate Credentials
       Database-->>Backend: User Details
       Backend->>Backend: Generate JWT
       Backend-->>Frontend: Set HTTPOnly Cookie
       Frontend->>Frontend: Update Vuex Store
       Frontend-->>User: Redirect to Home
   ```

2. **Protected Route Access**
   ```mermaid
   sequenceDiagram
       participant User
       participant Router
       participant Store
       participant Backend
       
       User->>Router: Access Admin Route
       Router->>Store: Check User Role
       Store->>Backend: Validate Token
       Backend-->>Store: Token Valid/Invalid
       Store-->>Router: Allow/Deny Access
       Router-->>User: Show Page/Redirect
   ```

## Key Technologies

### Frontend Stack
- **Vue 3**: Core framework
- **Vuex**: State management
- **Vue Router**: Client-side routing
- **TailwindCSS**: Utility-first styling
- **Vite**: Build tool and dev server

### Backend Stack
- **Go**: Core language
- **Gin**: Web framework
- **JWT**: Authentication
- **MySQL**: Database
- **SQLC**: Type-safe SQL

## Security Implementation

### Token Flow
```mermaid
flowchart LR
    Login[Login Request] -->|Credentials| Auth[Auth Handler]
    Auth -->|Generate| JWT[JWT Token]
    JWT -->|Store| Cookie[HTTPOnly Cookie]
    Cookie -->|Send| Client[Client]
    Client -->|Include| Requests[API Requests]
    Requests -->|Validate| Middleware[Auth Middleware]
```

### Data Flow
```mermaid
flowchart TD
    Client[Client Request] -->|JWT| API[API Endpoint]
    API -->|Validate| Auth[Auth Middleware]
    Auth -->|Check| Role[Role Permission]
    Role -->|Query| DB[Database]
    DB -->|Response| API
    API -->|JSON| Client
```

## Directory Structure Explained

```
mospolytech-web-app/
├── client/                 # Frontend Vue application
│   ├── src/
│   │   ├── components/     # Reusable Vue components
│   │   ├── views/         # Page components
│   │   ├── store/         # Vuex state management
│   │   └── routes.js      # Vue Router configuration
│   └── ...
├── server/                 # Backend Go application
│   ├── middleware/        # Request processing
│   ├── config/           # Application configuration
│   ├── db/              # Database management
│   └── ...
└── docs/                  # Documentation
    ├── API.md            # API specifications
    ├── SECURITY.md       # Security implementation
    └── API_EXAMPLES.md   # API usage examples
```

## DevOps Considerations

### Environment Variables
Located in `.env`:
```env
# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_db_name

# JWT Configuration
JWT_SECRET=your_jwt_secret

# Server Configuration
SERVER_PORT=8086
CLIENT_PORT=8087
```

### Docker Configuration
The application is containerized with separate services:
- Frontend container (Vue)
- Backend container (Go)
- Database container (MySQL)

### Development Workflow
1. Local development with hot-reload
2. Testing in containerized environment
3. Production deployment with optimized builds

## Common Development Tasks

### Adding New Features
1. Create backend endpoint in `server/middleware/`
2. Add route in `server/main.go`
3. Create frontend component in `client/src/components/`
4. Add route in `client/src/routes.js`
5. Update documentation in `docs/`

### Security Updates
1. Update JWT configuration
2. Modify CORS settings
3. Add security headers
4. Update cookie settings
5. Document changes

This architecture overview should help you understand the system's components and their interactions. The mermaid diagrams provide a visual representation of the flows, making it easier to grasp the system's behavior.