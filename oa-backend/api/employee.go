package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"oa-backend/utils/pwd"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddEmployee(c *gin.Context) {
	var employee models.Employee
	_ = c.ShouldBindJSON(&employee)

	_, code = models.SelectEmployeeByPhone(employee.Phone)

	if code == msg.FAIL {
		//默认密码等于编号+手机号
		employee.Password = employee.Number + employee.Phone
		employee.Password, err = pwd.Scrypt(employee.Password)

		if err == nil {
			employeeCre := models.Employee{
				ID:            employee.ID,
				IsDelete:      false,
				Phone:         employee.Phone,
				Name:          employee.Name,
				Password:      employee.Password,
				WechatID:      employee.WechatID,
				Email:         employee.Email,
				OfficeID:      employee.OfficeID,
				Number:        employee.Number,
				ContractCount: 0,
				PreCount:      0,
				Money:         employee.Money,
				Credit:        employee.Credit,
				OfficeCredit:  employee.Credit,
				RoleCredit:    employee.RoleCredit,
				RoleCreditDel: 0,
				Roles:         employee.Roles,
			}
			code = models.GeneralInsert(&employeeCre)
		} else {
			code = msg.ERROR
		}
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func DelEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.GeneralDelete(&models.Employee{}, id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

// 修改员工基本信息
func EditEmployeeBase(c *gin.Context) {
	var employee models.Employee
	var maps = make(map[string]interface{})
	_ = c.ShouldBindJSON(&employee)

	maps["name"] = employee.Name
	maps["phone"] = employee.Phone
	maps["wechat_id"] = employee.WechatID
	maps["email"] = employee.Email

	code = models.GeneralUpdate(&models.Employee{}, employee.ID, maps)
	msg.Message(c, code, nil)
}

// 修改员工资金信息
func EditEmployeeExpense(c *gin.Context) {
	var employee models.Employee
	// var maps = make(map[string]interface{})
	_ = c.ShouldBindJSON(&employee)

	// maps["money"] = employee.Money
	// maps["credit"] = employee.Credit
	// maps["office_credit"] = employee.OfficeCredit

	// code = models.GeneralUpdate(&models.Employee{}, employee.ID, maps)
	code = models.UpdateEmployeeMoney(&employee, c.MustGet("employeeID").(int))

	msg.Message(c, code, nil)
}

// 修改员工人事信息
func EditEmployeeOffice(c *gin.Context) {
	var employee models.Employee
	_ = c.ShouldBindJSON(&employee)

	code = models.UpdateEmployeeOffice(&employee)
	msg.Message(c, code, nil)
}

func QueryEmployee(c *gin.Context) {
	var employee models.Employee
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		employee, code = models.SelectEmployee(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, employee)
}

func QueryEmployees(c *gin.Context) {
	var employeeQuery models.Employee
	_ = c.ShouldBindJSON(&employeeQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectEmployees(&employeeQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllEmployee(c *gin.Context) {
	var employeeQuery models.Employee
	var employees []models.Employee
	_ = c.ShouldBindJSON(&employeeQuery)
	var maps = make(map[string]interface{})
	if employeeQuery.OfficeID != 0 {
		maps["office_id"] = employeeQuery.OfficeID
	}
	code = models.GeneralSelectAll(&employees, maps)
	msg.Message(c, code, employees)
}

func ResetEmployeePwd(c *gin.Context) {
	var employee models.Employee
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.GeneralSelect(&employee, id, nil)
	} else {
		code = msg.ERROR
	}

	employee.Password, err = pwd.Scrypt(employee.Phone)

	if err == nil {
		var maps = make(map[string]interface{})
		maps["password"] = employee.Password
		code = models.GeneralUpdate(&employee, employee.ID, maps)
	} else {
		code = msg.ERROR
	}

	msg.Message(c, code, nil)
}
