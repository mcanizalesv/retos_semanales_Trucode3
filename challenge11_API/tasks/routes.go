package tasks

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {
	router.POST("/tasks", CreateTask)
}