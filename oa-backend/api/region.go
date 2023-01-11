package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddRegion(c *gin.Context) {
	var region models.Region
	_ = c.ShouldBindJSON(&region)

	regionCre := models.Region{
		ID:       region.ID,
		IsDelete: false,
		Name:     region.Name,
	}

	code = models.GeneralInsert(&regionCre)
	msg.Message(c, code, nil)
}

func EditRegion(c *gin.Context) {
	var region models.Region
	_ = c.ShouldBindJSON(&region)
	var maps = make(map[string]interface{})
	maps["name"] = region.Name

	code = models.GeneralUpdate(&models.Region{}, region.ID, maps)
	msg.Message(c, code, nil)
}

func QueryRegions(c *gin.Context) {
	var regionQuery models.Region
	_ = c.ShouldBindJSON(&regionQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectRegions(&regionQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllRegion(c *gin.Context) {
	var regions []models.Region
	code = models.GeneralSelectAll(&regions, nil)
	msg.Message(c, code, regions)
}
