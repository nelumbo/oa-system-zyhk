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

func AddBidbond(c *gin.Context) {
	var bidbond models.Bidbond
	_ = c.ShouldBindJSON(&bidbond)

	bidbond.EmployeeID = c.MustGet("employeeID").(int)
	bidbond.CreateDate.Time = time.Now()
	bidbond.Status = magic.BIDBOND_STATUS_NOT_APPROVAL

	code = models.GeneralInsert(&bidbond)
	msg.Message(c, code, nil)
}

func DelBidbond(c *gin.Context) {
	var bidbondBak models.Bidbond
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.GeneralSelect(&bidbondBak, id, nil)
		if code == msg.SUCCESS && bidbondBak.EmployeeID == c.MustGet("employeeID").(int) &&
			(bidbondBak.Status == magic.BIDBOND_STATUS_FAIL || bidbondBak.Status == magic.BIDBOND_STATUS_NOT_APPROVAL) {
			code = models.GeneralDelete(&models.Bidbond{}, id)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditBidbond(c *gin.Context) {
	var bidbond, bidbondBak models.Bidbond
	_ = c.ShouldBindJSON(&bidbond)
	code = models.GeneralSelect(&bidbondBak, bidbond.ID, nil)

	if code == msg.SUCCESS && bidbondBak.Status == magic.BIDBOND_STATUS_NOT_APPROVAL &&
		bidbondBak.EmployeeID == c.MustGet("employeeID").(int) {
		var maps = make(map[string]interface{})
		maps["create_date"] = time.Now()
		maps["money"] = bidbond.Money
		maps["create_remark"] = bidbond.CreateRemark
		code = models.GeneralUpdate(&models.Bidbond{}, bidbond.ID, maps)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func ApproveBidbond(c *gin.Context) {
	var bidbond, bidbondBak models.Bidbond
	_ = c.ShouldBindJSON(&bidbond)
	code = models.GeneralSelect(&bidbondBak, bidbond.ID, nil)

	if code == msg.SUCCESS && bidbond.Status == bidbondBak.Status {
		var maps = make(map[string]interface{})
		if bidbondBak.Status == magic.BIDBOND_STATUS_NOT_APPROVAL {
			maps["audit_date"] = time.Now()
			maps["auditor_id"] = c.MustGet("employeeID").(int)
			if bidbond.IsPass {
				maps["status"] = magic.BIDBOND_STATUS_NOT_FINAL
			} else {
				maps["status"] = magic.BIDBOND_STATUS_FAIL
			}
			code = models.GeneralUpdate(&models.Bidbond{}, bidbond.ID, maps)
		} else if bidbondBak.Status == magic.BIDBOND_STATUS_NOT_FINAL {
			maps["final_date"] = time.Now()
			maps["finance_id"] = c.MustGet("employeeID").(int)
			maps["final_remark"] = bidbond.FinalRemark
			if bidbond.IsPass {
				maps["status"] = magic.BIDBOND_STATUS_FINAL
			} else {
				maps["status"] = magic.BIDBOND_STATUS_FAIL
			}
			code = models.GeneralUpdate(&models.Bidbond{}, bidbond.ID, maps)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func QueryBidbonds(c *gin.Context) {
	var bidbondQuery models.Bidbond
	_ = c.ShouldBindJSON(&bidbondQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectBidbonds(&bidbondQuery, &xForms)
	msg.Message(c, code, xForms)
}
