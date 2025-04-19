# Simple Book Management API

A RESTful API service built with Golang for managing books and categories with user authentication.

## Features

- User registration and login with JWT authentication
- CRUD operations for books
- Category management
- Secure password hashing with bcrypt
- SQLite database with GORM
- Clean, modular code structure

## Technologies Used

- Go (1.24.1)
- Gin Web Framework
- GORM
- SQLite
- JWT for authentication
- Bcrypt for password hashing

## API Endpoints

### Authentication

- `POST /register` - Register a new user
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- `POST /login` - Login and get JWT token
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

### Categories

- `GET /categories` - List all categories
- `POST /categories` - Create a new category
  ```json
  {
    "name": "string"
  }
  ```

### Books

- `GET /books` - List all books (optional query parameters: `?title=SampleBook` or `?category=Fanatsy`)
- `POST /books` - Create a new book (requires authentication)
  ```json
  {
    "title": "string",
    "author": "string",
    "category": "string"
  }
  ```
- `PUT /books/{id}` - Update a book (requires authentication)
- `DELETE /books/{id}` - Delete a book (requires authentication)

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/rishad004/SimpleBookManagementAPI.git
   cd SimpleBookManagementAPI
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```
4. Create a folder called `database`

5. The API will be available at `http://localhost:8080`

## Database Setup

The application uses SQLite and will automatically create a database file (`main.db`) inside the folder called `database` on first run.

## Sample Requests

Register a new user:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"testpass"}'
```

Login and get JWT token:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"testpass"}'
```

Create a book (with JWT):
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title":"SampleBook", "author":"John Doe", "category":"Fantasy"}'
```