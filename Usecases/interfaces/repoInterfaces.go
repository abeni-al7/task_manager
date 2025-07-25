package usecases

import (
	"github.com/abeni-al7/task_manager/Domain"
)

type IUserRepo interface {
	Register(user *domain.User) (domain.User, error)
	Promote(user *domain.User) (domain.User, error)
	FetchAll() ([]domain.User, error)
	Fetch(idStr string) (domain.User, error)
	Update(idStr string, updatedUser domain.User) (domain.User, error)
	ChangePassword(idStr string, prevPassword string, newPassword string) error
	Remove(idStr string) error
	FetchByUsername(username string) (domain.User, error)
	CountUsers() (int, error)
}

type ITaskRepo interface {
	Create(task *domain.Task) (domain.Task, error)
	FetchAll() ([]domain.Task, error)
	Fetch(idStr string ) (domain.Task, error)
	Update(idStr string, task domain.Task) (domain.Task, error)
	Remove(idStr string) error
}