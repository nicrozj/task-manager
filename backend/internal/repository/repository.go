package repository

import (
	"backend/internal/models"
)

type AuthRepositoryInterface interface {
	CreateUser(user *models.User) (int, error)
	LogoutUser(userID int) (int, error)
	DeleteUser(userID int) (int, error)
	GetUserByID(userID int) (*models.User, int, error)
	LoginUser(user *models.User) (*models.TokenResponse, int, error)
	GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, int, error)
	getAuthTokens(userID int) (*models.TokenResponse, int, error)
}

type TasksRepositoryInterface interface {
	CreateTask(task *models.Task) (int, int, error)
	GetTaskByID(taskID int, userID int) (*models.Task, int, error)
	GetTasks(userID int) ([]*models.Task, int, error)
	UpdateTask(task *models.Task) (int, error)
	DeleteTask(taskID int, userID int) (int, error)
}
