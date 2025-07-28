package mocks

import (
	domain "github.com/abeni-al7/task_manager/Domain"
	"github.com/stretchr/testify/mock"
)

type MockInfrastructure struct {
	mock.Mock
}

func (m *MockInfrastructure) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockInfrastructure) ComparePassword(correctPassword []byte, inputPassword []byte) error {
	args := m.Called(correctPassword, inputPassword)
	return args.Error(0)
}

func (m *MockInfrastructure) GenerateJwtToken(user *domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}