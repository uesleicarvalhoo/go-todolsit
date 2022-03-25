package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
)

// GetMe godoc
// @Sumary       Endpoint to Get User info
// @Description  Return current user info
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.User
// @Failure      400  {object}  handler.Message
// @Router       /api/v1/user/me  [get]
func (h *Handler) GetMe(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.GetMe")

	defer span.End()

	trace.AddSpanTags(span, map[string]string{"app.user-id": c.GetHeader("X-User-Id")})

	userId, err := uuid.Parse(c.GetHeader("X-User-Id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Message{Message: fmt.Sprintf("Invalid userId: %s", userId)})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Invalid User")
		return
	}

	user, err := h.UserSvc.Get(ctx, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Message{Message: err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "User not Found")
		return
	}

	c.JSON(http.StatusOK, user)
}
