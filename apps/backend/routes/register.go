package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	AuthRoutes(r)
	SessionRoutes(r)
	TeamRoutes(r)
	TaskRoutes(r)
	UserRoutes(r)
	UserTaskRoutes(r)
}