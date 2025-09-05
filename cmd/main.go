package main

import (
	"mini-erp/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
