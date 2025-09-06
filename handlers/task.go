package handlers

import (
    "mini-promise/internal/db"
    "mini-promise/internal/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.DB.Create(&task)
    c.JSON(http.StatusOK, task)
}

func GetTasksByProject(c *gin.Context) {
    projectID, _ := strconv.Atoi(c.Param("projectId"))

    var tasks []models.Task
    db.DB.Where("project_id = ?", projectID).Find(&tasks)

    c.JSON(http.StatusOK, tasks)
}

func UpdateTaskStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var input struct {
        Status string `json:"status"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var task models.Task
    if err := db.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }

    task.Status = input.Status
    db.DB.Save(&task)

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    db.DB.Delete(&models.Task{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
