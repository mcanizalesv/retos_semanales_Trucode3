package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
)

// CreateTask creates a new task associated with a specific project.
// It expects a JSON request body containing the task's title, description, estimated hours, engineer ID, and project ID.
// On success, it returns the created task with a 201 Created status.
// If the project does not exist, it returns a 404 Not Found error.
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dbconn := database.CreateDbConnection()

	var project models.Project
    if err := dbconn.First(&project, task.ProjectID).Error; err != nil {
		fmt.Println("Project ID:", task.ProjectID)
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find project"})
        }
        return
    }

	if err := dbconn.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }


	c.JSON(http.StatusCreated, task)
}