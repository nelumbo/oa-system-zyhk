package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddProductType(c *gin.Context) {
	var productType models.ProductType
	_ = c.ShouldBindJSON(&productType)

	code = models.GeneralInsert(&productType)
	msg.Message(c, code, nil)
}

func EditProductType(c *gin.Context) {
	var productType models.ProductType
	_ = c.ShouldBindJSON(&productType)
	var maps = make(map[string]interface{})
	maps["name"] = productType.Name
	maps["push_money_percentages"] = productType.PushMoneyPercentages
	maps["min_push_money_percentages"] = productType.MinPushMoneyPercentages
	maps["push_money_percentages_up"] = productType.PushMoneyPercentagesUp
	maps["push_money_percentages_down"] = productType.PushMoneyPercentagesDown
	maps["business_money_percentages"] = productType.BusinessMoneyPercentages
	maps["business_money_percentages_up"] = productType.BusinessMoneyPercentagesUp
	maps["is_task_load"] = productType.IsTaskLoad

	code = models.GeneralUpdate(&productType, productType.ID, maps)
	msg.Message(c, code, nil)
}

func QueryProductTypes(c *gin.Context) {
	var productTypeQuery models.ProductType
	_ = c.ShouldBindJSON(&productTypeQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectProductTypes(&productTypeQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllProductType(c *gin.Context) {
	var productTypes []models.ProductType
	code = models.GeneralSelectAll(&productTypes, nil)
	msg.Message(c, code, productTypes)
}
