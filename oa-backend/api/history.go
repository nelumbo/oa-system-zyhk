package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func QueryHistoryEmployees(c *gin.Context) {
	var historyEmployee models.HistoryEmployee
	_ = c.ShouldBindJSON(&historyEmployee)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectHistoryEmployees(&historyEmployee, &xForms)
	msg.Message(c, code, xForms)
}

func QueryHistoryOffices(c *gin.Context) {
	var historyOffice models.HistoryOffice
	_ = c.ShouldBindJSON(&historyOffice)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectHistoryOffices(&historyOffice, &xForms)
	msg.Message(c, code, xForms)
}

func QueryHistoryProducts(c *gin.Context) {
	var historyProduct models.HistoryProduct
	_ = c.ShouldBindJSON(&historyProduct)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectHistoryProducts(&historyProduct, &xForms)
	msg.Message(c, code, xForms)
}
