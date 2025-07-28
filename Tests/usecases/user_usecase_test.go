package tests

import (
	"errors"
	"testing"

	domain "github.com/abeni-al7/task_manager/Domain"
	usecases "github.com/abeni-al7/task_manager/Usecases"
	"github.com/abeni-al7/task_manager/Usecases/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockUserRepo
	usecase  usecases.UserUsecase
	mockinfra    *mocks.MockInfrastructure
}

func (suite *UserTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockUserRepo)
	suite.mockinfra = new(mocks.MockInfrastructure)
	suite.usecase = *usecases.NewUserUsecase(suite.mockRepo, suite.mockinfra)
}

func (suite *UserTestSuite) TestRegularUserRegister() {
	user := &domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("FetchByUsername", user.Username).Return(domain.User{}, errors.New("not found"))
	suite.mockRepo.On("CountUsers").Return(2, nil)
	suite.mockinfra.On("HashPassword", user.Password).Return("hashedpassword", nil)
	suite.mockRepo.On("Register", user).Return(*user, nil)

	createdUser, err := suite.usecase.Register(user)
	suite.NoError(err)
	suite.Equal("testuser", createdUser.Username)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestAdminUserRegister() {
	user := &domain.User{
		Username: "adminuser",
		Password: "adminpassword",
		Email:    "testadmin@example.com",
	}

	suite.mockRepo.On("FetchByUsername", user.Username).Return(domain.User{}, errors.New("not found"))
	suite.mockRepo.On("CountUsers").Return(0, nil)
	suite.mockinfra.On("HashPassword", user.Password).Return("hashedadminpassword", nil)
	suite.mockRepo.On("Register", user).Return(*user, nil)

	createdUser, err := suite.usecase.Register(user)
	suite.NoError(err)
	suite.Equal("adminuser", createdUser.Username)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestRegisterMissingFields() {
	user := &domain.User{
		Username: "",
		Password: "password123",
		Email:    "testuser@example.com",
	}
	
	createdUser, err := suite.usecase.Register(user)
	suite.Error(err)
	suite.Equal(domain.User{}, createdUser)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestDuplicateRegister() {
	user := &domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("FetchByUsername", user.Username).Return(*user, nil)
	createdUser, err := suite.usecase.Register(user)
	suite.Error(err)
	suite.Equal(domain.User{}, createdUser)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestUserLogin() {
	user := &domain.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "testuser@example.com",
	}

	hashedPassword := "hashedpassword"

	suite.mockRepo.On("FetchByUsername", user.Username).Return(*user, nil)
	suite.mockinfra.On("ComparePassword", []byte(hashedPassword), []byte(user.Password)).Return(nil)
	suite.mockinfra.On("GenerateJwtToken", user).Return("jwt_token", nil)

	jwtToken, err := suite.usecase.Login(user.Username, user.Password)
	suite.NoError(err)
	suite.Equal("jwt_token", jwtToken)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestLoginMissingFields() {
	jwtToken, err := suite.usecase.Login("", "")
	suite.Error(err)
	suite.Equal("", jwtToken)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestLoginInvalidUsername() {
	user := &domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("FetchByUsername", user.Username).Return(domain.User{}, errors.New("not found"))
	jwtToken, err := suite.usecase.Login(user.Username, user.Password)
	suite.Error(err)
	suite.Equal("", jwtToken)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestLoginInvalidCredentials() {
	user := domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("FetchByUsername", user.Username).Return(user, nil)
	suite.mockinfra.On("ComparePassword", []byte(user.Password), []byte("wrong")).Return(errors.New("invalid password"))
	jwtToken, err := suite.usecase.Login(user.Username, "wrong")
	suite.Error(err)
	suite.Equal("", jwtToken)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestPromoteUser() {
	user := &domain.User{
		ID: primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("Fetch", user.ID.Hex()).Return(*user, nil)
	suite.mockRepo.On("Promote", user).Return(*user, nil)
	_, err := suite.usecase.Promote(user.ID.Hex())
	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestPromoteNonExistentUser() {
	userID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("Fetch", userID).Return(domain.User{}, errors.New("not found"))

	promotedUser, err := suite.usecase.Promote(userID)
	suite.Error(err)
	suite.Equal(domain.User{}, promotedUser)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestUserFetch() {
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}

	suite.mockRepo.On("Fetch", user.ID.Hex()).Return(*user, nil)

	fetchedUser, err := suite.usecase.Fetch(user.ID.Hex())
	suite.NoError(err)
	suite.Equal(*user, fetchedUser)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestUserUpdate() {
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
	}
	suite.mockRepo.On("Update", user.ID.Hex(), *user).Return(*user, nil)

	updatedUser, err := suite.usecase.Update(user.ID.Hex(), *user)
	suite.NoError(err)
	suite.Equal(*user, updatedUser)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestChangePassword() {
	userID := primitive.NewObjectID().Hex()

	prevPassword := "oldpassword"
	newPassword := "newpassword"

	suite.mockRepo.On("Fetch", userID).Return(domain.User{ID: primitive.NewObjectID(), Password: "hashedoldpassword"}, nil)
	suite.mockinfra.On("ComparePassword", []byte("hashedoldpassword"), []byte(prevPassword)).Return(nil)
	suite.mockRepo.On("ChangePassword", userID, prevPassword, newPassword).Return(nil)

	err := suite.usecase.ChangePassword(userID, prevPassword, newPassword)
	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestChangePasswordIncorrectOldPassword() {
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "oldpassword",
		Email:    "testuser@example.com",
	}
	prevPassword := "wrongpassword"
	newPassword := "newpassword"

	suite.mockRepo.On("Fetch", user.ID.Hex()).Return(*user, nil)
	suite.mockinfra.On("ComparePassword", []byte(user.Password), []byte(prevPassword)).Return(errors.New("incorrect password"))

	err := suite.usecase.ChangePassword(user.ID.Hex(), prevPassword, newPassword)
	suite.Error(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestRemoveNonexistantUser() {
	userID := primitive.NewObjectID().Hex()

	suite.mockRepo.On("Fetch", userID).Return(domain.User{}, errors.New("not found"))
	err := suite.usecase.Remove(userID)
	suite.Error(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestRemoveUser() {
	userID := primitive.NewObjectID()

	suite.mockRepo.On("Fetch", userID.Hex()).Return(domain.User{ID: userID}, nil)
	suite.mockRepo.On("Remove", userID.Hex()).Return(nil)

	err := suite.usecase.Remove(userID.Hex())
	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func (suite *UserTestSuite) TestRemoveAdminUser() {
	userID := primitive.NewObjectID()

	suite.mockRepo.On("Fetch", userID.Hex()).Return(domain.User{ID: userID, Role: "admin"}, nil)

	err := suite.usecase.Remove(userID.Hex())
	suite.Error(err)

	suite.mockRepo.AssertExpectations(suite.T())
	suite.mockinfra.AssertExpectations(suite.T())
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}