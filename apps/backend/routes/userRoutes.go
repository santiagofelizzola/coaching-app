package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
	"github.com/santiagofelizzola/coaching-app/middleware"
)

func UserRoutes(router *gin.Engine) {
	auth := router.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/users", controllers.GetAllUsers)
		auth.GET("/users/:id", controllers.GetUserByID)
		auth.POST("/users", controllers.CreateUser)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.DELETE("/users/:id", controllers.DeleteUser)
	}
}