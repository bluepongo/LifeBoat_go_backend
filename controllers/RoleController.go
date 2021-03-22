package controllers

import (
	"github.com/gin-gonic/gin"
	"lifeboat_go_backend/models"
	"net/http"
	"strconv"
)

// Get role by id
func GetRoleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	roleModel := models.Role{}
	role, err := roleModel.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

// Get all roles
func GetAllRoles(c *gin.Context) {
	roleModel := models.Role{}
	roles, err := roleModel.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": roles,
	})
}
