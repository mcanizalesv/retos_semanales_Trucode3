package engineers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
)

// UpdateProjectEngineer updates the project assignment for a specific engineer.
// It takes the engineer ID from the URL parameters and the new project ID from the request body.
// It checks if both the engineer and the new project exist before making the update.
// Returns a 200 OK status on success or an appropriate error message if something goes wrong.
func UpdateProjectEngineer(c *gin.Context) {
    var engineer models.Engineer
    var input struct {
        ProjectID uint `json:"project_id"`
    }

    engineerID, match := c.Params.Get("id")
    if !match {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get engineer id"})
        return
    }

    dbconn := database.CreateDbConnection()

    if tx := dbconn.First(&engineer, engineerID); tx.Error != nil {
        if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Engineer not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": tx.Error.Error()})
        }
        return
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var project models.Project
    if tx := dbconn.First(&project, input.ProjectID); tx.Error != nil {
        if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": tx.Error.Error()})
        }
        return
    }

    if tx := dbconn.Model(&engineer).Update("project_id", input.ProjectID); tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update engineer's project", "details": tx.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Engineer reassigned to new project successfully", "engineer": engineer})
}

// GetEngineerByID retrieves a specific engineer by their ID, including their current project.
// Returns the engineer's details along with the associated project or a 404 Not Found error if not found.
func ReadEngineerByID(c *gin.Context) {
	var engineer models.Engineer

	engineerID, match := c.Params.Get("id")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get engineer id"})
		return
	}

	dbconn := database.CreateDbConnection()

	if err := dbconn.Preload("Project").First(&engineer, engineerID).Error; err != nil {
		fmt.Println("Error details:", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Engineer not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve engineer"})
		}
		return
	}

	c.JSON(http.StatusOK, engineer)
}


// ReadAllEngineers retrieves a list of all engineers from the database.
// It returns the list of engineers as a JSON response with a 200 OK status.
func ReadAllEngineers(c *gin.Context) {
	var engineer []models.Engineer
	dbconn := database.CreateDbConnection()
	dbconn.Find(&engineer)
	c.JSON(http.StatusOK, engineer)
}

// CreateEngineer creates a new engineer in the database.
// It expects a JSON request body containing the engineer's name, surname, and associated project ID.
// On success, it returns the created engineer along with the associated project details with a 201 Created status.
// If the input is invalid or there is an error, it returns a 400 Bad Request error
func CreateEngineer(c *gin.Context) {
	var engineer models.Engineer
	if err := c.BindJSON(&engineer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	dbconn := database.CreateDbConnection()

	tx := dbconn.Create(&engineer)
	if tx.Error != nil {
		fmt.Println("Error al crear ingeniero:", tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a new engineer"})
		return
	}

	var project models.Project
	if err := dbconn.First(&project, engineer.ProjectID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find project"})
		return
	}

	engineer.Project = project

	c.JSON(http.StatusCreated, engineer)
}
