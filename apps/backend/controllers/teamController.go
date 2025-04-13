package controllers

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/models"
)

// GET /teams
func GetAllTeams(c *gin.Context) {
	var teams []models.Team
	result := database.DB.Find(&teams)

	if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	fmt.Println("Getting all teams!")
	c.JSON(http.StatusOK, teams)
}

// GET /teams/:id
func GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := database.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// POST /teams
func CreateTeam(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team := models.Team{
		Name: input.Name,
	}

	if err := database.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// PUT /teams/:id
func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := database.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	var input struct {
		Name string `json:"name" binding:"required"`
	}

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	team.Name = input.Name

	database.DB.Save(&team)

	c.JSON(http.StatusOK, team)
}

// DELETE /teams/:id
func DeleteTeam(c *gin.Context) {
    id := c.Param("id")
    var team models.Team

    if err := database.DB.First(&team, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
        return
    }

    database.DB.Delete(&team)
    c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}

// POST /teams/:id/user
func AddUserToTeam(c *gin.Context) {
    var input struct {
        UserID uint `json:"user_id" binding:"required"`
    }

	id := c.Param("id")
	var team models.Team
	if err := database.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Model(&team).Association("Users").Append(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to team"})
		return
	}
	
	fmt.Println(user.Name + " added to " + team.Name)
	c.JSON(http.StatusOK, gin.H{"message": "User added to team"})
}