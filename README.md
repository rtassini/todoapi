# Todo API

A RESTful API for managing todo items built with Go, Gin web framework, and PostgreSQL database. This project follows clean architecture principles with separation of concerns through domain, usecase, and infrastructure layers.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Project](#running-the-project)
  - [Local Development](#local-development)
  - [Docker Compose](#docker-compose)
  - [Kubernetes Deployment](#kubernetes-deployment)
- [API Endpoints](#api-endpoints)
- [Request and Response Examples](#request-and-response-examples)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Development](#development)
  - [Running Tests](#running-tests)
  - [Building the Project](#building-the-project)

## Overview

This Todo API provides a complete CRUD (Create, Read, Update, Delete) interface for managing todo items. It demonstrates:

- **Clean Architecture**: Separation between domain logic, use cases, and infrastructure
- **Dependency Injection**: All components are injected for better testability
- **Database Integration**: PostgreSQL with proper connection management
- **REST API**: Built with the Gin framework for high performance
- **Docker Support**: Containerization for easy deployment
- **Kubernetes Ready**: Includes deployment manifests for Kubernetes orchestration

## Architecture

The project follows a layered architecture:

```
┌─────────────────────────────────────┐
│      HTTP Controller Layer          │
│    (Gin Router & Handlers)          │
├─────────────────────────────────────┤
│      Application Layer (Usecase)    │
│    (Business Logic & Orchestration) │
├─────────────────────────────────────┤
│      Domain Layer                   │
│    (Core Business Logic & Models)   │
├─────────────────────────────────────┤
│      Infrastructure Layer           │
│    (Database & Repository)          │
└─────────────────────────────────────┘
```

### Layers Explained

- **Domain**: Pure business logic with no external dependencies (entities, value objects, business rules)
- **Use Cases**: Application logic that orchestrates domain objects and repositories
- **Infrastructure**: Technical concerns like database access and HTTP handling
- **Controller**: HTTP request/response handling and routing

## Prerequisites

- **Go 1.25+**: Programming language runtime
- **PostgreSQL 13+**: Relational database
- **Docker** (optional): For containerized development
- **Kubernetes** (optional): For orchestrated deployment

## Installation

### 1. Clone the Repository

```bash
cd /home/rtassini/GolandProjects/todoapi
```

### 2. Install Go Dependencies

```bash
go mod download
```

### 3. Set Up PostgreSQL

The project expects a PostgreSQL database named `go_db` with user `postgres` and password `postgres`.

#### Option A: Using Docker Compose (Recommended for Development)

```bash
docker-compose up -d
```

This will start a PostgreSQL container with the correct configuration.

#### Option B: Manual PostgreSQL Setup

```bash
createdb -U postgres go_db
```

The database will be created automatically by PostgreSQL. Ensure the `pg_hba.conf` file allows connections:

```
host    all             all             0.0.0.0/0               md5
```

## Running the Project

### Local Development

1. **Start PostgreSQL** (if not already running):

```bash
docker-compose up -d go_db
```

2. **Run the application**:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

3. **Test the API**:

```bash
curl http://localhost:8080/ping
# Response: {"message":"pong"}
```

### Docker Compose

Start the entire application stack (API + PostgreSQL) with a single command:

```bash
docker-compose up -d
```

To stop:

```bash
docker-compose down
```

To view logs:

```bash
docker-compose logs -f go_db
docker-compose logs -f
```

### Kubernetes Deployment

Deploy to a Kubernetes cluster:

```bash
# Create the ConfigMap with database configuration
kubectl apply -f configmap.yaml

# Deploy PostgreSQL
kubectl apply -f postgres-deployment.yaml

# Deploy the API
kubectl apply -f deployment.yaml

# Create service for the API
kubectl apply -f service.yaml
```

Access the API:

```bash
# Forward the service to localhost
kubectl port-forward svc/todo-api 8080:8080

# Test the API
curl http://localhost:8080/ping
```

To clean up:

```bash
kubectl delete -f configmap.yaml -f postgres-deployment.yaml -f deployment.yaml -f service.yaml
```

## API Endpoints

All endpoints return JSON responses. The base URL is `http://localhost:8080`.

### Health Check

#### `GET /ping`

Health check endpoint to verify the API is running.

**Response:**
```json
{
  "message": "pong"
}
```

---

### Todo Operations

#### `POST /todos` - Create a New Todo

Creates a new todo item with the given title and completion status.

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Buy groceries",
  "completed": false
}
```

**Response (200 OK):**
```json
{
  "response": {
    "id": "1",
    "title": "Buy groceries",
    "completed": false
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "todo invalid input: title"
}
```

**Validation Rules:**
- `title` is required and cannot be empty or whitespace-only
- `completed` is a boolean (true/false)

---

#### `GET /todos` - Get All Todos

Retrieves a list of all todo items in the database.

**Response (200 OK):**
```json
{
  "response": [
    {
      "id": "1",
      "title": "Buy groceries",
      "completed": false
    },
    {
      "id": "2",
      "title": "Finish project",
      "completed": true
    }
  ]
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "error": "fail to get all todos in the repository"
}
```

---

#### `PUT /todos/:id` - Update a Todo

Updates an existing todo item's title and completion status.

**URL Parameter:**
- `id` (required): The unique identifier of the todo item

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "Buy groceries at the store",
  "completed": true
}
```

**Response (200 OK):**
```json
{
  "response": {
    "message": "Rows updated: 1"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "todo invalid input: title"
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "error": "fail to update a todo in the repository"
}
```

---

#### `DELETE /todos/:id` - Delete a Todo

Deletes an existing todo item by its ID.

**URL Parameter:**
- `id` (required): The unique identifier of the todo item

**Response (200 OK):**
```json
{
  "response": {
    "message": "Rows deleted: 1"
  }
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "error": "fail to delete a todo in the repository"
}
```

---

## Request and Response Examples

### Example 1: Complete Todo Workflow

```bash
# 1. Create a new todo
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","completed":false}'

# Response:
# {"response":{"id":"1","title":"Learn Go","completed":false}}

# 2. Get all todos
curl http://localhost:8080/todos

# Response:
# {"response":[{"id":"1","title":"Learn Go","completed":false}]}

# 3. Update the todo to mark it as completed
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","completed":true}'

# Response:
# {"response":{"message":"Rows updated: 1"}}

# 4. Delete the todo
curl -X DELETE http://localhost:8080/todos/1

# Response:
# {"response":{"message":"Rows deleted: 1"}}
```

### Example 2: Using Python requests

```python
import requests
import json

BASE_URL = "http://localhost:8080"

# Create a todo
response = requests.post(
    f"{BASE_URL}/todos",
    json={"title": "Complete project report", "completed": False}
)
print(response.json())

# Get all todos
response = requests.get(f"{BASE_URL}/todos")
print(response.json())

# Update a todo
response = requests.put(
    f"{BASE_URL}/todos/1",
    json={"title": "Complete project report", "completed": True}
)
print(response.json())

# Delete a todo
response = requests.delete(f"{BASE_URL}/todos/1")
print(response.json())
```

### Example 3: Using JavaScript fetch

```javascript
const BASE_URL = "http://localhost:8080";

// Create a todo
fetch(`${BASE_URL}/todos`, {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ title: "Learn TypeScript", completed: false })
})
.then(res => res.json())
.then(data => console.log(data));

// Get all todos
fetch(`${BASE_URL}/todos`)
  .then(res => res.json())
  .then(data => console.log(data));

// Update a todo
fetch(`${BASE_URL}/todos/1`, {
  method: "PUT",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ title: "Learn TypeScript", completed: true })
})
.then(res => res.json())
.then(data => console.log(data));

// Delete a todo
fetch(`${BASE_URL}/todos/1`, { method: "DELETE" })
  .then(res => res.json())
  .then(data => console.log(data));
```

## Project Structure

```
todoapi/
├── main.go                              # Application entry point
├── go.mod                               # Go module definition
├── go.sum                               # Go dependencies checksums
├── Dockerfile                           # Docker image configuration
├── docker-compose.yml                   # Docker Compose configuration
├── deployment.yaml                      # Kubernetes deployment manifest
├── service.yaml                         # Kubernetes service manifest
├── postgres-deployment.yaml             # Kubernetes PostgreSQL deployment
├── configmap.yaml                       # Kubernetes configuration
├── postgres-data/                       # PostgreSQL data directory
│   └── pg_hba.conf                      # PostgreSQL authentication config
└── internal/
    ├── app/
    │   └── usecase/                     # Application use cases (business logic)
    │       ├── usecase.go               # Error handling utilities
    │       └── todo/
    │           ├── create.go            # Create todo use case
    │           ├── getall.go            # Get all todos use case
    │           ├── update.go            # Update todo use case
    │           ├── delete.go            # Delete todo use case
    │           ├── create_test.go       # Tests for create use case
    │           └── usecase_test.go      # Tests for use cases
    ├── domain/                          # Domain models (core business logic)
    │   ├── todo.go                      # Todo domain entity
    │   └── todo_test.go                 # Tests for domain logic
    └── infra/                           # Infrastructure layer
        ├── controller/
        │   └── todo.go                  # HTTP handlers for todo endpoints
        ├── db/
        │   └── conn.go                  # Database connection logic
        └── repository/
            ├── todo.go                  # Todo repository (data access)
            └── todo_mock.go             # Mock repository for testing
```

## Database Schema

The application uses a single `todos` table with the following schema:

```sql
CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  completed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Table Structure

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| `id` | SERIAL | PRIMARY KEY | Auto-incremented unique identifier |
| `title` | VARCHAR(255) | NOT NULL | The todo item title/description |
| `completed` | BOOLEAN | DEFAULT FALSE | Whether the todo is completed |
| `created_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Last update timestamp |

The database is automatically created when you start PostgreSQL with Docker Compose or manually with `docker-compose up`.

## Development

### Running Tests

Run all unit tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test ./... -v
```

Run tests for a specific package:

```bash
go test ./internal/domain -v
```

Run tests with coverage:

```bash
go test ./... -cover
```

Generate coverage report:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Building the Project

Build the binary for your current OS:

```bash
go build -o api
```

Run the compiled binary:

```bash
./api
```

Build for a specific OS and architecture:

```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o api-mac

# Windows
GOOS=windows GOARCH=amd64 go build -o api.exe
```

### Code Quality

Format code:

```bash
go fmt ./...
```

Analyze code for issues:

```bash
go vet ./...
```

### Troubleshooting

#### Connection Error: "no pg_hba.conf entry"

The PostgreSQL database is not accepting connections from your application IP.

**Solution**: Ensure `pg_hba.conf` contains:
```
host    all             all             0.0.0.0/0               md5
```

#### "Rows updated: 0" on PUT request

The todo ID doesn't exist in the database.

**Solution**: Verify the ID exists by running `GET /todos` first.

#### "Port 8080 already in use"

Another service is using port 8080.

**Solution**: Either:
- Stop the other service
- Modify the port in `main.go` (line: `server.Run(":8080")`)

## License

This project is provided as-is for educational and development purposes.

## Contact

For issues or questions, please refer to the project documentation or create an issue in the repository.
