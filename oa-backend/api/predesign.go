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

func AddPredesign(c *gin.Context) {
	var predesign models.Predesign
	_ = c.ShouldBindJSON(&predesign)

	predesign.CreateDate.Time = time.Now()
	predesign.EmployeeID = c.MustGet("employeeID").(int)
	predesign.Status = magic.PREDESIGN_STATUS_NOT_APPROVAL

	code = models.GeneralInsert(&predesign)
	msg.Message(c, code, nil)
}

func DelPredesign(c *gin.Context) {
	var predesignBak models.Predesign
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.GeneralSelect(&predesignBak, id, nil)
		if code == msg.SUCCESS && predesignBak.EmployeeID == c.MustGet("employeeID").(int) &&
			(predesignBak.Status == magic.PREDESIGN_STATUS_FAIL || predesignBak.Status == magic.PREDESIGN_STATUS_NOT_APPROVAL) {
			code = models.GeneralDelete(&models.Predesign{}, id)
		} else {
			code = msg.FAIL
		}
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditPredesign(c *gin.Context) {
	var predesign, predesignBak models.Predesign
	_ = c.ShouldBindJSON(&predesign)
	code = models.GeneralSelect(&predesignBak, predesign.ID, nil)

	if code == msg.SUCCESS && predesignBak.Status == magic.PREDESIGN_STATUS_NOT_APPROVAL &&
		predesignBak.EmployeeID == c.MustGet("employeeID").(int) {
		var maps = make(map[string]interface{})
		maps["create_date"] = time.Now()
		maps["create_remark"] = predesign.CreateRemark
		code = models.GeneralUpdate(&models.Predesign{}, predesign.ID, maps)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func ApprovePredesign(c *gin.Context) {
	var predesign, predesignBak models.Predesign
	_ = c.ShouldBindJSON(&predesign)
	code = models.GeneralSelect(&predesignBak, predesign.ID, nil)

	if code == msg.SUCCESS && predesign.Status == predesignBak.Status &&
		predesign.Status == magic.PREDESIGN_STATUS_NOT_APPROVAL {
		var maps = make(map[string]interface{})
		maps["audit_date"] = time.Now()
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["approve_remark"] = predesign.ApproveRemark
		if predesign.IsPass {
			predesign.AuditorID = c.MustGet("employeeID").(int)
			maps["status"] = magic.PREDESIGN_STATUS_NOT_FINAL
			code = models.UpdatePredesign(&predesign, maps)
		} else {
			maps["status"] = magic.PREDESIGN_STATUS_FAIL
			code = models.GeneralUpdate(&models.Predesign{}, predesign.ID, maps)
		}
	} else {
		code = msg.FAIL
	}
	msg.Message(c, code, nil)
}

func QueryPredesign(c *gin.Context) {
	var predesign models.Predesign
	id, err := strconv.Atoi(c.Param("id"))

	if err == nil {
		predesign, code = models.SelectPredesign(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, predesign)
}

func QueryPredesigns(c *gin.Context) {
	var predesignQuery models.Predesign
	_ = c.ShouldBindJSON(&predesignQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectPredesigns(&predesignQuery, &xForms)
	msg.Message(c, code, xForms)
}
