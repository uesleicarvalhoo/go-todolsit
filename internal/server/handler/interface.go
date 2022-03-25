package handler

import (
	"github.com/uesleicarvalhoo/go-todolist/internal/services/task"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/user"
)

type Handler struct {
	UserSvc *user.Service
	TaskSvc *task.Service
}

type Message struct {
	Message string
}

func NewHandler(userSvc *user.Service, taskSvc *task.Service) *Handler {
	return &Handler{
		UserSvc: userSvc,
		TaskSvc: taskSvc,
	}
}
