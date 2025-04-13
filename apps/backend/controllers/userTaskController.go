package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/models"
)

// GET /user-tasks
func GetUserTasks(c *gin.Context) {
    userID := c.Query("user_id")
    var tasks []models.UserTask

    if err := database.DB.Preload("Task").Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

// PATCH /user-tasks/:id/complete
func MarkTaskComplete(c *gin.Context) {
    id := c.Param("id")
    var userTask models.UserTask

    if err := database.DB.First(&userTask, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User task not found"})
        return
    }

    now := time.Now()
    userTask.Status = "complete"
    userTask.CompletedAt = &now

    database.DB.Save(&userTask)
    c.JSON(http.StatusOK, userTask)
}

// POST /assign-task
func AssignTaskToUser(c *gin.Context) {
    var input struct {
        UserID uint `json:"user_id" binding:"required"`
        TaskID uint `json:"task_id" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    assignment := models.UserTask{
        UserID: input.UserID,
        TaskID: input.TaskID,
        Status: "pending",
    }

    if err := database.DB.Create(&assignment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, assignment)
}

// GET /user-tasks/summary
func GetUserTaskStatusSummary(c *gin.Context) {
    userID := c.Query("user_id")
    var assignments []models.UserTask

    if err := database.DB.Preload("Task").Where("user_id = ?", userID).Find(&assignments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var completed, overdue, pending int
    now := time.Now()

    for _, task := range assignments {
        if task.Status == "complete" {
            completed++
        } else if task.Task.DueDate.Before(now) {
            overdue++
        } else {
            pending++
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "completed": completed,
        "overdue":   overdue,
        "pending":   pending,
    })
}