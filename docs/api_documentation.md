# Task Manager API
This is an API to perform basic CRUD operations on an in-memory task list.

It allows clients to

- Create new tasks

- Update existing tasks

- View lists of tasks

- View details of a specific task

- Delete tasks

## GET Tasks
### http://localhost:8080/tasks/

#### Example Request
```curl --location 'http://localhost:8080/tasks/'```
#### Example Response
```
{
  "tasks": [
    {
      "id": 1,
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

## GET Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location 'http://localhost:8080/tasks/1'
```
#### Example Response
```
{
  "id": 1,
  "title": "not urgent",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "pending",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T11:51:41.028011929+03:00"
}
```

## PUT Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location --request PUT 'http://localhost:8080/tasks/1' \
--data '{
    "title": "Updated Title",
    "status": "completed"
}'
```
#### Example Response
```
{
  "id": 1,
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

## POST Task
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
  "id": 1,
  "title": "Updated Title",
  "description": "good but far",
  "due_date": "2025-12-16T08:30:00Z",
  "status": "completed",
  "CreatedAt": "2025-07-16T11:51:41.028011851+03:00",
  "UpdatedAt": "2025-07-16T14:36:57.945307958+03:00"
}
```

## DELETE Task
### http://localhost:8080/tasks/:id

#### Example Request
```
curl --location --request DELETE 'http://localhost:8080/tasks/1'
```
#### Example Response
```
Status code: 204
```
