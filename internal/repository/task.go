package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(dbInstance *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: dbInstance,
	}
}

func (tr *TaskRepository) Get(ctx context.Context, id uuid.UUID) (task *entity.Task, err error) {
	ctx, span := trace.NewSpan(ctx, "Task.Get")
	defer span.End()

	tx := tr.DB.First(&task, "id = ?", id)

	return task, tx.Error
}

func (tr *TaskRepository) GetAll(ctx context.Context, ownerId uuid.UUID) (tasks []*entity.Task, err error) {
	ctx, span := trace.NewSpan(ctx, "Task.GetAll")
	defer span.End()

	tx := tr.DB.Find(&tasks, "owner_id = ?", ownerId)

	return tasks, tx.Error
}

func (tr *TaskRepository) Create(ctx context.Context, task *entity.Task) error {
	ctx, span := trace.NewSpan(ctx, "Task.Create")
	defer span.End()

	tx := tr.DB.Create(&task)
	return tx.Error
}

func (tr *TaskRepository) Update(ctx context.Context, task *entity.Task) error {
	ctx, span := trace.NewSpan(ctx, "Task.Update")
	defer span.End()

	tx := tr.DB.Save(&task)
	return tx.Error
}

func (tr *TaskRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	ctx, span := trace.NewSpan(ctx, "Task.Delete")
	defer span.End()

	tx := tr.DB.Delete(&entity.Task{}, "id = ?", id)
	return tx.Error
}
