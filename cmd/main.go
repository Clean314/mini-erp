package main

import (
	"log"
	"mini-erp/internal/db"
	"mini-erp/internal/routes"

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