API documentation:

---

# API Documentation

## Overview

This API allows you to manage Engineers, Projects, and Tasks. It includes the following key features:

- **Engineers**: Each engineer can belong to one project at a time.
- **Projects**: A project can have multiple engineers and tasks.
- **Tasks**: Each task is associated with a project and can be assigned to one engineer.

## Setup Instructions

### 1. **Clone the Repository**

Clone the repository to your local environment.

### 2. **Configure Docker**

Ensure you have Docker installed. Set up the Docker configuration as specified in the provided Docker file to containerize your application.

### 3. **Create Database**

Use PostgreSQL as the database. Make sure the database is created before running the API.

### 4. **Set Up Environment Variables**

In the `.env.example` file, the required environment variables are specified. Copy this file to a new file named `.env` and fill in your database credentials and configuration:

- `POSTGRES_HOST=your_host`
- `POSTGRES_USER=your_user`
- `POSTGRES_PASSWORD=your_password`
- `POSTGRES_NAME=your_dbname`
- `POSTGRES_PORT=your_port`
- `POSTGRES_SSL=disable`

**Note**: The `.env` file is not included in the repository for security reasons.

### 5. **Install Dependencies**

Install the required Go modules, including Gin for routing and GORM for database operations:

```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 6. **Run Migrations**

Run the database migrations to set up the required tables for Engineers, Projects, and Tasks.

### 7. **Run the Application**

Start the application using the following command:

```
go run main.go
```

### 8. **Database Connection**

The application connects to a PostgreSQL database using credentials from the `.env` file.

## Endpoints

### 1. **Create a Project**

- **Endpoint**: `POST /projects`
- **Request Body**:
  ```json
  {
    "name": "Project Alpha",
    "delivery_date": "2024-12-31"
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Project Alpha",
    "delivery_date": "2024-12-31"
  }
  ```

### 2. **Get All Projects**

- **Endpoint**: `GET /projects`
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Project Alpha"
    },
    {
      "id": 2,
      "name": "Project Beta"
    }
  ]
  ```

### 3. **Get Project By ID**

- **Endpoint**: `GET /projects/:id`
- **Response**:

  ```json
  [
    {
      "id": 1,
      "name": "Project Alpha",
      "delivery_date": "2024-12-31",
      "engineers": [
        {
          "id": 1,
          "name": "John",
          "surname": "Doe"
        }
      ],
      "tasks": [
        {
          "id": 1,
          "title": "Create database",
          "description": "Setting up the project database"
        }
      ]
    }
  ]
  ```

  ```

  ```

### 4. **Create an Engineer**

- **Endpoint**: `POST /engineers`
- **Request Body**:
  ```json
  {
    "name": "John",
    "surname": "Doe",
    "project_id": 1
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John",
    "surname": "Doe",
    "project_id": 1
  }
  ```

### 5. **Get Engineer by ID (Including Current Project)**

- **Endpoint**: `GET /engineers/:id`
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John",
    "surname": "Doe",
    "project": {
      "id": 1,
      "name": "Project Alpha"
    }
  }
  ```

### 6. **Reassign Engineer to a Different Project**

- **Endpoint**: `PUT /engineers/:id/project`
- **Request Body**:
  ```json
  {
    "project_id": 2
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John",
    "surname": "Doe",
    "project_id": 2
  }
  ```

### 7. **Create a Task**

- **Endpoint**: `POST /tasks`
- **Request Body**:
  ```json
  {
    "title": "Task Title",
    "description": "Task description",
    "estimated_hours": 8,
    "project_id": 1
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "title": "Task Title",
    "description": "Task description",
    "estimated_hours": 8,
    "project_id": 1
  }
  ```

## Additional Information

- **API Structure**: The API has models defined in the `models` directory for `Project`, `Engineer`, and `Task`.
- **Routes**: The routes for handling requests are organized in the `routes.go` files within their respective folders (`projects`, `engineer`, `tasks`).
- **Request Examples**: There is a `requests.http` file to test API requests.

**Go Version**: The project uses Go version `1.22.5`.

---
