package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uesleicarvalhoo/go-todolist/internal/entity"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
)

// SiginUp godoc
// @Sumary   Endpoint to Create new User
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Success  201  {object}  entity.User
// @Failure  400  {object}  handler.Message
// @Router   /api/v1/auth/signup [post]
func (h *Handler) SiginUp(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.SiginUp")
	defer span.End()

	var payload entity.SiginUp
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	user, err := h.UserSvc.SiginUp(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Internal server error")
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Sumary       Endpoint to get Auth token
// @Description  Return auth token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.LoginResponse
// @Failure      400  {object}  handler.Message
// @Router       /api/v1/auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	ctx, span := trace.NewSpan(c.Request.Context(), "Handler.Login")
	defer span.End()

	var payload entity.Login
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Unprocessable Entity")
		return
	}

	res, err := h.UserSvc.Login(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "Login Unauthorized")
	}

	c.JSON(http.StatusOK, res)
}
