package handlers

import (
    "mini-erp/internal/models"
    "net/http"
    "strconv"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

func CreateTask(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var task models.Task
        if err := c.ShouldBindJSON(&task); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&task).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "작업 생성 실패"})
            return
        }
        c.JSON(http.StatusOK, task)
    }
}

func GetTasks(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var tasks []models.Task
        if err := db.Find(&tasks).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "작업 불러오기 실패"})
            return
        }
        c.JSON(http.StatusOK, tasks)
    }
}

func UpdateTaskStatus(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var input struct {
            Status string `json:"status"`
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var task models.Task
        if err := db.First(&task, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "작업을 찾을 수 없습니다."})
            return
        }

        task.Status = input.Status
        db.Save(&task)

        c.JSON(http.StatusOK, task)
    }
}

func DeleteTask(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        db.Delete(&models.Task{}, id)
        c.JSON(http.StatusOK, gin.H{"message": "작업이 삭제되었습니다."})
    }
}