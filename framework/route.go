package framework

import (
	"app/handler"
	"app/database"
	"github.com/gin-gonic/gin"
)

func Route(g *gin.Engine) {
	g.GET("/v1/students", handler.GetUserHandler)
	g.GET("/v1/students/search", handler.SearchHandler)
	g.Run(":8080")
	defer database.CloseConnection()
}
