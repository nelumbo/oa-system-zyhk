package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddOffice(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)

	code = models.GeneralInsert(&office)
	msg.Message(c, code, nil)
}

func DelOffice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		code = models.GeneralDelete(&models.Office{}, id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditOffice(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)
	var maps = make(map[string]interface{})
	maps["name"] = office.Name
	maps["number"] = office.Number
	maps["business_money"] = office.BusinessMoney
	maps["money"] = office.Money
	maps["money_cold"] = office.MoneyCold
	maps["task_load"] = office.TaskLoad
	maps["target_load"] = office.TargetLoad
	if office.RoleID == 0 {
		maps["role_id"] = nil
	} else {
		maps["role_id"] = office.RoleID
	}
	maps["is_ranking"] = office.IsRanking

	code = models.GeneralUpdate(&models.Office{}, office.ID, maps)
	msg.Message(c, code, nil)
}

func QueryOffice(c *gin.Context) {
	var office models.Office
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		office, code = models.SelectOffice(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, office)
}

func QueryOffices(c *gin.Context) {
	var officeQuery models.Office
	_ = c.ShouldBindJSON(&officeQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectOffices(&officeQuery, &xForms)
	msg.Message(c, code, xForms)
}

func QueryAllOffice(c *gin.Context) {
	var offices []models.Office
	code = models.GeneralSelectAll(&offices, nil)
	msg.Message(c, code, offices)
}
