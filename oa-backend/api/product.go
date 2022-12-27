package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product
	_ = c.ShouldBindJSON(&product)

	// code = models.GeneralInsert(&product)
	code = models.InsertProduct(&product)
	msg.Message(c, code, nil)
}

func EditProductBase(c *gin.Context) {
	var product models.Product
	_ = c.ShouldBindJSON(&product)

	code = models.UpdateProduct(&product)
	msg.Message(c, code, nil)
}

func EditProductAttribute(c *gin.Context) {
	var product models.Product
	_ = c.ShouldBindJSON(&product)

	code = models.UpdateProductAttribute(&product)
	msg.Message(c, code, nil)
}

func QueryProduct(c *gin.Context) {
	var product models.Product
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		product, code = models.SelectProduct(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, product)
}

func QueryProducts(c *gin.Context) {
	var productQuery models.Product
	_ = c.ShouldBindJSON(&productQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectProducts(&productQuery, &xForms)
	msg.Message(c, code, xForms)
}
