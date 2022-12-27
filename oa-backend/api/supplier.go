package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddSupplier(c *gin.Context) {
	var supplier models.Supplier
	_ = c.ShouldBindJSON(&supplier)

	code = models.GeneralInsert(&supplier)
	msg.Message(c, code, nil)
}

func EditSupplier(c *gin.Context) {
	var supplier models.Supplier
	_ = c.ShouldBindJSON(&supplier)
	var maps = make(map[string]interface{})
	maps["name"] = supplier.Name
	maps["web"] = supplier.Web
	maps["address"] = supplier.Address
	maps["linkman"] = supplier.Linkman
	maps["phone"] = supplier.Phone
	maps["wechat_id"] = supplier.WechatID
	maps["email"] = supplier.Email
	maps["description"] = supplier.Description
	maps["remark"] = supplier.Remark

	code = models.GeneralUpdate(&models.Supplier{}, supplier.ID, maps)
	msg.Message(c, code, nil)
}

func QuerySuppliers(c *gin.Context) {
	var supplierQuery models.Supplier
	_ = c.ShouldBindJSON(&supplierQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectSuppliers(&supplierQuery, &xForms)
	msg.Message(c, code, xForms)
}
