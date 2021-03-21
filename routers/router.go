package routes

import (
	"github.com/bluepongo/lifeboat_go_backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", controllers.GetRoleByID)
}
