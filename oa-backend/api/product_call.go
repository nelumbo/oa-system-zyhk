package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func QueryProductCalls(c *gin.Context) {
	var productCallQuery models.ProductCall
	// _ = c.ShouldBindJSON(&productCallQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectProductCalls(&productCallQuery, &xForms)
	msg.Message(c, code, xForms)
}
