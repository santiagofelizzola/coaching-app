package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/models"
)

func GetSessionsByUser(c *gin.Context) {
    userID := c.Query("user_id")
    var sessions []models.Session

    if err := database.DB.Where("user_id = ?", userID).Find(&sessions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, sessions)
}