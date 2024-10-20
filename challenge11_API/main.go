package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/engineers"
	"trucode.app/apitaskmanager/models"
	"trucode.app/apitaskmanager/projects"
	"trucode.app/apitaskmanager/tasks"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env vars")
	}

	dbconn := database.CreateDbConnection()
	dbconn.AutoMigrate(&models.Project{}, &models.Task{}, &models.Engineer{})
	router := gin.Default()
	projects.AddRoutes(router)
	tasks.AddRoutes(router)
	engineers.AddRoutes(router)
	router.Run("0.0.0.0:2222")
}