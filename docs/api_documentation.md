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
```bash
git clone https://github.com/abeni-al7/task_manager
```
2. Get into the directory
```bash
cd task_manager
```
3. Configure environment variables with your database URI and HOST URL
```bash
mv .env.example .env
```
Get into .env and edit the default values with your own credentials
3. Run the server 
```bash
go run main.go
```

## Task Manager API Documentation
For the APIs which are protected, use "bearer xxxxxxxxxxxx" on the authorization header with your JWT token which expires after 24 hours and need to be generated vial login.

You can find the postman API documentation at: https://documenter.getpostman.com/view/46775407/2sB34ijKAe

### GET Tasks (open for all users)
### http://localhost:8080/tasks/

#### Example Request
```bash
curl --location 'http://localhost:8080/tasks/'
```
#### Example Response
```bash
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

### GET Task (open for all users)
### http://localhost:8080/tasks/:id

#### Example Request
```bash
curl --location 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```bash
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

### PUT Task (admin previledge)
### http://localhost:8080/tasks/:id

#### Example Request
```bash
curl --location --request PUT 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a' \
--data '{
    "title": "Updated Title",
    "status": "completed"
}'
```
#### Example Response
```bash
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

### POST Task (admin previledge)
### http://localhost:8080/tasks/:id

#### Example Request
```bash
curl --location 'http://localhost:8080/tasks/' \
--data '{
    "title": "not urgent",
    "description": "good but far",
    "due_date": "2025-12-16T08:30:00Z",
    "status": "pending"
}'
```
#### Example Response
```bash
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

### DELETE Task (admin previledge)
### http://localhost:8080/tasks/:id

#### Example Request
```bash
curl --location --request DELETE 'http://localhost:8080/tasks/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```bash
Status code: 204
```

### GET Users (admin previledge)
### http://localhost:8080/users/

#### Example Request
```bash
curl --location 'http://localhost:8080/users/'
```
#### Example Response
```bash
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

### GET User (account owner previledge)
### http://localhost:8080/users/:id

#### Example Request
```bash
curl --location 'http://localhost:8080/users/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```bash
{
    "id": "6878eb6ddfbd2f90f0d2c60a",
    "username": "joe",
    "role": "regular",
    "email": "abebe@abeb.co",
    "created_at": "2025-07-20T12:48:43.095Z",
    "updated_at": "2025-07-20T12:48:43.095Z"
}
```

### PUT User (account owner previledge)
### http://localhost:8080/users/:id

#### Example Request
```bash
curl --location --request PUT 'http://localhost:8080/users/6878eb6ddfbd2f90f0d2c60a' \
--data '{
    "email": "updated@email.co",
}'
```
#### Example Response
```bash
{
    "id": "6878eb6ddfbd2f90f0d2c60a",
    "username": "heisenberg",
    "role": "admin",
    "email": "h@h.co",
    "created_at": "2025-07-20T12:47:00.633Z",
    "updated_at": "2025-07-20T13:08:48.492Z"
}
```

### POST Login (anyone can access this one)
### http://localhost:8080/login

#### Example Request
```bash
curl --location 'http://localhost:8080/login/' \
--data '{
    "username": "heisenberg",
    "password": "testpass"
}'
```
#### Example Response
```bash
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiZWJlQGFiZWIuY28iLCJleHAiOjE3NTMxMDI1MDAsInJvbGUiOiJhZG1pbiIsInVzZXJfaWQiOiI2ODdjZTU0NDMzZmQ0ODQ1OTYxNGNhNGUiLCJ1c2VybmFtZSI6ImhlaXNlbmJlcmcifQ.jRKtw8UG0jNT20Yf4wjZS9kQTo9-WAEfSH4kmxXXI_M"
}
```

### POST Register (anyone can access this one)
### http://localhost:8080/register

#### Example Request
```bash
curl --location 'http://localhost:8080/register/' \
--data '{
    "username": "heisenberg",
    "email": "h@h.co",
    "password": "testpass"
}'
```
#### Example Response
```bash
{
    "id": "687ce5ab33fd48459614ca4f",
    "username": "heisenberg",
    "role": "regular", // the first to register is an admin
    "email": "h@h.co",
    "created_at": "2025-07-20T15:48:43.095064617+03:00",
    "updated_at": "2025-07-20T15:48:43.095064663+03:00"
}
```

### POST Promote (admin previledge)
### http://localhost:8080/promote/:id

#### Example Request
```bash
curl --location 'http://localhost:8080/promote/687ce5ab33fd48459614ca4f' \
```
#### Example Response
```bash
{
    "id": "687ce5ab33fd48459614ca4f",
    "username": "heisenberg",
    "role": "admin", // promoted
    "email": "h@h.co",
    "created_at": "2025-07-20T15:48:43.095064617+03:00",
    "updated_at": "2025-07-20T15:48:43.095064663+03:00"
}
```

### POST Change-Password (account owner previledge)
### http://localhost:8080/change-password/:id

#### Example Request
```
curl --location 'http://localhost:8080/687ce5ab33fd48459614ca4f/change-password' \
--data '{
    "prev_password": "test",
    "new_password": "mytpass",
}'
```
#### Example Response
```
{"message":"password updated successfully"}
```

### DELETE User (admin previledge)
### http://localhost:8080/users/:id

#### Example Request
```bash
curl --location --request DELETE 'http://localhost:8080/users/6878eb6ddfbd2f90f0d2c60a'
```
#### Example Response
```bash
Status code: 204
```

## Architecture
The project is structured in the following format
```bash
.
├── Delivery
│   ├── controllers
│   │   ├── task_controller.go
│   │   └── user_controller.go
│   ├── main.go
│   └── router
│       └── router.go
├── Domain
│   └── domain.go
├── Infrastructure
│   ├── auth_middleware.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories
│   ├── task_repository.go
│   └── user_repository.go
├── Usecases
│   ├── task_usecases.go
│   └── user_usecases.go
├── docs
│   └── api_documentation.md
├── go.mod
└── go.sum
```

This format of structuring a project is called the Clean architecture and the different layers of this project in accordance with the Clean architecture principle are the following:

- Delivery: This is the layer which contains logic for the mod of delivery of the Task manager project. The mode of delivery implemented here is REST APIs which can be consumed by client applications. For that there are routers and controllers. The controllers interact with the usecases to deliver the functionality needed.

- Domain: This is the layer which contains the core data structures and system wide implementations that feed every other layer in the architecture. In our case the domain contains the 2 core models for this project namely task and user. It also contains the database connection object as well as input structs. The 2 core structs for this project are:
```Golang
type Task struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	DueDate time.Time `bson:"due_date" json:"due_date"`
	Status string `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Role string `bson:"role" json:"role"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
```

- Infrastructure: This layer contains the implementation for the security features that support the whole project. It has JWT token generation and validation, password hashing and validation as well as authorization middlewares to limit access to secure routes.

- Repositories: This layer contains the logic for the database interaction this project would have to perform the use cases. It currently supports the mongoDB database and it can be swapped with any other database if needed.

- Usecases: This layer contains the core business logic of the application. This layer is agnostic towards the framework used or the database utilized. It supports data validation and communicates with an interface that is implemented by the repositories layer for database functionality. It also depends on an interface that is implemented by the infrastructure layer for security features.

This approach was chosen for structuring this project because it allows flexibility in responding to changes in requirement or tools that could arise later during the maintenance of the project.

## Testing
The project contains mocks that implement the repository interfaces so that usecases can be tested separately.
The project contains unittests for usecases covering happy paths as well as error paths. The tests utilize their separate suits to avoid repetition when initializing the mock repository and the usecase to be tested.