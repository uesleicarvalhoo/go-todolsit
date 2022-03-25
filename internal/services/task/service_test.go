package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/internal/repository"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/user"
	"github.com/uesleicarvalhoo/go-todolist/pkg/database"
	"github.com/uesleicarvalhoo/go-todolist/pkg/utils"
)

var taskService *Service
var existingUser *entity.User

func init() {
	db, err := database.NewSQLiteConnection()
	if err != nil {
		fmt.Println(utils.ErrDatabaseConnection)
		panic(err)
	}

	err = repository.AutoMigrate(db)
	if err != nil {
		fmt.Println(utils.ErrRunMigrations)
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userService := user.NewService(userRepository)
	taskService = NewService(userService, taskRepository)

	u, err := userService.SiginUp(context.TODO(), newFakeSignUp())
	if err != nil {
		panic(err)
	}
	existingUser = u
}

func newFakeSignUp() entity.SiginUp {
	return entity.SiginUp{Name: "Fake Name", Email: "fake@email.com", Password: "MySecretPasswd"}
}

func newFakeRegisterTask(ownerId uuid.UUID) entity.RegisterTask {
	return entity.RegisterTask{
		OwnerId:     ownerId,
		Title:       "Title of task",
		Description: "Description of task",
	}
}

func TestNewTaskSaveTaskInRepository(t *testing.T) {
	// Arrange
	data := newFakeRegisterTask(existingUser.Id)

	// Action
	task, err := taskService.NewTask(context.TODO(), data)
	assert.Nil(t, err)

	// Assert
	searchedTask, err := taskService.Repository.Get(context.TODO(), task.Id)
	assert.Nil(t, err)
	assert.Equal(t, searchedTask.Id, task.Id)
}

func TestGetErrorReturnNilWhenTaskExist(t *testing.T) {
	// Arrange
	data := newFakeRegisterTask(existingUser.Id)

	// Action
	existingTask, err := taskService.NewTask(context.TODO(), data)
	assert.Nil(t, err)

	// Assert
	_, err = taskService.Get(context.TODO(), existingTask.Id)
	assert.Nil(t, err)
}

func TestGetErrorReturnTaskNotFoundWhenTaskNotExist(t *testing.T) {
	// Arrange
	fakeTaskId := uuid.New()

	// Action
	_, err := taskService.Get(context.TODO(), fakeTaskId)

	// Assert
	assert.NotNil(t, err)
}

func TestGetErrorReturnErrorWhenTaskNotInRepository(t *testing.T) {
	// Arrange
	fakeTaskId := uuid.New()

	// Action
	_, err := taskService.Get(context.TODO(), fakeTaskId)

	// Assert
	assert.NotNil(t, err)
}

func TestRemoveTaskErrorRemoveTaskFromRepository(t *testing.T) {
	// Arrange
	existingTask, err := taskService.NewTask(context.TODO(), newFakeRegisterTask(existingUser.Id))
	assert.Nil(t, err)

	// Action
	err = taskService.RemoveTask(context.TODO(), existingTask.Id)
	assert.Nil(t, err)

	// Assert
	_, err = taskService.Get(context.TODO(), existingTask.Id)
	assert.NotNil(t, err)
}

func TestUpdateTaskUpdateMemoryValues(t *testing.T) {
	// Arrange
	existingTask, err := taskService.NewTask(context.TODO(), newFakeRegisterTask(existingUser.Id))
	assert.Nil(t, err)

	newTitle := "New task title"
	newDescription := "New task description"

	// Action
	err = taskService.UpdateTask(context.TODO(), existingTask, newTitle, newDescription)
	assert.Nil(t, err)

	// Assert
	assert.Equal(t, existingTask.Title, newTitle)
	assert.Equal(t, existingTask.Description, newDescription)
}

func TestUpdateTaskUpdateSiginUpInRepository(t *testing.T) {
	// Prepare
	newTitle := "New task title"
	newDescription := "New task description"

	// Arrange
	existingTask, err := taskService.NewTask(context.TODO(), newFakeRegisterTask(existingUser.Id))
	assert.Nil(t, err)

	// Action
	err = taskService.UpdateTask(context.TODO(), existingTask, newTitle, newDescription)
	assert.Nil(t, err)

	// Assert
	updatedTask, err := taskService.Repository.Get(context.TODO(), existingTask.Id)
	assert.Nil(t, err)
	assert.Equal(t, updatedTask.Title, newTitle)
	assert.Equal(t, updatedTask.Description, newDescription)
}

func TestFinishTaskUpdateSiginUpInRepository(t *testing.T) {
	// Arrange
	existingTask, err := taskService.NewTask(context.TODO(), newFakeRegisterTask(existingUser.Id))
	assert.Nil(t, err)

	// Action
	err = taskService.FinishTask(context.TODO(), existingTask)
	assert.Nil(t, err)

	// Assert
	updatedTask, err := taskService.Repository.Get(context.TODO(), existingTask.Id)
	assert.Nil(t, err)
	assert.True(t, updatedTask.Done)
	assert.False(t, updatedTask.FinishedAt.IsZero())
}
