package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/util/config"
	"github.com/kairo913/tasclock-server/app/web/controller"
)

func NewTaskHandler(ctx context.Context, taskGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	taskController := controller.NewTaskController(service.NewTaskAppService(ctx, repository.NewTaskRepository(sqlHandler)))

	sessonConfig := config.NewSessionConfig(ctx)

	taskGroup.Use(AuthMiddleware(sessonConfig.JWTSecret))

	taskGroup.POST("/create", taskController.CreateTask)
	taskGroup.GET("/get", taskController.GetTask)
	taskGroup.GET("/getall", taskController.GetTasks)
	taskGroup.DELETE("/delete", taskController.DeleteTask)
	taskGroup.PUT("/update", taskController.UpdateTask)
}
