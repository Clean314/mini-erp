package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	
	"mini-erp/internal/models"
	"mini-erp/internal/db"
)

var SecretKey = []byte("cnvuiqawhrh38hhkasbfb")

func Register(c *gin.Context) {

	var input sturct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{
		Username: input.Username,
		Password: string(hashed),
	}

	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "회원가입 성공"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db.DB.Where("username = ?", input.Username).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "잘못된 비밀번호입니다."})
        return
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "expire":     time.Now().Add(time.Hour * 72).Unix(),
    })
    tokenStr, _ := token.SignedString(SecretKey)

    c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}