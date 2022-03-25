package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Sumary   Return status of Service
// @Tags     General
// @Accept   json
// @Produce  json
// @Success  200  {object}  handler.Message
// @Router   /api/health-check [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Message: "ok",
	})
}
