package entity

import "errors"

var (
	// User
	ErrMissingName  = errors.New("The name must be informed")
	ErrMissingEmail = errors.New("The email must be informed")
	ErrMissingPassword = errors.New("The password must be informed")

	// Task
	ErrMissingTitle        = errors.New("The Title must be informed")
	ErrMissingDescription  = errors.New("The Description must be informed")
	ErrTaskAlreadyIsClosed = errors.New("The Task already is closed")
)
