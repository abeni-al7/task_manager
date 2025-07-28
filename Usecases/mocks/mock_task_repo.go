package mocks

import (
	domain "github.com/abeni-al7/task_manager/Domain"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) Create(task *domain.Task) (domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepo) FetchAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Fetch(idStr string) (domain.Task, error) {
	args := m.Called(idStr)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Update(idStr string, task domain.Task) (domain.Task, error) {
	args := m.Called(idStr, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepo) Remove(idStr string) error {
	args := m.Called(idStr)
	return args.Error(0)
}