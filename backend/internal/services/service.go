package services

import (
	"backend/internal/models"
)

type AuthServiceInterface interface {
	CreateUser(body *models.AuthRequest) (*models.Response, *models.ErrorResponse)
	DeleteUser(userID int) (*models.Response, *models.ErrorResponse)
	LogoutUser(userID int) (*models.Response, *models.ErrorResponse)
	GetUserByID(userID int) (*models.Response, *models.ErrorResponse)
	LoginUser(body *models.AuthRequest) (*models.TokenResponse, *models.ErrorResponse)
	GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, *models.ErrorResponse)
}

type TasksServiceInterface interface {
	CreateTask(taskRequest *models.TaskRequest, userID int) (*models.Response, *models.ErrorResponse)
	GetTaskByID(taskID int, userID int) (*models.Response, *models.ErrorResponse)
	GetTasks(userID int) (*models.Response, *models.ErrorResponse)
	UpdateTask(taskRequest *models.TaskRequest, userID int, taskID int) (*models.Response, *models.ErrorResponse)
	DeleteTask(taskID int, userID int) (*models.Response, *models.ErrorResponse)
}
