package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (user *entity.User, err error)
	GetByEmail(ctx context.Context, email string) (user *entity.User, err error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
