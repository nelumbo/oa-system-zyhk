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

	_ = models.SelectAllPermission(&employee)

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

func QueryMyPredesigns(c *gin.Context) {
	var predesignQuery models.Predesign
	_ = c.ShouldBindJSON(&predesignQuery)
	predesignQuery.EmployeeID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectPredesigns(&predesignQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryMyPredesignTasks(c *gin.Context) {
	var predesignTaskQuery models.PredesignTask
	_ = c.ShouldBindJSON(&predesignTaskQuery)
	predesignTaskQuery.EmployeeID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectPredesignTasks(&predesignTaskQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryMyContracts(c *gin.Context) {
	var contactQuery models.Contract
	_ = c.ShouldBindJSON(&contactQuery)
	contactQuery.EmployeeID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectContracts(&contactQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryMySaveContracts(c *gin.Context) {
	xForms := ginUtil.GinArrayPreprocessing(c)
	xForms.Data, code = models.SelectMySaveContracts(c.MustGet("employeeID").(int), &xForms)
	msg.Message(c, code, xForms)
}

func QueryMyTasks(c *gin.Context) {
	var taskQuery models.Task
	_ = c.ShouldBindJSON(&taskQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectMyTasks(&taskQuery, c.MustGet("employeeID").(int), &xForms)
	msg.Message(c, code, xForms)
}

func QueryMySavePurchasings(c *gin.Context) {
	var purchasing models.Purchasing
	var purchasings []models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	purchasings, code = models.SelectAllSavePurchasings(purchasing.ContractID, purchasing.TaskID, c.MustGet("employeeID").(int))

	msg.Message(c, code, purchasings)
}

func QueryMyHistorys(c *gin.Context) {
	var historyEmployee models.HistoryEmployee
	_ = c.ShouldBindJSON(&historyEmployee)
	historyEmployee.UserID = c.MustGet("employeeID").(int)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectHistoryEmployees(&historyEmployee, &xForms)
	msg.Message(c, code, xForms)
}
