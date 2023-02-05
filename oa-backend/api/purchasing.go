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

func SavePurchasing(c *gin.Context) {
	var purchasing, purchasingNew models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	purchasingNew.ContractID = purchasing.ContractID
	purchasingNew.TaskID = purchasing.TaskID
	purchasingNew.EmployeeID = c.MustGet("employeeID").(int)
	purchasingNew.Type = magic.PURCHASING_TYPE_CONTRACT
	purchasingNew.ProductID = purchasing.ProductID
	purchasingNew.Number = purchasing.Number
	code = models.GeneralInsert(&purchasingNew)
	msg.Message(c, code, nil)
}

func SubmitPurchasing(c *gin.Context) {
	var purchasing models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	purchasing.EmployeeID = c.MustGet("employeeID").(int)
	var maps = make(map[string]interface{})
	maps["status"] = magic.PURCHASING_STATUS_NO_CHECK
	maps["create_date"] = time.Now()

	code = models.SubmitPurchasing(&purchasing, maps)
	msg.Message(c, code, nil)
}

func AddPurchasing(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	var productBak models.Product
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&productBak, purchasing.ProductID, nil)

	purchasingBak.EmployeeID = c.MustGet("employeeID").(int)
	if purchasing.Type == magic.PURCHASING_TYPE_TASK {
		purchasingBak.No = productBak.Name + " " + time.Now().Format("2006-01-02")
		purchasingBak.Type = magic.PURCHASING_TYPE_TASK
	} else if purchasing.Type == magic.PURCHASING_TYPE_BANK {
		purchasingBak.No = "库存下单"
		purchasingBak.Type = magic.PURCHASING_TYPE_TASK
	}
	purchasingBak.ProductID = productBak.ID
	purchasingBak.Price = purchasing.Price
	purchasingBak.Number = purchasing.RealNumber
	purchasingBak.RealNumber = purchasing.RealNumber
	purchasingBak.TotalPrice = float64(purchasing.RealNumber) * purchasing.Price
	purchasingBak.Status = magic.PURCHASING_STATUS_NO_APPROVE
	purchasingBak.PurchaseManID = c.MustGet("employeeID").(int)

	purchasingBak.CreateDate.Time = time.Now()
	purchasingBak.EndDate.Time = purchasingBak.CreateDate.Time

	code = models.GeneralInsert(&purchasingBak)
	msg.Message(c, code, nil)
}

func DelPurchasing(c *gin.Context) {
	var purchasingBak models.Purchasing
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.GeneralSelect(&purchasingBak, id, nil)
		if code == msg.SUCCESS && purchasingBak.EmployeeID == c.MustGet("employeeID").(int) &&
			purchasingBak.Status == magic.PURCHASING_STATUS_SAVE {
			code = models.DeletePurchasing(&purchasingBak)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func ApprovePurchasing(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)

	if purchasingBak.Status == magic.PURCHASING_STATUS_NO_CHECK {
		var maps = make(map[string]interface{})
		maps["status"] = magic.PURCHASING_STATUS_NO_APPROVE
		maps["real_number"] = purchasing.RealNumber
		maps["price"] = purchasing.Price
		maps["total_price"] = float64(purchasing.RealNumber) * purchasing.Price
		maps["purchase_man_id"] = c.MustGet("employeeID").(int)
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
	} else if purchasingBak.Status == magic.PURCHASING_STATUS_NO_APPROVE {
		purchasingBak.IsPass = purchasing.IsPass
		var maps = make(map[string]interface{})
		maps["status"] = magic.PURCHASING_STATUS_NO_FINAL
		maps["product_status"] = magic.PURCHASING_PRODUCT_STATUS_NO_FINAL
		maps["pay_status"] = magic.PURCHASING_PAY_STATUS_NO_FINAL
		maps["invoice_status"] = magic.PURCHASING_INVOICE_STATUS_NO_FINAL
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["audit_date"] = time.Now()
		code = models.ApprovePurchasing(&purchasingBak, maps)
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func FinalPurchasingProductStatus(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)

	if purchasingBak.Status == magic.PURCHASING_STATUS_NO_FINAL &&
		purchasingBak.ProductStatus == magic.PURCHASING_PRODUCT_STATUS_NO_FINAL {
		var maps = make(map[string]interface{})
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["real_end_date"] = time.Now()
		maps["product_status"] = magic.PURCHASING_PRODUCT_STATUS_FINAL
		maps["product_remark"] = purchasing.ProductRemark
		code = models.UpdatePurchasingProductStatus(&purchasingBak, maps, c.MustGet("employeeID").(int))
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func FinalPurchasingPayStatus(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)

	if purchasingBak.Status == magic.PURCHASING_STATUS_NO_FINAL &&
		purchasingBak.PayStatus == magic.PURCHASING_PAY_STATUS_NO_FINAL {
		var maps = make(map[string]interface{})
		maps["pay_man_id"] = c.MustGet("employeeID").(int)
		maps["pay_date"] = time.Now()
		maps["pay_create_date"] = time.Now()
		maps["pay_status"] = magic.PURCHASING_PRODUCT_STATUS_FINAL
		maps["pay_remark"] = purchasing.PayRemark
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func FinalPurchasingInvoiceStatus(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)

	if purchasingBak.Status == magic.PURCHASING_STATUS_NO_FINAL &&
		purchasingBak.InvoiceStatus == magic.PURCHASING_INVOICE_STATUS_NO_FINAL {
		var maps = make(map[string]interface{})
		maps["invoice_man_id"] = c.MustGet("employeeID").(int)
		maps["invoice_date"] = time.Now()
		maps["invoice_status"] = magic.PURCHASING_INVOICE_STATUS_FINAL
		maps["invoice_remark"] = purchasing.InvoiceRemark
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func FinalPurchasing(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)

	if purchasingBak.Status == magic.PURCHASING_STATUS_NO_FINAL &&
		purchasingBak.ProductStatus == magic.PURCHASING_PRODUCT_STATUS_FINAL &&
		purchasingBak.PayStatus == magic.PURCHASING_PAY_STATUS_FINAL &&
		purchasingBak.InvoiceStatus == magic.PURCHASING_INVOICE_STATUS_FINAL {
		var maps = make(map[string]interface{})
		maps["status"] = magic.PURCHASING_STATUS_FINAL
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func QueryPurchasings(c *gin.Context) {
	var purchasingQuery models.Purchasing
	_ = c.ShouldBindJSON(&purchasingQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectPurchasings(&purchasingQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllPurchasing(c *gin.Context) {
	var purchasing models.Purchasing
	var purchasings []models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)
	var maps = make(map[string]interface{})
	if purchasing.ContractID != 0 {
		maps["contract_id"] = purchasing.ContractID
	}
	if purchasing.TaskID != 0 {
		maps["task_id"] = purchasing.TaskID
	}

	purchasings, code = models.SelectAllPurchasings(maps)

	msg.Message(c, code, purchasings)
}
