package engineers

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {
	router.POST("/engineers", CreateEngineer)
	router.GET("/engineers/:id", ReadEngineerByID)
	router.GET("/engineers", ReadAllEngineers)
	router.PUT("/engineers/:id", UpdateProjectEngineer)
}
