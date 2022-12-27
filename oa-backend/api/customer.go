package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCustomer(c *gin.Context) {
	var customer models.Customer
	_ = c.ShouldBindJSON(&customer)

	code = models.GeneralInsert(&customer)
	msg.Message(c, code, nil)
}

func DelCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.DeleteCustomer(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditCustomer(c *gin.Context) {
	var customer models.Customer
	_ = c.ShouldBindJSON(&customer)
	var maps = make(map[string]interface{})
	maps["customer_company_id"] = customer.CustomerCompanyID
	maps["name"] = customer.Name
	maps["research_group"] = customer.ResearchGroup
	maps["phone"] = customer.Phone
	maps["wechat_id"] = customer.WechatID
	maps["email"] = customer.Email

	code = models.GeneralUpdate(&models.Customer{}, customer.ID, maps)
	msg.Message(c, code, nil)
}

func QueryCustomers(c *gin.Context) {
	var customerQuery models.Customer
	_ = c.ShouldBindJSON(&customerQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectCustomers(&customerQuery, &xForms)
	msg.Message(c, code, xForms)
}
