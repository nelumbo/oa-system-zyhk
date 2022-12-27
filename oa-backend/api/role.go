package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRole(c *gin.Context) {
	var role models.Role
	_ = c.ShouldBindJSON(&role)

	code = models.GeneralInsert(&role)
	msg.Message(c, code, nil)
}

func DelRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.DeleteRole(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditRole(c *gin.Context) {
	var role models.Role
	_ = c.ShouldBindJSON(&role)

	code = models.UpdateRole(&role)
	msg.Message(c, code, nil)
}

func QueryRole(c *gin.Context) {
	var role models.Role
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		role, code = models.SelectRole(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, role)
}

func QueryRoles(c *gin.Context) {
	var roleQuery models.Role
	_ = c.ShouldBindJSON(&roleQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectRoles(&roleQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllRole(c *gin.Context) {
	var roles []models.Role
	code = models.GeneralSelectAll(&roles, nil)
	msg.Message(c, code, roles)
}
