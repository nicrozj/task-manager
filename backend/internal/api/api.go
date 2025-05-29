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

	authGroup := s.Engine.Group("/auth")
	{
		authGroup.GET("/greet", authHandlers.Greet)
		authGroup.POST("/login", authHandlers.LoginUser)
		authGroup.POST("/registration", authHandlers.CreateUser)
	}

	protectedGroup := s.Engine.Group("/")
	protectedGroup.Use(jwtMiddleware())
	{
		protectedGroup.POST("/auth/logout", authHandlers.LogoutUser)
		protectedGroup.POST("/auth/refresh", authHandlers.RefreshToken)
		protectedGroup.DELETE("/users", authHandlers.DeleteUser)
		protectedGroup.GET("/users/me", authHandlers.GetUserByID)
	}
}
