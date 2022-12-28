package api

import (
	"oa-backend/middleware"
	"oa-backend/models"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

var code int
var err error

func Login(c *gin.Context) {
	var employee models.Employee
	var token string
	_ = c.ShouldBindJSON(&employee)

	employee, code = models.SelectEmployeeByPhoneAndPwd(employee.Phone, employee.Password)

	if code == msg.SUCCESS {
		token, _ = middleware.GenerateToken(employee.ID)
	}

	msg.Message(c, code, gin.H{
		"employee": employee,
		"token":    token,
	})
}
