package projects

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
)

// GetProjectByID retrieves a specific project by its ID.
// It returns the project details if found, or a 404 Not Found error if the project does not exist.
func ReadProjectByID(c *gin.Context) {
	var project models.Project
	id, match := c.Params.Get("id")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get project id"})
		return
	}

	dbconn := database.CreateDbConnection()
	if tx := dbconn.Preload("Tasks").Preload("Engineers").First(&project, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		}
		return
	}

	c.JSON(http.StatusOK, project)
}

// ReadAllProjects retrieves a list of all projects from the database.
// It returns the list of projects as a JSON response with a 200 OK status.
func ReadAllProjects(c *gin.Context) {
	var projects []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	dbconn := database.CreateDbConnection()
	dbconn.Model(&models.Project{}).Select("id", "name").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// CreateProject creates a new project in the database.
// It accepts a JSON request body containing the project's name and other details.
// On success, it returns the created project with a 201 Created status.
// If the input is invalid, it returns a 400 Bad Request error.
func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dbconn := database.CreateDbConnection()
	if tx := dbconn.Create(&project); tx.Error != nil {
		fmt.Println(tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a new project"})
		return
	}
	c.JSON(http.StatusCreated, project)
}