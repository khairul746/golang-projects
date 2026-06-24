# API Contracts: Daily Task Tracker

## Base URL
`/api/v1`

## Endpoints

### 1. Get All Tasks
- **Method**: `GET`
- **Path**: `/tasks`
- **Response**: `200 OK`
  ```json
  [
    {
      "id": 1,
      "title": "Buy groceries",
      "status": "NOT_STARTED",
      "estimated_duration": 30,
      "created_at": "2026-06-10T10:00:00Z"
    }
  ]
  ```

### 2. Create Task
- **Method**: `POST`
- **Path**: `/tasks`
- **Request Body**:
  ```json
  {
    "title": "Read a book",
    "estimated_duration": 60
  }
  ```
- **Response**: `201 Created`
  ```json
  {
    "id": 2,
    "title": "Read a book",
    "status": "NOT_STARTED",
    "estimated_duration": 60,
    "created_at": "2026-06-10T10:05:00Z"
  }
  ```

### 3. Update Task Status
- **Method**: `PATCH`
- **Path**: `/tasks/{id}/status`
- **Request Body**:
  ```json
  {
    "status": "IN_PROGRESS"
  }
  ```
- **Response**: `200 OK`
  ```json
  {
    "id": 2,
    "title": "Read a book",
    "status": "IN_PROGRESS",
    "estimated_duration": 60,
    "updated_at": "2026-06-10T10:10:00Z"
  }
  ```
