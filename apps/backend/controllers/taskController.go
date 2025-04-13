package controllers

import (
    "net/http"
    "time"
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/models"
)

// GET /tasks
func GetAllTasks(c *gin.Context) {
    var tasks []models.Task
    if err := database.DB.Preload("CreatedBy").Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

// GET /tasks/:id
func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    var task models.Task

    if err := database.DB.Preload("CreatedBy").First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

// POST /tasks
func CreateTask(c *gin.Context) {
    var input struct {
        Title       string `json:"title" binding:"required"`
        Description string `json:"description"`
        DueDate     string `json:"due_date"`
        CreatedByID uint   `json:"created_by_id"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    parsedDate, err := time.Parse("2006-01-02", input.DueDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format. Use YYYY-MM-DD."})
        return
    }

    task := models.Task{
        Title:       input.Title,
        Description: input.Description,
        DueDate:     parsedDate,
        CreatedByID: input.CreatedByID,
    }

    if err := database.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("Task Created")
    c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task

    if err := database.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    var input struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        DueDate     string `json:"due_date"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Title != "" {
        task.Title = input.Title
    }
    if input.Description != "" {
        task.Description = input.Description
    }
    if input.DueDate != "" {
        parsedDate, err := time.Parse("2006-01-02", input.DueDate)
        if err == nil {
            task.DueDate = parsedDate
        }
    }

    database.DB.Save(&task)
    fmt.Println("Task Updated")
    c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task

    if err := database.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    database.DB.Delete(&task)
    fmt.Println("Task Deleted")
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
