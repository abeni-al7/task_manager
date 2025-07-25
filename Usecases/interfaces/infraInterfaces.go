package usecases

import "github.com/abeni-al7/task_manager/Domain"

type IInfrastructure interface {
	HashPassword(password string) (string, error)
	ComparePassword(correctPassword []byte, inputPassword []byte) error
	GenerateJwtToken(user *domain.User) (string, error)
}