package usecases

import (
	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseInterface interface {
	Register(user *domain.User) (domain.User, error)
	Login(username string, password string) (string, error)
	Promote(id primitive.ObjectID) (domain.User, error)
	FetchAll() ([]domain.User, error)
	Fetch(id primitive.ObjectID) (domain.User, error)
	Update(id primitive.ObjectID, updatedUser domain.User) (domain.User, error)
	ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error
	Remove(id primitive.ObjectID) error
}

type UserUsecase struct {
	userRepo repositories.UserRepository
}

func (uu *UserUsecase) Register(user *domain.User) (domain.User, error) {
	
}