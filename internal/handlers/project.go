package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"mini-erp/internal/models"
	"mini-erp/internal/db"
)

func CreateProject(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project models.Project
		if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&project).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "프로젝트 생성 실패"})
			return
		}
		c.JSON(http.StatusOK, project)
	}
}

func GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var projects []models.Project
		if err := db.Find(&projects).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "프로젝트 불러오기 실패"})
			return
		}
		c.JSON(http.StatusOK, projects)
	}
}

func DeleteProject(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    db.DB.Delete(&models.Project{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "프로젝트 삭제됨"})
}