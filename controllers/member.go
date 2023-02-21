package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/models"
	"net/http"
)

type MemberController struct{}

var memberModel = new(models.Member)
var memberDeviceModel = new(models.MemberDevice)

func (u MemberController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		member, err := memberModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "member": member})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
