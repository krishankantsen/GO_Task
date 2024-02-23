# CRUD Operations with SQL Database in Go

This repository demonstrates basic CRUD (Create, Read, Update, Delete) operations using an SQL database in Go. It utilizes the popular Gin framework for building web applications in Go and connects to an SQL database for data storage.

## Prerequisites

- Go installed on your machine
- An SQL database (e.g., SQLite, MySQL) and the necessary driver
- [Gin](https://github.com/gin-gonic/gin) framework installed (`go get -u github.com/gin-gonic/gin`)
- Any additional dependencies specified in `go.mod`

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/krishankantsen/GO_Task.git
   ```

2. Navigate to the project directory:
   ```bash
   cd GO_Task
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### 1. Create Task
   ```http
   POST /createTask
   ```

### 2. Get All Tasks
   ```http
   GET /getAllTask
   ```

### 3. Get Task by ID
   ```http
   GET /getTask/:id
   ```

### 4. Update Task by ID
   ```http
   PUT /updateTask/:id
   ```

### 5. Delete Task by ID
   ```http
   DELETE /deleteTask/:id
   ```

## Usage

- Use a tool like [Postman](https://www.postman.com/) to test the API endpoints.
- Modify the database schema and models as needed for your application.
- Customize the routes and handlers in `routes.go` based on your requirements.
