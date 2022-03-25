package entity_test

import (
	"fmt"
	"testing"

	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
)

func TestNewUser(t *testing.T) {
	type testCase struct {
		testName     string
		userName     string
		userEmail    string
		userPassword string
		expectedErr  error
	}

	testCases := []testCase{
		// {
		// 	testName:     "error is ErrMissingTitle when title is null",
		// 	userName:     "",
		// 	userEmail:    "email@teste.com",
		// 	userPassword: "MySecretPassword",
		// 	expectedErr:  entity.ErrMissingName,
		// },
		// {
		// 	testName:     "error is ErrMissingEmail when email is null",
		// 	userName:     "Ueslei Carvalho",
		// 	userEmail:    "",
		// 	userPassword: "MySecretPassword",
		// 	expectedErr:  entity.ErrMissingEmail,
		// },
		// {
		// 	testName:     "error is ErrPasswordMissing when password is null",
		// 	userName:     "Ueslei Carvalho",
		// 	userEmail:    "email@teste.com",
		// 	userPassword: "",
		// 	expectedErr:  entity.ErrMissingPassword,
		// },
		{
			testName:    "error nil when name and email is valid",
			userName:    "Ueslei Carvalho",
			userEmail:   "email@teste.com",
			userPassword: "MySecretPassword",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewUser(tc.userName, tc.userEmail, tc.userPassword)
			fmt.Println(tc.userName, tc.userEmail, tc.userPassword)
			if err != tc.expectedErr {
				t.Errorf("Expected %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
