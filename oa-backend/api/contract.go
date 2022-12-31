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

// 暂存合同
func SaveContract(c *gin.Context) {
	var contract models.Contract
	_ = c.ShouldBindJSON(&contract)

	contract.EmployeeID = c.MustGet("employeeID").(int)
	contract.Status = magic.CONTRACT_STATUS_SAVE
	contract.CreateDate.Time = time.Now()
	contract.TotalAmount = 0
	contract.Tasks = nil

	code = models.GeneralInsert(&contract)

	msg.Message(c, code, nil)
}

func AddContract(c *gin.Context) {
	var contract models.Contract
	var employee models.Employee
	_ = c.ShouldBindJSON(&contract)

	code = models.GeneralSelect(&employee, c.MustGet("employeeID").(int), nil)

	if code == msg.SUCCESS {
		contract.EmployeeID = c.MustGet("employeeID").(int)
		contract.OfficeID = employee.OfficeID
		contract.Status = magic.CONTRACT_STATUS_NOT_APPROVAL
		contract.CreateDate.Time = time.Now()

		code = models.GeneralInsert(&contract)
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func DelContract(c *gin.Context) {
	var contractBak models.Contract
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.GeneralSelect(&contractBak, id, nil)
		if code == msg.SUCCESS && contractBak.EmployeeID == c.MustGet("employeeID").(int) &&
			(contractBak.Status == magic.CONTRACT_STATUS_REJECT ||
				contractBak.Status == magic.CONTRACT_STATUS_SAVE ||
				contractBak.Status == magic.CONTRACT_STATUS_NOT_APPROVAL) {
			code = models.GeneralDelete(&models.Contract{}, id)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func ApproveContract(c *gin.Context) {
	var contract, contractBak models.Contract
	_ = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.EmployeeID, nil)

	if contractBak.Status == magic.CONTRACT_STATUS_NOT_APPROVAL {
		var maps = make(map[string]interface{})
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["audit_date"] = time.Now()
		code = models.ApproveContract(&contractBak, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func RejectContract(c *gin.Context) {
	var contract, contractBak models.Contract
	_ = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.EmployeeID, nil)

	if contractBak.Status == magic.CONTRACT_STATUS_NOT_FINISH {
		code = models.RejectContract(&contractBak, c.MustGet("employeeID").(int))
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func QueryContract(c *gin.Context) {
	var contract models.Contract
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		contract, code = models.SelectContract(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, contract)
}

func QueryContracts(c *gin.Context) {
	var contractQuery models.Contract
	_ = c.ShouldBindJSON(&contractQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectContracts(&contractQuery, &xForms)
	msg.Message(c, code, xForms)
}
