package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	FinishedAt  time.Time
	OwnerId     uuid.UUID
}

func NewTask(ownerId uuid.UUID, title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrMissingTitle
	}

	if description == "" {
		return nil, ErrMissingDescription
	}
	return &Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		OwnerId:     ownerId,
	}, nil
}

func (t *Task) Finish() error {
	if t.Done {
		return ErrTaskAlreadyIsClosed
	}
	t.Done = true
	t.FinishedAt = time.Now()
	return nil
}
