package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense
	_ = c.ShouldBindJSON(&expense)

	var employee models.Employee
	employee, code = models.SelectEmployee(c.MustGet("employeeID").(int))
	if code != msg.SUCCESS ||
		(expense.Type == magic.EXPENSE_TYPE_BuZhu && employee.Money < expense.Amount) ||
		(expense.Type == magic.EXPENSE_TYPE_TiChen && employee.Office.Money < expense.Amount) {
		code = msg.FAIL
	} else {
		expenseCre := models.Expense{
			ID:         expense.ID,
			IsDelete:   false,
			EmployeeID: c.MustGet("employeeID").(int),
			Type:       expense.Type,
			Text:       expense.Text,
			Amount:     expense.Amount,
			Status:     magic.EXPENSE_STATUS_NOT_APPROVAL_1,
			CreateDate: models.XDate{
				Time: time.Now(),
			},
		}
		code = models.GeneralInsert(&expenseCre)
	}

	msg.Message(c, code, nil)
}

func DelExpense(c *gin.Context) {
	var expenseBak models.Expense
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.GeneralSelect(&expenseBak, id, nil)
		if code == msg.SUCCESS && expenseBak.EmployeeID == c.MustGet("employeeID").(int) &&
			(expenseBak.Status == magic.EXPENSE_STATUS_FAIL || expenseBak.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_1) {
			code = models.GeneralDelete(&models.Expense{}, id)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func ApprovalExpense(c *gin.Context) {
	var expense, expenseBak models.Expense
	_ = c.ShouldBindJSON(&expense)
	code = models.GeneralSelect(&expenseBak, expense.ID, nil)

	if expenseBak.ID != 0 && expenseBak.Status == expense.Status &&
		expenseBak.Type == expense.Type && expenseBak.Amount == expense.Amount {
		var maps = make(map[string]interface{})
		maps["approve_date"] = time.Now()
		if expenseBak.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_1 {
			if expense.IsPass {
				maps["approver_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_NOT_APPROVAL_2
				code = models.GeneralUpdate(&models.Expense{}, expenseBak.ID, maps)
			} else {
				maps["approver_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_FAIL
				code = models.GeneralUpdate(&models.Expense{}, expenseBak.ID, maps)
			}
		} else if expenseBak.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_2 {
			if expense.IsPass {
				maps["finance_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_NOT_PAYMENT
				code = models.UpdateExpense(&expense, &expenseBak, maps)
			} else {
				maps["finance_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_FAIL
				code = models.GeneralUpdate(&models.Expense{}, expenseBak.ID, maps)
			}
		} else if expenseBak.Status == magic.EXPENSE_STATUS_NOT_PAYMENT {
			if expense.IsPass {
				maps["cashier_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_FINISH
				code = models.GeneralUpdate(&models.Expense{}, expenseBak.ID, maps)
			} else {
				maps["cashier_id"] = c.MustGet("employeeID").(int)
				maps["status"] = magic.EXPENSE_STATUS_FAIL
				code = models.UpdateExpense(&expense, &expenseBak, maps)
			}
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func QueryExpenses(c *gin.Context) {
	var expenseQuery models.Expense
	_ = c.ShouldBindJSON(&expenseQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectExpenses(&expenseQuery, &xForms)
	msg.Message(c, code, xForms)
}
