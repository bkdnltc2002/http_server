package main

import (
	"app/framework"
	"github.com/gin-gonic/gin"
)

func main() {
	framework.Route(gin.Default())
}
