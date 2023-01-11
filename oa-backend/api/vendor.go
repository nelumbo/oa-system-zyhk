package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddVendor(c *gin.Context) {
	var vendor models.Vendor
	_ = c.ShouldBindJSON(&vendor)

	vendorCre := models.Vendor{
		ID:       vendor.ID,
		IsDelete: false,
		Name:     vendor.Name,
	}

	code = models.GeneralInsert(&vendorCre)
	msg.Message(c, code, nil)
}

func EditVendor(c *gin.Context) {
	var vendor models.Vendor
	_ = c.ShouldBindJSON(&vendor)
	var maps = make(map[string]interface{})
	maps["name"] = vendor.Name

	code = models.GeneralUpdate(&models.Vendor{}, vendor.ID, maps)
	msg.Message(c, code, nil)
}

func QueryVendors(c *gin.Context) {
	var vendorQuery models.Vendor
	_ = c.ShouldBindJSON(&vendorQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectVendors(&vendorQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllVendor(c *gin.Context) {
	var vendors []models.Vendor
	code = models.GeneralSelectAll(&vendors, nil)
	msg.Message(c, code, vendors)
}
