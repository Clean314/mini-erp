package main

import (
	"log"
	"mini-promise/internal/db"
	"mini-promise/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, database)

	r.Run(":8080")
}