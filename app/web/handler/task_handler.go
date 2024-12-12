package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/web/controller"
)

func NewTaskHandler(ctx context.Context, taskGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	taskController := controller.NewTaskController(service.NewTaskAppService(ctx, repository.NewTaskRepository(sqlHandler)))

	taskGroup.POST("/", taskController.Create)
	taskGroup.GET("/", taskController.GetAll)
	taskGroup.GET("/:task_id", taskController.Get)
	taskGroup.PUT("/:task_id", taskController.Put)
	taskGroup.DELETE("/:task_id", taskController.Delete)
}
