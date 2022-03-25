package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(dbInstance *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: dbInstance,
	}
}

func (ur *UserRepository) Get(ctx context.Context, id uuid.UUID) (user *entity.User, err error) {
	ctx, span := trace.NewSpan(ctx, "User.Get")
	defer span.End()

	tx := ur.DB.First(&user, "id = ?", id)
	return user, tx.Error
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	ctx, span := trace.NewSpan(ctx, "User.GetByEmail")
	defer span.End()

	tx := ur.DB.First(&user, "email = ?", email)
	return user, tx.Error
}

func (ur *UserRepository) Create(ctx context.Context, user *entity.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Create")
	defer span.End()

	tx := ur.DB.Create(&user)
	return tx.Error
}

func (ur *UserRepository) Update(ctx context.Context, user *entity.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Update")
	defer span.End()

	tx := ur.DB.Save(user)
	return tx.Error
}

func (ur *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	ctx, span := trace.NewSpan(ctx, "User.Delete")
	defer span.End()

	tx := ur.DB.Delete(&entity.User{}, "id = ?")
	return tx.Error
}
