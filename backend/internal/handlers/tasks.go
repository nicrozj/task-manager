package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TasksHandlers struct {
	svc services.TasksServiceInterface
}

func NewTasksHandlers() *TasksHandlers {
	return &TasksHandlers{
		svc: services.NewTasksService(),
	}
}

func (h TasksHandlers) CreateTask(c *gin.Context) {
	var taskRequest models.TaskRequest

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("failed to create task")))
		return
	}

	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.CreateTask(&taskRequest, userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}
func (h TasksHandlers) GetTaskByID(c *gin.Context) {
	taskIDStr := c.Param("id")

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("failed to get task")))
		return
	}

	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, er := h.svc.GetTaskByID(taskID, userID)
	if er != nil {
		c.JSON(er.Status, er)
		return
	}

	c.JSON(result.Status, result)
}
func (h TasksHandlers) GetTasks(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.GetTasks(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}

func (h TasksHandlers) UpdateTask(c *gin.Context) {
	var taskRequest models.TaskRequest

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("failed to create task")))
		return
	}

	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("failed to delete task")))
		return
	}

	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, er := h.svc.UpdateTask(&taskRequest, userID, taskID)
	if er != nil {
		c.JSON(er.Status, er)
		return
	}

	c.JSON(result.Status, result)
}

func (h TasksHandlers) DeleteTask(c *gin.Context) {
	taskIDStr := c.Param("id")

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("failed to delete task")))
		return
	}

	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, er := h.svc.DeleteTask(taskID, userID)
	if er != nil {
		c.JSON(er.Status, er)
		return
	}

	c.JSON(result.Status, result)
}
