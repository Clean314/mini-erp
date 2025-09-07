package routes

import (
    "github.com/gin-gonic/gin"
    "mini-erp/internal/handlers"
    "mini-erp/internal/middlewares"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())

    // 인증 요구 API

	api.POST("/projects", middlewares.RoleMiddleware("USER", "MANAGER", "ADMIN"), handlers.CreateProject(db))
	api.GET("/projects", middlewares.RoleMiddleware("USER", "MANAGER", "ADMIN"), handlers.GetProjects(db))

	api.POST("/tasks", middlewares.RoleMiddleware("USER", "MANAGER", "ADMIN"), handlers.CreateTask(db))
	api.GET("/tasks", middlewares.RoleMiddleware("USER", "MANAGER", "ADMIN"), handlers.GetTasks(db))

	api.POST("/approvals", middlewares.RoleMiddleware("USER", "MANAGER", "ADMIN"), handlers.CreateApproval(db))
	api.POST("/approvals/:id/approve", middlewares.RoleMiddleware("MANAGER", "ADMIN"), handlers.ApproveApproval(db))
	api.POST("/approvals/:id/reject", middlewares.RoleMiddleware("MANAGER", "ADMIN"), handlers.RejectApproval(db))
}