package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uesleicarvalhoo/go-todolist/internal/config"
)

func LogMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/health-check") {
			return
		}
		c.Next()
		statusCode := c.Writer.Status()
		entry := logger.WithFields(logrus.Fields{
			"log_version": "1.0.0",
			"date_time":   time.Now(),
			"product": map[string]interface{}{
				"name":        config.ServiceName,
				"application": config.ServiceName,
				"version":     config.Version,
				"http": map[string]string{
					"method": c.Request.Method,
					"path":   c.Request.URL.Path,
				},
			},
			"origin": map[string]interface{}{
				"application": config.ServiceName,
				"ip":          c.ClientIP(),
				"headers": map[string]string{
					"user_agent": c.Request.UserAgent(),
					"origin":     c.GetHeader("Origin"),
					"refer":      c.Request.Referer(),
				},
			},
			"context": map[string]interface{}{
				"service":     config.ServiceName,
				"status_code": statusCode,
			},
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode > 499 {
				entry.Error()
			} else if statusCode > 399 {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}
