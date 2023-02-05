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

	productCre := models.Product{
		ID:            product.ID,
		IsDelete:      false,
		Name:          product.Name,
		Version:       product.Version,
		Brand:         product.Brand,
		Specification: product.Specification,
		Number:        product.Number,
		CallNumber:    product.CallNumber,
		Unit:          product.Unit,
		DeliveryCycle: product.DeliveryCycle,
		Remark:        product.Remark,
		IsFree:        product.IsFree,
		TypeID:        product.TypeID,
		PurchasePrice: product.PurchasePrice,
		Suppliers:     product.Suppliers,

		Attribute: models.ProductAttribute{
			ID:               product.Attribute.ID,
			IsDelete:         false,
			StandardPrice:    product.Attribute.StandardPrice,
			StandardPriceUSD: product.Attribute.StandardPriceUSD,
		},
	}

	code = models.InsertProduct(&productCre)
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

func EditProductNumber(c *gin.Context) {
	var product models.Product
	_ = c.ShouldBindJSON(&product)

	code = models.UpdateProductNumber(&product, c.MustGet("employeeID").(int))
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
