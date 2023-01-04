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

	if contract.ID == 0 {
		contract.EmployeeID = c.MustGet("employeeID").(int)
		contract.Status = magic.CONTRACT_STATUS_SAVE
		contract.CreateDate.Time = time.Now()
		contract.TotalAmount = 0
		contract.Tasks = nil

		code = models.GeneralInsert(&contract)
	} else {
		var contractBak models.Contract
		_ = models.GeneralSelect(&contractBak, contract.ID, nil)
		if contractBak.ID == contract.ID && contractBak.EmployeeID == c.MustGet("employeeID").(int) {
			var maps = make(map[string]interface{})
			maps["region_id"] = contract.RegionID
			maps["customer_id"] = contract.CustomerID
			maps["vendor_id"] = contract.VendorID
			maps["contract_date"] = contract.ContractDate
			maps["estimated_delivery_date"] = contract.ContractDate
			maps["pay_type"] = contract.PayType
			maps["is_pre_deposit"] = contract.IsPreDeposit
			maps["pre_deposit_record"] = contract.PreDepositRecord
			maps["is_special"] = contract.IsSpecial
			maps["invoice_type"] = contract.InvoiceType
			maps["invoice_content"] = contract.InvoiceContent
			maps["payment_content"] = contract.PaymentContent
			maps["remark"] = contract.Remark
			code = models.GeneralUpdate(&models.Contract{}, contractBak.ID, maps)
		} else {
			code = msg.FAIL
		}
	}

	msg.Message(c, code, nil)
}

func AddContract(c *gin.Context) {
	var contract, contractBak models.Contract
	var employee models.Employee
	_ = c.ShouldBindJSON(&contract)

	code = models.GeneralSelect(&employee, c.MustGet("employeeID").(int), nil)

	if code == msg.SUCCESS {
		contract.OfficeID = employee.OfficeID
		contract.Status = magic.CONTRACT_STATUS_NOT_APPROVAL

		if contract.IsPreDeposit {
			contract.Tasks = nil
		} else {
			for i := range contract.Tasks {
				var product models.Product
				_ = models.GeneralSelect(&product, contract.Tasks[i].ProductID, nil)
				contract.TotalAmount += contract.Tasks[i].TotalPrice
				contract.Tasks[i].ProductAttributeID = product.AttributeID
			}
		}

		if contract.ID != 0 {
			_ = models.GeneralSelect(&contractBak, contract.ID, nil)
			if contractBak.ID == contract.ID && contractBak.EmployeeID == c.MustGet("employeeID").(int) {
				code = models.InsertContract(&contract)
			} else {
				code = msg.FAIL
			}
		} else {
			contract.EmployeeID = c.MustGet("employeeID").(int)
			contract.CreateDate.Time = time.Now()

			code = models.InsertContract(&contract)
		}
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
	err = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.ID, nil)

	if contractBak.Status == magic.CONTRACT_STATUS_NOT_APPROVAL {
		contractBak.IsPass = contract.IsPass
		var maps = make(map[string]interface{})
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["audit_date"] = time.Now()
		code = models.ApproveContract(&contractBak, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func EditContractEstimatedDeliveryDate(c *gin.Context) {
	var contract models.Contract
	_ = c.ShouldBindJSON(&contract)

	code = models.GeneralUpdate(&models.Contract{}, contract.ID, map[string]interface{}{"estimated_delivery_date": contract.EstimatedDeliveryDate})

	msg.Message(c, code, nil)
}

// 预存款合同完成
func FinalContract(c *gin.Context) {
	var contract models.Contract
	_ = c.ShouldBindJSON(&contract)

	code = models.FinalContract(&contract)

	msg.Message(c, code, nil)
}

// 重置合同状态
func ResetContractCollectionStatus(c *gin.Context) {
	var contract, contractBak models.Contract
	_ = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.ID, nil)

	if contractBak.ID != 0 && contractBak.Status == magic.CONTRACT_STATUS_NOT_FINISH {
		code = models.GeneralUpdate(&models.Contract{}, contract.ID, map[string]interface{}{"collection_status": magic.CONTRATCT_COLLECTION_STATUS_ING})
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

// 驳回已通过的合同
func RejectContract(c *gin.Context) {
	var contract, contractBak models.Contract
	_ = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.ID, nil)

	if contractBak.Status == magic.CONTRACT_STATUS_NOT_FINISH || contract.Status == magic.CONTRACT_STATUS_FINISH {
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
