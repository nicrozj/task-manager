package repository

import (
	"backend/internal/config"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TasksRepo struct {
	db *sqlx.DB
}

func NewTasksRepo() TasksRepositoryInterface {
	return &TasksRepo{
		db: config.NewAppConfig().DB,
	}
}

func (r TasksRepo) CreateTask(task *models.Task) (int, int, error) {
	var id int

	query := `
        INSERT INTO tasks (user_id, title, description, status)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	err := r.db.QueryRow(
		query,
		task.UserId,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&id)

	if err != nil {
		return 0, http.StatusInternalServerError, fmt.Errorf("failed to create task: %s", err)
	}

	return id, http.StatusOK, nil
}

func (r TasksRepo) GetTaskByID(taskID int, userID int) (*models.Task, int, error) {
	var task models.Task

	fmt.Println(taskID, userID)
	query := `
        SELECT id, user_id, title, description, status, created_at 
        FROM tasks 
        WHERE id = $1 AND user_id = $2`

	err := r.db.QueryRow(query, taskID, userID).Scan(
		&task.Id,
		&task.UserId,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, fmt.Errorf("task not found: %s", err)
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get task")
	}

	return &task, http.StatusOK, nil
}

func (r TasksRepo) GetTasks(userID int) ([]*models.Task, int, error) {
	var tasks []*models.Task

	query := `
        SELECT id, user_id, title, description, status, created_at 
        FROM tasks 
        WHERE user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to fetch tasks")
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
		); err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("failed to scan task")
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error during iteration")
	}

	return tasks, http.StatusOK, nil
}
func (r TasksRepo) UpdateTask(task *models.Task) (int, error) {
	query := `
        UPDATE tasks 
        SET title = $1, description = $2, status = $3 
        WHERE id = $4 AND user_id = $5 
        RETURNING id, user_id, title, description, status, created_at`

	var updatedTask models.Task

	fmt.Println(task.Id, task.UserId)

	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Id,
		task.UserId,
	).Scan(
		&updatedTask.Id,
		&updatedTask.UserId,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.Status,
		&updatedTask.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, fmt.Errorf("task not found or not owned by user: %s", err)
		}
		return http.StatusInternalServerError, fmt.Errorf("failed to update task")
	}

	return http.StatusOK, nil
}

func (r TasksRepo) DeleteTask(taskID int, userID int) (int, error) {
	query := `DELETE FROM tasks WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(query, taskID, userID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete task")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to check affected rows")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("task not found or not owned by user")
	}

	return http.StatusOK, nil
}
