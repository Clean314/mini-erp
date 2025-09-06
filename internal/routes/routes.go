package routes

import (
    "github.com/gin-gonic/gin"
    "mini-promise/internal/handlers"
    "mini-promise/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
    // 인증
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    // 인증 요구 API
    auth := r.Group("/api", middlewares.AuthMiddleware())
    {
        auth.POST("/projects", handlers.CreateProject)
        auth.GET("/projects", handlers.GetProjects)
        auth.DELETE("/projects/:id", handlers.DeleteProject)

		auth.POST("/tasks", handlers.CreateTask)
		auth.GET("/projects/:projectId/tasks", handlers.GetTasksByProject)
		auth.PUT("/tasks/:id/status", handlers.UpdateTaskStatus)
		auth.DELETE("/tasks/:id", handlers.DeleteTask)
    }
}