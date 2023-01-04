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
	expense.EmployeeID = c.MustGet("employeeID").(int)
	expense.CreateDate.Time = time.Now()

	licence := true

	if expense.Type == magic.EXPENSE_TYPE_BuZhu {
		var employee models.Employee
		employee, code = models.SelectEmployee(expense.EmployeeID)
		if code != msg.SUCCESS || employee.Money < expense.Amount {
			licence = false
			code = msg.FAIL
		}
	} else if expense.Type == magic.EXPENSE_TYPE_TiChen {
		var employee models.Employee
		employee, code = models.SelectEmployee(expense.EmployeeID)
		if code != msg.SUCCESS || employee.Office.Money < expense.Amount {
			licence = false
			code = msg.FAIL
		}
	}

	if licence {
		expense.Status = magic.EXPENSE_STATUS_NOT_APPROVAL_1
		code = models.GeneralInsert(&expense)
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

func EditExpense(c *gin.Context) {
	var expense, expenseBak models.Expense
	_ = c.ShouldBindJSON(&expense)
	code = models.GeneralSelect(&expenseBak, expense.ID, nil)

	if code == msg.SUCCESS && expenseBak.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_1 &&
		expenseBak.EmployeeID == c.MustGet("employeeID").(int) {
		var maps = make(map[string]interface{})
		maps["create_date"] = time.Now()
		maps["type"] = expense.Type
		maps["amount"] = expense.Amount
		maps["text"] = expense.Text
		code = models.GeneralUpdate(&models.Expense{}, expense.ID, maps)
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
