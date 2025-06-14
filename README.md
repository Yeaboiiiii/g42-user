# G42-USERR

## 1. Project Overview

A Golang backend that handles JWT authentication and exposes endpoints such as `/login` and `/logout`. User authentication data is stored and validated via MongoDB.

---

## 2. Technology Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: Gin Web Framework
- **HTTP Router**: Gin Router
- **CORS**: `gin-contrib/cors` middleware
- **Database**: MongoDB
- **JSON Handling**: `encoding/json` package

### MongoDB Details
- **Connection String**: `mongodb://localhost:27017/`
- **Database Name**: `g42-user`
- **Collection Name**: `auth`

---

## 3. API Overview

### Available Endpoints

| Endpoint   | Method | Description               |
|------------|--------|---------------------------|
| `/login`   | POST   | User login using JWT      |
| `/logout`  | GET    | Invalidate JWT / Logout   |

---

## 4. System Architecture

```
┌────────────────────────┐    HTTP/REST   ┌────────────────────┐
│     Go Backend API     │◄──────────────►│    External APIs   │
│   (Gin + MongoDB + JWT)│                │ (e.g., client apps)│
│     Port: 8002         │                │                    │
└────────────────────────┘                └────────────────────┘
```

---

## 5. Project Structure

```
.
├── main.go                 # Entry point of application
├── go.mod                 # Go module definition
│
├── cmd/
│   ├── handler/           # HTTP route handlers
│   │   └── user_handler.go
│   └── logic/             # Core business logic
│       └── user_logic.go
│
├── repositories/          # MongoDB operations
│   └── user_repository.go
│
├── utils/                 # Utility functions (e.g., JWT utils)
│   └── jwt_utils.go
```

---

## 6. API Specification

### POST `/login`

- **Purpose**: Authenticates user credentials, returns JWT token
- **Request Body**:
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```
- **Response**:
```json
{
  "token": "your.jwt.token.here"
}
```

---

### GET `/logout`

- **Purpose**: Invalidate or clear token on client side
- **Note**: Server-side logout is stateless (JWT is not stored), so actual token invalidation depends on token expiration or client-side handling.

---

## 7. Run the Project

### Prerequisites

- Go installed (v1.18+ recommended)
- MongoDB running on `localhost:27017`

### Steps

```bash
# clone the repository
git clone https://github.com/Yeaboiiiii/g42-user.git
cd g42-user

# install dependencies
go mod tidy

# run the server
go run main.go
```

The server will start on `http://localhost:8080`.

---

## 8. Notes

- JWT secret should be stored securely (e.g., via environment variables).
- Passwords must be hashed before storing into MongoDB.
- Add middleware to protect private routes using JWT verification.