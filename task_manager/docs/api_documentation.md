# Task Management API Documentation

Base URL: `http://localhost:8080`

---

## Authentication

All task endpoints require authentication except login and register.
Clients must include a valid Authorization header:

```makefile
Authorization: Bearer <your-token>

```

If the token is invalid or missing, the server will respond with 401 Unauthorized.

### GET /register

**Authentication Not Required**
**Description**: Register a new user.
**Request**:

```json
{
  "username": "username",
  "password": "password",
  "role": "user or admin"
}
```

**Response:**

```json
{
  "message": "User registered successfully"
}
```

### POST /login

**Authentication Not Required**

**Description**: Login and receive a JWT token.
**Request**:

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Response**: 200 OK

```json
{
  "token": "your_jwt_token_here"
}
```

### GET /api/tasks

**Authentication Required**
**Description:** Fetch all tasks.

**Response:**

```json
[
  {
    "id": "uuid",
    "title": "Task Title",
    "description": "Description here",
    "due_date": "2024-04-10",
    "status": "pending"
  }
]
```

### GET /api/tasks/:id

**Authentication Required**
**Description:**: Fetch a specific task.

**Response:**

```json
{
  "id": "uuid",
  "title": "Task Title",
  "description": "Description here",
  "due_date": "2024-04-10",
  "status": "pending"
}
```

### POST /api/tasks

**Authentication Required**
**Description**: Create a new task owned by the authenticated user.

**Request**:

```json
{
  "title": "New Task",
  "description": "Task details",
  "due_date": "2024-04-15",
  "status": "pending"
}
```

**Response**: 201 Created.

```json
{
  "message": "Task created successfully"
}
```

### PUT /api/tasks/:id

**Authentication Required**
**Description**: Update a task.
**Request**:

```json
{
  "title": "Updated Title",
  "description": "Updated details",
  "due_date": "2024-04-20",
  "status": "done"
}
```

**Response**: 200 OK.

```json
{
  "message": "Task updated successfully"
}
```

### DELETE /api/admin/tasks/:id

**Authentication Required**
**Authorization**: Only the admin can delete it.
**Description**: Delete a task.

**Response**: 204.

```json
{
  "message": "Task deleted successfully"
}
```

### Error Codes

**_- 400 Bad Request_**: Invalid request payload.

**_-404 Not Found_**: Task not found.

---

### ✅ Run Instructions

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run main.go
```

## For more information

https://restless-comet-915812.postman.co/workspace/MyFirst~a6132409-17d4-48b1-99cc-40b7228d33ff/collection/32029702-6699d8c5-5400-4306-b94e-6eb8bf32cf04?action=share&creator=32029702
