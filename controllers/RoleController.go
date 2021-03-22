package controllers

import (
	"github.com/bluepongo/lifeboat_go_backend/models"
	"github.com/gin-gonic/gin"
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

// Add a role
func AddRole(c *gin.Context) {
	roleModel := models.Role{}
	if err := c.ShouldBind(&roleModel); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	id, err := roleModel.AddRole()
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":     "success",
		"role_id": id,
	})
}
