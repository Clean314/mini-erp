package handlers

import (
	"net/http"
	"strconv"
	"time"

	"mini-erp/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateApproval(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ProcessType string `json:"process_type" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetUint("userID")
		approval := models.Approval{
			ProcessType: req.ProcessType,
			Status:      "REQUESTED",
			RequestedBy: userID,
		}

		if err := db.Create(&approval).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "승인 요청이 생성되었습니다."})
			return
		}

		c.JSON(http.StatusOK, approval)
	}
}

func ApproveApproval(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var approval models.Approval
		if err := db.First(&approval, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "승인 요청을 찾을 수 없습니다."})
			return
		}

		userID := c.GetUint("userID")
		approval.Status = "APPROVED"
		approval.ApprovedBy = &userID
		approval.UpdatedAt = time.Now()
		db.Save(&approval)

		c.JSON(http.StatusOK, approval)
	}
}

func RejectApproval(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var approval models.Approval
		if err := db.First(&approval, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "승인 요청을 찾을 수 없습니다."})
			return
		}

		userID := c.GetUint("userID")
		approval.Status = "REJECTED"
		approval.ApprovedBy = &userID
		approval.UpdatedAt = time.Now()
		db.Save(&approval)

		c.JSON(http.StatusOK, approval)
	}
}