# Quickstart: Daily Task Tracker

## Prerequisites
- Go 1.22+
- PostgreSQL
- `psql` command line tool (optional, for db init)

## Setup

1. **Initialize Database**:
   ```bash
   createdb simple_task_list
   ```

2. **Run Migrations / Schema Creation**:
   Since we don't have a complex migration tool, schema creation can be handled automatically on startup or via a simple init file.
   ```sql
   CREATE TABLE IF NOT EXISTS tasks (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       status VARCHAR(20) DEFAULT 'NOT_STARTED',
       estimated_duration INT DEFAULT 0,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

3. **Install Dependencies**:
   ```bash
   go mod init github.com/username/simpletasklist
   go get github.com/gorilla/mux
   go get github.com/lib/pq
   ```

## Running the Application

```bash
export DB_CONN="postgres://localhost/simple_task_list?sslmode=disable"
go run backend/main.go
```

The application will start the server at `http://localhost:8080`.

## Validation Scenario

1. Open `http://localhost:8080` in a browser.
2. The plain HTML interface should load.
3. Enter a task title (e.g., "Complete speckit workflow") and a duration in minutes (e.g., 30).
4. Submit the form.
5. The task should appear in the "Not Started" column.
6. Click the task or a button to move it to "In Progress".
7. Verify it instantly moves to the correct column without requiring a manual page refresh (via JS).
