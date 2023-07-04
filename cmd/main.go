package main

import (
	"app/framework"
	"app/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/v1/students", framework.getUserHandler)
	r.GET("/v1/students/search", framework.searchHandler)
	r.Run(":8080")
	defer database.CloseConnection()
}
