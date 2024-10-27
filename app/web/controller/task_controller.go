package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type TaskController struct {
	taskAppService *service.TaskAppService
}

func NewTaskController(taskAppService *service.TaskAppService) *TaskController {
	return &TaskController{taskAppService}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	userId := uuid.MustParse(c.GetString("userId"))

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskAppService.CreateTask(userId, req.Title, req.Description, req.Reward, req.Deadline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res CreateTaskResponse
	res.Id = task.Id.String()
	res.Title = task.Title
	res.Description = task.Description
	res.Reward = task.Reward
	res.Deadline = task.Deadline
	res.CreatedAt = task.CreatedAt

	c.JSON(http.StatusCreated, res)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	var req GetTaskRequest

	userId := uuid.MustParse(c.GetString("userId"))

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskAppService.GetTask(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	var res GetTaskResponse

	res.Id = task.Id.String()
	res.Title = task.Title
	res.Description = task.Description
	res.IsDone = task.IsDone
	res.Reward = task.Reward
	res.Elapsed = task.Elapsed
	res.Deadline = task.Deadline
	res.CreatedAt = task.CreatedAt
	res.UpdatedAt = task.UpdatedAt

	c.JSON(http.StatusOK, res)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	userId := c.GetString("userId")

	tasks, err := tc.taskAppService.GetTasks(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res GetTasksResponse
	for _, task := range tasks {
		res.Tasks = append(res.Tasks, GetTaskResponse{
			Id:          task.Id.String(),
			Title:       task.Title,
			Description: task.Description,
			IsDone:      task.IsDone,
			Reward:      task.Reward,
			Elapsed:     task.Elapsed,
			Deadline:    task.Deadline,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var req UpdateTaskRequest

	userId := uuid.MustParse(c.GetString("userId"))

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskAppService.GetTask(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err = tc.taskAppService.UpdateTask(req.Id, req.Title, req.Description, req.IsDone, req.Reward, req.Elapsed, req.Deadline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var req DeleteTaskRequest

	userId := uuid.MustParse(c.GetString("userId"))

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskAppService.GetTask(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	err = tc.taskAppService.DeleteTask(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
