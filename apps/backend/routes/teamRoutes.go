package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
	"github.com/santiagofelizzola/coaching-app/middleware"
)

func TeamRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/teams", controllers.GetAllTeams)
		auth.GET("/teams/:id", controllers.GetTeamByID)
		auth.POST("/teams", controllers.CreateTeam)
		auth.PUT("/teams/:id", controllers.UpdateTeam)
		auth.DELETE("/teams/:id", controllers.DeleteTeam)
		auth.POST("teams/:id/user", controllers.AddUserToTeam)
	}
}