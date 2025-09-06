package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"mini-erp/internal/models"
	"mini-erp/internal/db"
)

func CreateProject(c *gin.Context) {

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")
    project.OwnerID = uint(userID)

    db.DB.Create(&project)
    c.JSON(http.StatusOK, project)
}

func GetProjects(c *gin.Context) {
    var projects []models.Project
    db.DB.Find(&projects)
    c.JSON(http.StatusOK, projects)
}

func DeleteProject(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    db.DB.Delete(&models.Project{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}