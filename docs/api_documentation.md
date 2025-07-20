# Task Manager API
This is an API to perform basic CRUD operations on a task list with persistent data storage
utilizing MongoDB.

It allows clients to

- Create new tasks

- Update existing tasks

- View lists of tasks

- View details of a specific task

- Delete tasks

## Prerequisites
- Golang version 1.24.5 or later
- MongoDB version 8.0.11 or later

## Dependencies
- github.com/gin-gonic/gin - The Gin Framework
- github.com/joho/godotenv - Godotenv for environment variable management
- go.mongodb.org/mongo-driver - MongoDB driver for Go

## Usage
1. Clone the github repository
```
git clone https://github.com/abeni-al7/task_manager
```
2. Get into the directory
```
cd task_manager
```
3. Configure environment variables with your database URI and HOST URL
```
mv .env.example .env
```
Get into .env and edit the default values with your own credentials
3. Run the server 
```
go run main.go
```

## Task Manager API Documentation
### GET Tasks
### http://localhost:8080/tasks/

#### Example Request
```curl --location 'http://localhost:8080/tasks/'```
#### Example Response
```
{
  "tasks": [
    {
      "id": "6878eb6ddfbd2f90f0d2c60a",
      "title": "not urgent",
      "description": "good but far",
      "due_date": "2025-12-16T08:30:00Z",
      "status": "pending",
      "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
      "UpdatedAt": "2025-07-16T11:51:41.028011929+03:00"
    }
  ]
}
```

### GET Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```
{
  "id": "6878eb6ddfbd2f90f0d2c60a"1,
  "title": "not urgent",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "pending",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T11:51:41.028011929+03:00"
}
```

### PUT Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location --request PUT 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a' \
--data '{
    "title": "Updated Title",
    "status": "completed"
}'
```
#### Example Response
```
{
  "id": "6878eb6ddfbd2f90f0d2c60a",
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

### POST Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location 'http://localhost:8080/tasks/' \
--data '{
    "title": "not urgent",
    "description": "good but far",
    "due_date": "2025-12-16T08:30:00Z",
    "status": "pending"
}'
```
#### Example Response
```
{
  "id": "6878eb6ddfbd2f90f0d2c60a",
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

### DELETE Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location --request DELETE 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```
Status code: 204
```

### GET Users
### http://localhost:8080/users/

#### Example Request
```curl --location 'http://localhost:8080/users/'```
#### Example Response
```
{
    "users": [
        {
            "id": "687ce54433fd48459614ca4e",
            "username": "heisenberg",
            "role": "admin",
            "email": "abebe@abeb.co",
            "created_at": "2025-07-20T12:47:00.633Z",
            "updated_at": "2025-07-20T12:47:00.633Z"
        },
        {
            "id": "687ce5ab33fd48459614ca4f",
            "username": "joe",
            "role": "regular",
            "email": "abebe@abeb.co",
            "created_at": "2025-07-20T12:48:43.095Z",
            "updated_at": "2025-07-20T12:48:43.095Z"
        }
    ]
}
```

### GET User
### http://localhost:8080/users/:id

#### Example Request
```
curl --location 'http://localhost:8080/users/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```
{
    "id": "687ce5ab33fd48459614ca4f",
    "username": "joe",
    "role": "regular",
    "email": "abebe@abeb.co",
    "created_at": "2025-07-20T12:48:43.095Z",
    "updated_at": "2025-07-20T12:48:43.095Z"
}
```

### PUT User
### http://localhost:8080/users/:id

#### Example Request
```
curl --location --request PUT 'http://localhost:8080/users/6878eb6ddfbd2f90f0d2c60a' \
--data '{
    "email": "updated@email.co",
    "username": "completed"
}'
```
#### Example Response
```
{
  "id": "6878eb6ddfbd2f90f0d2c60a",
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

### POST Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location 'http://localhost:8080/tasks/' \
--data '{
    "title": "not urgent",
    "description": "good but far",
    "due_date": "2025-12-16T08:30:00Z",
    "status": "pending"
}'
```
#### Example Response
```
{
  "id": "6878eb6ddfbd2f90f0d2c60a",
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

### DELETE Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location --request DELETE 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```
Status code: 204
```
