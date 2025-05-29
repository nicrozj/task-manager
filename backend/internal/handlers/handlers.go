package handlers

import "github.com/gin-gonic/gin"

type AuthHandlersInterface interface {
	Greet(c *gin.Context)
	CreateUser(c *gin.Context)
	LogoutUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	LoginUser(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type TasksHandlersInterface interface {
	CreateTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTasks(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}
