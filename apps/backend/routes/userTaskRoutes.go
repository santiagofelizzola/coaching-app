package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
	"github.com/santiagofelizzola/coaching-app/middleware"
)

func UserTaskRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/user-tasks", controllers.GetUserTasks)
		auth.GET("/user-tasks/summary", controllers.GetUserTaskStatusSummary)
		auth.POST("/assign-tasks", controllers.AssignTaskToUser)
		auth.PATCH("/user-tasks/:id/complete", controllers.MarkTaskComplete)
		
	}
}