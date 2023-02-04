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

func AddProductTrial(c *gin.Context) {
	var productTrial models.ProductTrial
	_ = c.ShouldBindJSON(&productTrial)

	productTrialCre := models.ProductTrial{
		ID:         productTrial.ID,
		IsDelete:   false,
		ProductID:  productTrial.ProductID,
		Number:     productTrial.Number,
		Status:     magic.PRODUCT_TRIAL_STATUS_NO_APPROVE,
		EmployeeID: c.MustGet("employeeID").(int),
		CreateDate: models.XDate{Time: time.Now()},
		Remark1:    productTrial.Remark1,
	}

	code = models.GeneralInsert(&productTrialCre)
	msg.Message(c, code, nil)
}

func DelProductTrial(c *gin.Context) {
	var productTrialBak models.ProductTrial
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.GeneralSelect(&productTrialBak, id, nil)
		if code == msg.SUCCESS && productTrialBak.EmployeeID == c.MustGet("employeeID").(int) &&
			(productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_REJECT || productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_APPROVE) {
			code = models.GeneralDelete(&models.ProductTrial{}, id)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditProductTrial(c *gin.Context) {
	var productTrial, productTrialBak models.ProductTrial
	_ = c.ShouldBindJSON(&productTrial)
	code = models.GeneralSelect(&productTrialBak, productTrial.ID, nil)

	if code == msg.SUCCESS && productTrialBak.EmployeeID == c.MustGet("employeeID").(int) &&
		(productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_APPROVE ||
			productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_REJECT) {
		var maps = make(map[string]interface{})
		maps["product_id"] = productTrial.ProductID
		maps["number"] = productTrial.Number
		code = models.GeneralUpdate(&models.ProductTrial{}, productTrial.ID, maps)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func NextProductTrial(c *gin.Context) {
	var productTrial, productTrialBak models.ProductTrial
	_ = c.ShouldBindJSON(&productTrial)
	code = models.GeneralSelect(&productTrialBak, productTrial.ID, nil)
	if code == msg.SUCCESS {
		code = models.NextProductTrial(&productTrial, &productTrialBak, c.MustGet("employeeID").(int))
	} else {
		code = msg.ERROR
	}

	msg.Message(c, code, nil)
}

func QueryProductTrials(c *gin.Context) {
	var productTrialQuery models.ProductTrial
	_ = c.ShouldBindJSON(&productTrialQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectProductTrials(&productTrialQuery, &xForms)
	msg.Message(c, code, xForms)
}
