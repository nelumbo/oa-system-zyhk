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

		contractCre := models.Contract{
			ID:                    contract.ID,
			IsDelete:              false,
			RegionID:              contract.RegionID,
			EmployeeID:            c.MustGet("employeeID").(int),
			CustomerID:            contract.CustomerID,
			ContractDate:          contract.ContractDate,
			VendorID:              contract.VendorID,
			EstimatedDeliveryDate: contract.EstimatedDeliveryDate,
			InvoiceType:           contract.InvoiceType,
			InvoiceContent:        contract.InvoiceContent,
			PaymentContent:        contract.PaymentContent,
			IsSpecial:             contract.IsSpecial,
			IsPreDeposit:          contract.IsPreDeposit,
			PreDeposit:            0,
			PreDepositRecord:      contract.PreDepositRecord,
			PayType:               contract.PayType,
			TotalAmount:           0,
			PaymentTotalAmount:    0,
			Remark:                contract.Remark,
			ProductionStatus:      0,
			CollectionStatus:      0,
			Status:                magic.CONTRACT_STATUS_SAVE,
			CreateDate:            models.XDate{Time: time.Now()},
		}

		code = models.GeneralInsert(&contractCre)
	} else {
		var contractBak models.Contract
		_ = models.GeneralSelect(&contractBak, contract.ID, nil)
		if contractBak.ID == contract.ID && contractBak.EmployeeID == c.MustGet("employeeID").(int) {
			var maps = make(map[string]interface{})
			maps["region_id"] = contract.RegionID
			maps["customer_id"] = contract.CustomerID
			maps["contract_date"] = contract.ContractDate
			maps["vendor_id"] = contract.VendorID
			maps["estimated_delivery_date"] = contract.ContractDate
			maps["invoice_type"] = contract.InvoiceType
			maps["invoice_content"] = contract.InvoiceContent
			maps["payment_content"] = contract.PaymentContent
			maps["is_special"] = contract.IsSpecial
			maps["is_pre_deposit"] = contract.IsPreDeposit
			maps["pre_deposit_record"] = contract.PreDepositRecord
			maps["pay_type"] = contract.PayType
			maps["remark"] = contract.Remark
			code = models.GeneralUpdate(&models.Contract{}, contractBak.ID, maps)
		} else {
			code = msg.FAIL
		}
	}

	msg.Message(c, code, nil)
}

func AddContract(c *gin.Context) {
	var contract models.Contract
	var employeeBak models.Employee
	_ = c.ShouldBindJSON(&contract)

	code = models.GeneralSelect(&employeeBak, c.MustGet("employeeID").(int), nil)

	if code == msg.SUCCESS {

		if contract.IsPreDeposit {
			contract.Tasks = nil
		} else {
			for i := range contract.Tasks {
				var product models.Product
				_ = models.GeneralSelect(&product, contract.Tasks[i].ProductID, nil)
				contract.TotalAmount += contract.Tasks[i].TotalPrice
				contract.Tasks[i].ProductAttributeID = product.AttributeID
				contract.Tasks[i].Product = models.Product{}
			}
		}

		contractCre := models.Contract{
			ID:                    contract.ID,
			IsDelete:              false,
			RegionID:              contract.RegionID,
			OfficeID:              employeeBak.OfficeID,
			EmployeeID:            c.MustGet("employeeID").(int),
			CustomerID:            contract.CustomerID,
			ContractDate:          contract.ContractDate,
			VendorID:              contract.VendorID,
			EstimatedDeliveryDate: contract.EstimatedDeliveryDate,
			InvoiceType:           contract.InvoiceType,
			InvoiceContent:        contract.InvoiceContent,
			PaymentContent:        contract.PaymentContent,
			IsSpecial:             contract.IsSpecial,
			IsPreDeposit:          contract.IsPreDeposit,
			PreDeposit:            0,
			PreDepositRecord:      contract.PreDepositRecord,
			PayType:               contract.PayType,
			TotalAmount:           contract.TotalAmount,
			PaymentTotalAmount:    0,
			Remark:                contract.Remark,
			ProductionStatus:      0,
			CollectionStatus:      0,
			Status:                magic.CONTRACT_STATUS_NOT_APPROVAL,
			CreateDate:            models.XDate{Time: time.Now()},
			Tasks:                 contract.Tasks,
		}

		code = models.InsertContract(&contractCre)
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

	_ = models.GeneralSelect(&contractBak, contract.ID, nil)

	if contractBak.Status == magic.CONTRACT_STATUS_NOT_APPROVAL {
		contractBak.IsPass = contract.IsPass
		contractBak.Tasks = contract.Tasks
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

// 合同完成
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

func EditContractPreDeposit(c *gin.Context) {
	var contract, contractBak models.Contract
	_ = c.ShouldBindJSON(&contract)

	_ = models.GeneralSelect(&contractBak, contract.ID, nil)

	if contractBak.ID != 0 && contractBak.IsPreDeposit && contractBak.Status == 2 {
		code = models.UpdateContractPreDeposit(&contract, c.MustGet("employeeID").(int))
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}
