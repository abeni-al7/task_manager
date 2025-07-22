package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserInput struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type TaskUsecase interface {
	Create(task *Task) (Task, error)
	FetchAll() ([]Task, error)
	Fetch(id primitive.ObjectID) (Task, error)
	Update(id primitive.ObjectID, updatedTask Task) (Task, error)
	Remove(id primitive.ObjectID) error
}

type TaskRepository interface {
	Create(task *Task) (Task, error)
	FetchAll() ([]Task, error)
	Fetch(id primitive.ObjectID) (Task, error)
	Update(id primitive.ObjectID, updatedTask Task) (Task, error)
	Remove(id primitive.ObjectID) error
}

type UserUsecase interface {
	Register(user *User) (User, error)
	Login(username string, password string) (string, error)
	Promote(id primitive.ObjectID) (User, error)
	FetchAll() ([]User, error)
	Fetch(id primitive.ObjectID) (User, error)
	Update(id primitive.ObjectID, updatedUser User) (User, error)
	ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error
	Remove(id primitive.ObjectID) error
}

type UserRepository interface {
	Register(user *User) (User, error)
	Login(username string, password string) (string, error)
	Promote(id primitive.ObjectID) (User, error)
	FetchAll() ([]User, error)
	Fetch(id primitive.ObjectID) (User, error)
	Update(id primitive.ObjectID, updatedUser User) (User, error)
	ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error
	Remove(id primitive.ObjectID) error
}