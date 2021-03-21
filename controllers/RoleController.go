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
	data, err := roleModel.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
