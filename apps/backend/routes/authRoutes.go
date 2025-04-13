package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.POST("/refresh", controllers.RefreshToken)
}