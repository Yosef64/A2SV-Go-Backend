# Task Management API Documentation

Base URL: `http://localhost:8080`

---

### GET /tasks

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

### GET /tasks/:id

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

### POST /tasks

**Description**: Create a new task.

**Request**:

```json
{
  "title": "New Task",
  "description": "Task details",
  "due_date": "2024-04-15",
  "status": "pending"
}
```

**Response**: 201 Created with the full task including generated id.

### PUT /tasks/:id

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

**Response**: 200 OK with updated task data.

### DELETE /tasks/:id

**Description**: Delete a task.

**Response**: 204 No Content

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
