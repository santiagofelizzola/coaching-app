package controllers

import (
    "net/http"
    "time"
    "fmt"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/models"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"`
}

func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
        Role:     input.Role,
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    fmt.Println(user.Name + " created")
    c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    session := models.Session{
        UserID:    user.ID,
        Token:     tokenString,
        ExpiresAt: time.Now().Add(time.Hour * 72),
    }

    database.DB.Create(&session)
    fmt.Println("Welcome " + user.Name)
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
        return
    }

    if err := database.DB.Where("token = ?", token).Delete(&models.Session{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
        return
    }

    fmt.Println("Take care!")
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func RefreshToken(c *gin.Context) {
    oldToken := c.GetHeader("Authorization")
    if oldToken == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
        return
    }

    var session models.Session
    if err := database.DB.Where("token = ?", oldToken).First(&session).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
        return
    }

    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": session.UserID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := newToken.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
        return
    }

    session.Token = tokenString
    session.ExpiresAt = time.Now().Add(time.Hour * 72)
    database.DB.Save(&session)

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}