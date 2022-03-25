package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (task *entity.Task, err error)
	GetAll(ctx context.Context, ownerId uuid.UUID) (task []*entity.Task, err error)
	Create(ctx context.Context, task *entity.Task) error
	Update(ctx context.Context, task *entity.Task) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
