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

