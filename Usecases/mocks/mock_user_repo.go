package mocks

import (
	domain "github.com/abeni-al7/task_manager/Domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FetchByUsername(username string) (domain.User, error) {
	args := m.Called(username)
	result := args.Get(0)
	if result == nil {
		return domain.User{}, args.Error(1)
	}
	return result.(domain.User), args.Error(1)
}

func(m *MockUserRepo) CountUsers() (int, error) {
	args := m.Called()
	return args.Get(0).(int), args.Error(1)
}

func(m *MockUserRepo) Register(user *domain.User) (domain.User, error) {
	args := m.Called(user)
	result := args.Get(0)
	if result == nil {
		return domain.User{}, args.Error(1)
	}
	return result.(domain.User), args.Error(1)
}

func(m *MockUserRepo) Promote(user *domain.User) (domain.User, error) {
	args := m.Called(user)
	result := args.Get(0)
	if result == nil {
		return domain.User{}, args.Error(1)
	}
	return result.(domain.User), args.Error(1)
}

func(m *MockUserRepo) FetchAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func(m *MockUserRepo) Fetch(idStr string) (domain.User, error) {
	args := m.Called(idStr)
	return args.Get(0).(domain.User), args.Error(1)
}

func(m *MockUserRepo) Update(idStr string, updatedUser domain.User) (domain.User, error) {
	args := m.Called(idStr, updatedUser)
	return args.Get(0).(domain.User), args.Error(1)
}

func(m *MockUserRepo) ChangePassword(idStr string, prevPassword string, newPassword string) error {
	args := m.Called(idStr, prevPassword, newPassword)
	return args.Error(0)
}

func(m *MockUserRepo) Remove(idStr string) error {
	args := m.Called(idStr)
	return args.Error(0)
}