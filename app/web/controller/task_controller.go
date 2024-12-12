package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type TaskController struct {
	taskAppService *service.TaskAppService
}

func NewTaskController(taskAppService *service.TaskAppService) *TaskController {
	return &TaskController{taskAppService}
}

func (tc *TaskController) Create(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(*entity.User)

	task, err := tc.taskAppService.CreateTask(user.Id, req.Title, req.Description, req.Reward, req.Deadline)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var res CreateTaskResponse
	res.Id = task.TaskId.String()
	res.Title = task.Title
	res.Description = task.Description
	res.IsDone = task.IsDone
	res.Reward = task.Reward
	res.Elapsed = task.Elapsed
	res.Deadline = task.Deadline
	res.CreatedAt = task.CreatedAt
	res.UpdatedAt = task.UpdatedAt

	c.JSON(http.StatusCreated, res)
}

func (tc *TaskController) GetAll(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	tasks, err := tc.taskAppService.GetTasks(user.Id)
	if err == sql.ErrNoRows {
		c.Status(http.StatusNoContent)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if (*tasks)[0].UserId != user.Id {
		c.Status(http.StatusForbidden)
		return
	}

	var res GetAllTaskResponse
	res.Tasks = make([]TaskResponse, 0)
	for _, task := range *tasks {
		res.Tasks = append(res.Tasks, TaskResponse{
			Id:          task.TaskId.String(),
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

func (tc *TaskController) Get(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	taskId := c.Param("id")

	task, err := tc.taskAppService.GetTask(taskId)
	if err == sql.ErrNoRows {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if task.UserId != user.Id {
		c.Status(http.StatusForbidden)
		return
	}

	var res GetTaskResponse
	res.Id = task.TaskId.String()
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

func (tc *TaskController) Put(c *gin.Context) {
	var req UpdateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	taskId := c.Param("id")

	task, err := tc.taskAppService.GetTask(taskId)
	if err == sql.ErrNoRows {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := c.MustGet("user").(*entity.User)

	if task.UserId != user.Id {
		c.Status(http.StatusForbidden)
		return
	}

	if task.Title == req.Title && task.Description == req.Description && task.IsDone == req.IsDone && task.Reward == req.Reward && task.Elapsed == req.Elapsed && task.Deadline == req.Deadline {
		c.Status(http.StatusBadRequest)
		return
	}

	err = tc.taskAppService.UpdateTask(task, req.Title, req.Description, req.IsDone, req.Reward, req.Elapsed, req.Deadline)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (tc *TaskController) Delete(c *gin.Context) {
	taskId := c.Param("id")

	task, err := tc.taskAppService.GetTask(taskId)
	if err == sql.ErrNoRows {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := c.MustGet("user").(*entity.User)

	if task.UserId != user.Id {
		c.Status(http.StatusForbidden)
		return
	}

	err = tc.taskAppService.DeleteTask(task.Id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
