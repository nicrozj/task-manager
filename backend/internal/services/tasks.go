package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type TasksService struct {
	repo repository.TasksRepositoryInterface
}

func NewTasksService() TasksServiceInterface {
	return &TasksService{
		repo: repository.NewTasksRepo(),
	}
}

func (s TasksService) CreateTask(taskRequest *models.TaskRequest, userID int) (*models.Response, *models.ErrorResponse) {
	task := &models.Task{Title: taskRequest.Title, Description: taskRequest.Description, Status: taskRequest.Status, UserId: userID}

	id, status, err := s.repo.CreateTask(task)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status, Data: id}, nil
}
func (s TasksService) GetTaskByID(taskID int, userID int) (*models.Response, *models.ErrorResponse) {
	task, status, err := s.repo.GetTaskByID(taskID, userID)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status, Data: task}, nil
}

func (s TasksService) GetTasks(userID int) (*models.Response, *models.ErrorResponse) {
	tasks, status, err := s.repo.GetTasks(userID)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status, Data: tasks}, nil
}

func (s TasksService) UpdateTask(taskRequest *models.TaskRequest, userID int, taskID int) (*models.Response, *models.ErrorResponse) {
	status, err := s.repo.UpdateTask(&models.Task{Id: taskID, UserId: userID, Title: taskRequest.Title, Description: taskRequest.Description, Status: taskRequest.Status})
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status}, nil
}

func (s TasksService) DeleteTask(taskID int, userID int) (*models.Response, *models.ErrorResponse) {
	status, err := s.repo.DeleteTask(taskID, userID)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status}, nil
}
