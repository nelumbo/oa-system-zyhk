package api

import (
	"oa-backend/models"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func QueryAllPermission(c *gin.Context) {
	var permissions []models.Permission
	code = models.GeneralSelectAll2(&permissions, nil)
	msg.Message(c, code, permissions)
}
