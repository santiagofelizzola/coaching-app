package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
	"github.com/santiagofelizzola/coaching-app/middleware"
)

func TaskRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetAllTasks)
		auth.GET("/tasks/:id", controllers.GetTaskByID)
		auth.POST("/tasks", controllers.CreateTask)
		auth.PUT("/tasks", controllers.UpdateTask)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}