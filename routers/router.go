package routes

import (
	"github.com/bluepongo/lifeboat_go_backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	groupRoles := router.Group("/api/v1/lifeboat")
	{
		groupRoles.GET("/roles/:name", controllers.GetRoleByName)
		groupRoles.GET("/roles", controllers.GetAllRoles)
		groupRoles.POST("/roles", controllers.AddRole)
		groupRoles.POST("/roles/:id", controllers.UpdateRoleByID)
		groupRoles.DELETE("/roles/:id", controllers.DeleteUserByID)
	}
}
