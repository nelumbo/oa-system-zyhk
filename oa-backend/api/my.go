package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func QueryMe(c *gin.Context) {
	var employee models.Employee

	employee, code = models.SelectEmployee(c.MustGet("employeeID").(int))

	msg.Message(c, code, employee)
}

func QueryMyExpenses(c *gin.Context) {
	var expenseQuery models.Expense
	_ = c.ShouldBindJSON(&expenseQuery)
	expenseQuery.EmployeeID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectExpenses(&expenseQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryMyBidbonds(c *gin.Context) {
	var bidbondQuery models.Bidbond
	_ = c.ShouldBindJSON(&bidbondQuery)
	bidbondQuery.EmployeeID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectBidbonds(&bidbondQuery, &xForms)
	msg.Message(c, code, xForms)
}
