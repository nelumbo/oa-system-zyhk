package api

import (
	"oa-backend/models"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
)

func SavePurchasing(c *gin.Context) {
	var purchasing, purchasingNew models.Purchasing
	_ = c.ShouldBindJSON(&purchasing)

	purchasingNew.ContractID = purchasing.ContractID
	purchasingNew.TaskID = purchasing.TaskID
	purchasingNew.EmployeeID = c.MustGet("employeeID").(int)
	purchasingNew.Type = purchasing.Type
	purchasingNew.ProductID = purchasing.ProductID
	purchasingNew.Number = purchasing.Number
	code = models.GeneralInsert(&purchasingNew)
	msg.Message(c, code, nil)
}

func AddPurchasing(c *gin.Context) {
	var purchasing, purchasingBak models.Purchasing
	var productBak models.Product
	_ = c.ShouldBindJSON(&purchasing)

	_ = models.GeneralSelect(&purchasingBak, purchasing.ID, nil)
	_ = models.GeneralSelect(&productBak, purchasing.ProductID, nil)

	var maps = make(map[string]interface{})
	maps["create_date"] = time.Now()
	maps["product_attribute_id"] = productBak.AttributeID
	if purchasing.Type == magic.PURCHASING_TYPE_CONTRACT {
		maps["status"] = magic.PURCHASING_STATUS_NO_CHECK
		maps["price"] = purchasing.Price
		maps["real_number"] = purchasing.RealNumber
		maps["total_price"] = float64(purchasing.RealNumber) * purchasing.Price
		code = models.GeneralInsert(&purchasing)
	} else if purchasing.Type == magic.PURCHASING_TYPE_TASK || purchasing.Type == magic.PURCHASING_TYPE_BANK {
		maps["status"] = magic.PURCHASING_STATUS_NO_APPROVE
		maps["purchase_man_id"] = c.MustGet("employeeID").(int)
		maps["total_price"] = float64(purchasingBak.RealNumber) * purchasingBak.Price
		code = models.GeneralInsert(&purchasing)
	} else {
		code = msg.FAIL
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
		maps["realNumber"] = purchasing.RealNumber
		maps["price"] = purchasing.Price
		maps["total_price"] = float64(purchasing.RealNumber) * purchasing.Price
		maps["purchase_man_id"] = c.MustGet("employeeID").(int)
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
	} else if purchasingBak.Status == magic.PURCHASING_STATUS_NO_APPROVE {
		purchasingBak.IsPass = purchasing.IsPass
		var maps = make(map[string]interface{})
		maps["status"] = magic.PURCHASING_STATUS_NO_FINAL
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
		code = models.GeneralUpdate(&models.Purchasing{}, purchasingBak.ID, maps)
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
		maps["invocie_man_id"] = c.MustGet("employeeID").(int)
		maps["invoice_date"] = time.Now()
		maps["invoice_status"] = magic.PURCHASING_INVOICE_STATUS_FINAL
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
