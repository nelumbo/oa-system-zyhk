package api

import (
	"oa-backend/models"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	var payment models.Payment
	_ = c.ShouldBindJSON(&payment)

	payment.CreateDate.Time = time.Now()
	payment.EmployeeID = c.MustGet("employeeID").(int)

	code = models.InsertPayment(&payment)

	msg.Message(c, code, nil)
}

func EditPayment(c *gin.Context) {
	var payment models.Payment
	_ = c.ShouldBindJSON(&payment)

	payment.EmployeeID = c.MustGet("employeeID").(int)

	code = models.UpdatePayment(&payment)
	msg.Message(c, code, nil)
}

func FinalPayment(c *gin.Context) {
	var payment models.Payment
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&payment)

	_ = models.GeneralSelect(&contractBak, payment.ContractID, nil)

	if contractBak.ID != 0 {
		var maps = make(map[string]interface{})
		maps["collection_status"] = magic.CONTRATCT_COLLECTION_STATUS_FINISH
		// if !contractBak.IsPreDeposit && contractBak.ProductionStatus == magic.CONTRATCT_PRODUCTION_STATUS_FINISH {
		// 	maps["status"] = magic.CONTRACT_STATUS_FINISH
		// }
		code = models.GeneralUpdate(&models.Contract{}, payment.ContractID, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}
