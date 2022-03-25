package entity_test

import (
	"testing"

	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
)

func TestNewTask(t *testing.T) {
	type testCase struct {
		testName        string
		taskTitle       string
		taskDescription string
		expectedErr     error
	}
	user, err := entity.NewUser("Ueslei Carvalho", "teste@email.com", "myFakePassword")

	if err != nil {
		t.Error(err)
	}

	testCases := []testCase{
		{
			testName:        "error is ErrInvalidTitle when title is null",
			taskTitle:       "",
			taskDescription: "Any description",
			expectedErr:     entity.ErrMissingTitle,
		},
		{
			testName:        "error is ErrInvalidDescription when description is null",
			taskTitle:       "Any title",
			taskDescription: "",
			expectedErr:     entity.ErrMissingDescription,
		},
		{
			testName:        "error nil when title, description and password is valid",
			taskTitle:       "Any title",
			taskDescription: "Any description",
			expectedErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewTask(user.Id, tc.taskTitle, tc.taskDescription)
			if err != tc.expectedErr {
				t.Errorf("Expected %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
