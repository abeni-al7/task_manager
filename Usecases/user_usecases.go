package usecases

import (
	"errors"
	"time"

	domain "github.com/abeni-al7/task_manager/Domain"
	infrastructure "github.com/abeni-al7/task_manager/Infrastructure"
	repositories "github.com/abeni-al7/task_manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

func NewUserUsecase(ur repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu *UserUsecase) Register(user *domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	user.Password = hashedPassword

	*user, err = uu.userRepo.Register(user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return *user, nil
}

func (uu *UserUsecase) Login(username string, password string) (string, error) {
	jwtToken, err := uu.userRepo.Login(username, password)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return jwtToken, nil
}

func (uu *UserUsecase) Promote(id primitive.ObjectID) (domain.User, error) {
	user, err := uu.userRepo.Promote(id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (uu *UserUsecase) FetchAll() ([]domain.User, error) {
	users, err := uu.userRepo.FetchAll()
	if err != nil {
		return []domain.User{}, errors.New(err.Error())
	}
	return users, nil
}

func (uu *UserUsecase) Fetch(id primitive.ObjectID) (domain.User, error) {
	user, err := uu.userRepo.Fetch(id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (uu *UserUsecase) Update(id primitive.ObjectID, updatedUser domain.User) (domain.User, error) {
	user, err := uu.userRepo.Update(id, updatedUser)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (uu *UserUsecase) ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error {
	existingUser, err := uu.userRepo.Fetch(id)
	if err != nil {
		return errors.New("user does not exist")
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(prevPassword)) != nil {
		return errors.New("incorrect password")
	}

	err = uu.userRepo.ChangePassword(id, prevPassword, newPassword)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (uu *UserUsecase) Remove(id primitive.ObjectID) error {
	user, err := uu.userRepo.Fetch(id)
	if err != nil {
		return errors.New("user did not exist")
	}

	if user.Role == "admin" {
		return errors.New("admin cannot be deleted")
	}

	err = uu.userRepo.Remove(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}