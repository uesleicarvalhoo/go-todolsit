package user_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/internal/repository"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/user"
	"github.com/uesleicarvalhoo/go-todolist/pkg/database"
	"github.com/uesleicarvalhoo/go-todolist/pkg/utils"
)

var userService *user.Service

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
	userService = user.NewService(userRepository)
}

func newFakeSignUp() entity.SiginUp {
	return entity.SiginUp{
		Name:  "Fake name",
		Email: "fake@email.com",
		Password: "MySuperSecretPassword",
	}
}

func TestRegisterSaveUserInRepository(t *testing.T) {
	// Arrange
	user, err := userService.SiginUp(context.TODO(), newFakeSignUp())
	assert.Nil(t, err)

	// Assert
	searchedUser, err := userService.Get(context.TODO(), user.Id)
	assert.Nil(t, err)

	assert.Equal(t, searchedUser.Id, user.Id)
}
