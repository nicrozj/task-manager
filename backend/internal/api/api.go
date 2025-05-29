package api

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	s := &Server{
		Engine: gin.Default(),
	}
	s.mountMiddlewares()
	s.mountHandlers()
	return s
}

func (s *Server) mountMiddlewares() {
	s.Engine.Use(func(c *gin.Context) {
		c.Set("timeout", 1*time.Minute)
		c.Next()
	})

	s.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Envs.WEB_URL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func (s *Server) mountHandlers() {
	authHandlers := handlers.NewAuthHandlers()
	tasksHandlers := handlers.NewTasksHandlers()

	authGroup := s.Engine.Group("/auth")
	{
		authGroup.GET("/greet", authHandlers.Greet)
		authGroup.POST("/login", authHandlers.LoginUser)
		authGroup.POST("/registration", authHandlers.CreateUser)
	}

	protectedGroup := s.Engine.Group("/auth")
	protectedGroup.Use(jwtMiddleware())
	{
		protectedGroup.POST("/logout", authHandlers.LogoutUser)
		protectedGroup.POST("/refresh", authHandlers.RefreshToken)
		protectedGroup.DELETE("/users", authHandlers.DeleteUser)
		protectedGroup.GET("/me", authHandlers.GetUserByID)
	}

	tasksGroup := s.Engine.Group("/tasks")
	tasksGroup.Use(jwtMiddleware())
	{
		tasksGroup.POST("", tasksHandlers.CreateTask)
		tasksGroup.GET("/:id", tasksHandlers.GetTaskByID)
		tasksGroup.GET("", tasksHandlers.GetTasks)
		tasksGroup.PUT("/:id", tasksHandlers.UpdateTask)
		tasksGroup.DELETE("/:id", tasksHandlers.DeleteTask)
	}
}
