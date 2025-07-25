package usecases

import (
	"errors"
	"time"

	domain "github.com/abeni-al7/task_manager/Domain"
	usecases "github.com/abeni-al7/task_manager/Usecases/interfaces"
)

type UserUsecase struct {
	userRepo usecases.IUserRepo
	infra usecases.IInfrastructure
}

func NewUserUsecase(ur usecases.IUserRepo, infra usecases.IInfrastructure) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
		infra: infra,
	}
}

func (uu *UserUsecase) Register(user *domain.User) (domain.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return domain.User{}, errors.New("missing required fields")
	}

	_, err := uu.userRepo.FetchByUsername(user.Username)
	if err == nil {
		return domain.User{}, errors.New("user with this username already exists")
	}

	count, err := uu.userRepo.CountUsers()
	if err != nil {
		return domain.User{}, errors.New("unable to regiter user")
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "regular"
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := uu.infra.HashPassword(user.Password)
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
	if username == "" || password == "" {
		return "", errors.New("missing username or password")
	}

	existingUser, err := uu.userRepo.FetchByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = uu.infra.ComparePassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	jwtToken, err := uu.infra.GenerateJwtToken(&existingUser)
	if err != nil {
		return "", errors.New("unable to generate token")
	}

	return jwtToken, nil
}

func (uu *UserUsecase) Promote(id string) (domain.User, error) {
	user, err := uu.userRepo.Fetch(id)
	if err != nil {
		return domain.User{}, errors.New("user does not exist")
	}

	user, err = uu.userRepo.Promote(&user)
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

func (uu *UserUsecase) Fetch(id string) (domain.User, error) {
	user, err := uu.userRepo.Fetch(id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (uu *UserUsecase) Update(id string, updatedUser domain.User) (domain.User, error) {
	user, err := uu.userRepo.Update(id, updatedUser)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (uu *UserUsecase) ChangePassword(id string, prevPassword string, newPassword string) error {
	existingUser, err := uu.userRepo.Fetch(id)
	if err != nil {
		return errors.New("user does not exist")
	}

	if uu.infra.ComparePassword([]byte(existingUser.Password), []byte(prevPassword)) != nil {
		return errors.New("incorrect password")
	}

	err = uu.userRepo.ChangePassword(id, prevPassword, newPassword)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (uu *UserUsecase) Remove(id string) error {
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