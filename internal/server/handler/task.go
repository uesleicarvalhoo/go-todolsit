package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
)

// CreatTask godoc
// @Sumary       Endpoint to Register new Task
// @Description  Register new task and return data
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      201  {object}  entity.Task
// @Failure      400  {object}  handler.Message
// @Router       /api/v1/task/create [post]
func (h *Handler) CreateTask(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.CreateTask")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id")})

	var payload entity.RegisterTask
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	task, err := h.TaskSvc.NewTask(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad Request, internal server error")
	}

	c.JSON(http.StatusCreated, task)
}

// GetTaskData godoc
// @Sumary   Endpoint to list task data
// @Param    uuid  path  string  true  "The ID of Task"
// @Tags     Task
// @Accept   json
// @Produce  json
// @Success  200  {object}  entity.Task
// @Failure  400  {object}  handler.Message
// @Router   /api/v1/task/{taskId} [get]
func (h *Handler) GetTaskData(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.GetTaskData")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id"), "app.task-id": c.Param("taskId")})

	taskId, err := uuid.Parse(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("Invalid taskId: %s", taskId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	task, err := h.TaskSvc.Get(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad Request, internal server error")
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetTaskData godoc
// @Sumary       Endpoint to list all tasks
// @Description  Return all tasks of current user
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200  {array}   entity.Task
// @Failure      400  {object}  handler.Message
// @Router       /api/v1/task/ [get]
func (h *Handler) ListTasks(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.ListTasks")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id")})

	userId, err := uuid.Parse(c.Request.Header.Get("X-User-Id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("Invalid UserId: %s", userId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	tasks, err := h.TaskSvc.ListTasks(ctx, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad request, internal server error")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// FinishTask godoc
// @Sumary       Endpoint to finish task
// @Description  Set status of task with finished
// @Param        uuid  path  string  true  "The ID of Task"
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Task
// @Failure      400  {object}  handler.Message
// @Failure      500  {object}  string
// @Router       /api/v1/task/{taskId}/finish [post]
func (h *Handler) FinishTask(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.FinishTask")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id"), "app.task-id": c.Param("taskId")})

	taskId, err := uuid.Parse(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("Invalid taskId: %s", taskId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	task, err := h.TaskSvc.Get(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad request, internal server error")
		return
	}

	err = h.TaskSvc.FinishTask(ctx, task)
	if err != nil {
		trace.AddSpanError(span, err)

		var status int
		if err == entity.ErrTaskAlreadyIsClosed {
			status = http.StatusBadRequest
			trace.FailSpan(span, "Task alread is closed")
		} else {
			status = http.StatusInternalServerError
			trace.FailSpan(span, "Internal server error")
		}
		c.AbortWithStatusJSON(status, Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// HealthCheck godoc
// @Sumary       Endpoint Remove a Task
// @Description  Exclude task
// @Param        uuid  path  string  true  "The ID of Task"
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  handler.Message
// @Failure      500  {object}  string
// @Router       /api/v1/task/{taskId} [delete]
func (h *Handler) ExcludeTask(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.ExcludeTask")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id"), "app.task-id": c.Param("taskId")})

	taskId, err := uuid.Parse(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("Invalid taskId: %s", taskId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	err = h.TaskSvc.RemoveTask(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad request, internal server error")
		return
	}
	c.Status(http.StatusOK)
}

// UpdateTask godoc
// @Sumary       Endpoint Update a Task
// @Param        uuid  path  string  true  "The ID of Task"
// @Description  Update task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  handler.Message
// @Failure      500  {object}  handler.Message
// @Router       /api/v1/task/{taskId} [post]
func (h *Handler) UpdateTask(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.UpdateTask")
	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id"), "app.task-id": c.Param("taskId")})

	taskId, err := uuid.Parse(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("Invalid taskId: %s", taskId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, fmt.Sprintf("Task %s not found", taskId))
		return
	}

	task, err := h.TaskSvc.Get(ctx, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Task not found")
		return
	}

	var payload entity.RegisterTask
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	err = h.TaskSvc.UpdateTask(ctx, task, payload.Title, payload.Description)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Bad request, internal server error")
		return
	}
	c.Status(http.StatusOK)
}
