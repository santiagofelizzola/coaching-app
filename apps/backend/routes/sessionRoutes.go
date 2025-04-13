package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
	"github.com/santiagofelizzola/coaching-app/middleware"
)

func SessionRoutes(r *gin.Engine) {
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/session", controllers.GetSessionsByUser)
		
	}
}