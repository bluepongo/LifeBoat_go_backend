package controllers

import (
	"github.com/bluepongo/lifeboat_go_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get role by id
func GetRoleByName(c *gin.Context) {
	name := c.Param("name")
	roleModel := models.Role{}
	role, err := roleModel.GetRoleByName(name)
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

// Update the role info by id
func UpdateRoleByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}
	roleModel := models.Role{}
	if err := c.ShouldBind(&roleModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	_, err = roleModel.UpdateRoleByID(idInt)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"msg":     "success",
		"role_id": idInt,
	})
}

// Delete a role by id
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}
	roleModel := models.Role{}
	_, err = roleModel.DeleteRoleByID(idInt)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"msg":     "success",
		"role_id": idInt,
	})
}
