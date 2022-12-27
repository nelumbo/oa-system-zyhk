package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCustomerCompany(c *gin.Context) {
	var customerCompany models.CustomerCompany
	_ = c.ShouldBindJSON(&customerCompany)

	code = models.GeneralInsert(&customerCompany)
	msg.Message(c, code, nil)
}

func DelCustomerCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.DeleteCustomerCompany(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditCustomerCompany(c *gin.Context) {
	var customerCompany models.CustomerCompany
	_ = c.ShouldBindJSON(&customerCompany)
	var maps = make(map[string]interface{})
	maps["region_id"] = customerCompany.RegionID
	maps["name"] = customerCompany.Name
	maps["address"] = customerCompany.Address

	code = models.GeneralUpdate(&models.CustomerCompany{}, customerCompany.ID, maps)
	msg.Message(c, code, nil)
}

func QueryCustomerCompanys(c *gin.Context) {
	var customerCompanyQuery models.CustomerCompany
	_ = c.ShouldBindJSON(&customerCompanyQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectCustomerCompanys(&customerCompanyQuery, &xForms)
	msg.Message(c, code, xForms)
}
