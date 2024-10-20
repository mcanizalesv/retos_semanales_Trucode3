package projects

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {
	router.GET("/projects", ReadAllProjects)
	router.GET("/projects/:id", ReadProjectByID)
	router.POST("/projects", CreateProject)
}