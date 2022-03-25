package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uesleicarvalhoo/go-todolist/internal/config"
	"github.com/uesleicarvalhoo/go-todolist/internal/repository"
	"github.com/uesleicarvalhoo/go-todolist/internal/server"
	"github.com/uesleicarvalhoo/go-todolist/internal/server/handler"
	"github.com/uesleicarvalhoo/go-todolist/internal/server/middleware"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/task"
	"github.com/uesleicarvalhoo/go-todolist/internal/services/user"
	"github.com/uesleicarvalhoo/go-todolist/pkg/database"
	"github.com/uesleicarvalhoo/go-todolist/pkg/trace"
	"github.com/uesleicarvalhoo/go-todolist/pkg/utils"
)

func main() {
	env := config.GetEnv()
	ctx := context.Background()

	// Database
	dbInstance, err := database.NewPostgreSQLConnection(env.DBHost, env.DBPort, env.DBName, env.DBUser, env.DBPassword)
	if err != nil {
		log.Fatal(utils.ErrDatabaseConnection)
		panic(err)
	}

	err = repository.DBMigrate(dbInstance, env.DBName)
	if err != nil {
		fmt.Println(utils.ErrRunMigrations)
		panic(err)
	}

	// Gin
	engine := gin.New()

	// Middlewares
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	customCors := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	})

	provider, err := trace.NewProvider(trace.ProviderConfig{
		JaegerEndpoint: fmt.Sprintf("http://%s/api/traces", env.TraceHost),
		ServiceName:    config.ServiceName,
		ServiceVersion: config.Version,
		Environment:    env.Env,
		Disabled:       false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	defer provider.Close(ctx)

	engine.Use(
		customCors,
		gin.Recovery(),
		middleware.LogMiddleware(log),
		gzip.Gzip(gzip.DefaultCompression),
	)

	// Services
	userRepository := repository.NewUserRepository(dbInstance)
	taskRepository := repository.NewTaskRepository(dbInstance)

	userService := user.NewService(userRepository)
	taskService := task.NewService(userService, taskRepository)

	handler := handler.NewHandler(userService, taskService)
	server.RouterInit(engine, handler)

	gin.SetMode(gin.ReleaseMode)
	if env.Debug {
		gin.SetMode(gin.DebugMode)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", env.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("API server forced to shutdown:", err)
	}
}
