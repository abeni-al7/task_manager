package tests

import (
	"testing"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Usecases"
	"github.com/abeni-al7/task_manager/Usecases/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockTaskRepo
	usecase  usecases.TaskUsecase
}

func (suite *TaskTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockTaskRepo)
	suite.usecase = *usecases.NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskTestSuite) TestTaskCreate() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	suite.mockRepo.On("Create", task).Return(*task, nil)

	createdTask, err := suite.usecase.Create(task)
	suite.NoError(err)
	suite.Equal(task.Title, createdTask.Title)
	suite.Equal(task.Description, createdTask.Description)
	suite.Equal(task.Status, createdTask.Status)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskCreateMissingFields() {
	task := &domain.Task{
		Title:       "",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	createdTask, err := suite.usecase.Create(task)
	suite.Error(err)
	suite.Equal(domain.Task{}, createdTask)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskCreateInvalidStatus() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "invalid-status",
	}

	createdTask, err := suite.usecase.Create(task)
	suite.Error(err)
	suite.Equal(domain.Task{}, createdTask)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskFetchAll() {
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now().Add(48 * time.Hour),
			Status:      "pending",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now().Add(72 * time.Hour),
			Status:      "in-progress",
		},
	}

	suite.mockRepo.On("FetchAll").Return(tasks, nil)

	fetchedTasks, err := suite.usecase.FetchAll()
	suite.NoError(err)
	suite.Equal(len(tasks), len(fetchedTasks))

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskFetch() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	suite.mockRepo.On("Fetch", task.ID.Hex()).Return(task, nil)

	fetchedTask, err := suite.usecase.Fetch(task.ID.Hex())
	suite.NoError(err)
	suite.Equal(task.ID, fetchedTask.ID)
	suite.Equal(task.Title, fetchedTask.Title)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskUpdate() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "This is an updated task",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      "in-progress",
	}

	suite.mockRepo.On("Update", task.ID.Hex(), task).Return(task, nil)

	updatedTask, err := suite.usecase.Update(task.ID.Hex(), task)
	suite.NoError(err)
	suite.Equal(task.ID, updatedTask.ID)
	suite.Equal(task.Title, updatedTask.Title)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskUpdateInvalidStatus() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "This is an updated task",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      "invalid-status",
	}

	updatedTask, err := suite.usecase.Update(task.ID.Hex(), task)
	suite.Error(err)
	suite.Equal(domain.Task{}, updatedTask)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskTestSuite) TestTaskRemove() {
	taskID := primitive.NewObjectID()

	suite.mockRepo.On("Remove", taskID.Hex()).Return(nil)

	err := suite.usecase.Remove(taskID.Hex())
	suite.NoError(err)

	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}
