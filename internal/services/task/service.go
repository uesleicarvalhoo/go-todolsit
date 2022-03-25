package task

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/user"
)

var (
	ErrMissingValues = errors.New("The title or description must be informed!")
	ErrTaskIsClosed  = errors.New("The current task is already finished")
)

type Service struct {
	userService *user.Service
	Repository  Repository
}

func NewService(userService *user.Service, repository Repository) *Service {
	return &Service{
		userService: userService,
		Repository:  repository,
	}
}

func (s *Service) NewTask(ctx context.Context, payload entity.RegisterTask) (*entity.Task, error) {
	_, err := s.userService.Get(ctx, payload.OwnerId)
	if err != nil {
		return nil, err
	}

	task, err := entity.NewTask(payload.OwnerId, payload.Title, payload.Description)
	if err != nil {
		return nil, err
	}
	s.Repository.Create(ctx, task)
	return task, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	return s.Repository.Get(ctx, id)
}

func (s *Service) ListTasks(ctx context.Context, ownerId uuid.UUID) ([]*entity.Task, error) {
	return s.Repository.GetAll(ctx, ownerId)
}

func (s *Service) FinishTask(ctx context.Context, task *entity.Task) error {
	err := task.Finish()
	if err != nil {
		return err
	}

	return s.Repository.Update(ctx, task)
}

func (s *Service) UpdateTask(ctx context.Context, task *entity.Task, newTitle, newDescription string) error {
	if newTitle == "" && newDescription == "" {
		return ErrMissingValues
	}

	task.Description = newDescription
	task.Title = newTitle

	err := s.Repository.Update(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveTask(ctx context.Context, id uuid.UUID) error {
	task, err := s.Repository.Get(ctx, id)
	if err != nil {
		return err
	}
	return s.Repository.DeleteById(ctx, task.Id)
}
