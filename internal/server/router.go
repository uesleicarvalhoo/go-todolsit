package server

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/uesleicarvalhoo/go-todolist/docs"
	"github.com/uesleicarvalhoo/go-todolist/internal/server/handler"
	"github.com/uesleicarvalhoo/go-todolist/internal/server/middleware"
)

func RouterInit(engine *gin.Engine, handlers *handler.Handler) {
	api := engine.Group("/api")

	api.GET("/health-check", handlers.HealthCheck)
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := api.Group("v1/auth")
	authGroup.POST("/signup", handlers.SiginUp)
	authGroup.POST("/login", handlers.Login)

	userGroup := api.Group("v1/user")
	userGroup.GET("/me", middleware.AuthenticationMiddleware(handlers.GetMe))

	taskGroup := api.Group("v1/task")
	taskGroup.POST("/", middleware.AuthenticationMiddleware(handlers.CreateTask))
	taskGroup.GET("/", middleware.AuthenticationMiddleware(handlers.ListTasks))
	taskGroup.GET("/:taskId", middleware.AuthenticationMiddleware(handlers.GetTaskData))
	taskGroup.POST("/:taskId", middleware.AuthenticationMiddleware(handlers.UpdateTask))
	taskGroup.POST("/:taskId/finish", middleware.AuthenticationMiddleware(handlers.FinishTask))
	taskGroup.DELETE("/:taskId", middleware.AuthenticationMiddleware(handlers.ExcludeTask))
}
