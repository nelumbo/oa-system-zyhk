package api

import (
	"oa-backend/models"
	ginUtil "oa-backend/utils/gin"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
)

func SubmitPredesignTask(c *gin.Context) {
	var predesignTask, predesignTaskbak *models.PredesignTask
	_ = c.ShouldBindJSON(&predesignTask)
	code = models.GeneralSelect(&predesignTaskbak, predesignTask.ID, nil)

	if code == msg.SUCCESS && predesignTaskbak.Status == magic.PREDESIGN_TASK_STATUS_NOT_SUBMIT &&
		predesignTaskbak.EmployeeID == c.MustGet("employeeID").(int) {
		var maps = make(map[string]interface{})
		maps["final_remark"] = predesignTask.FinalRemark
		maps["final_date"] = time.Now()
		maps["status"] = magic.PREDESIGN_TASK_STATUS_NOT_APPROVAL
		code = models.GeneralUpdate(&models.PredesignTask{}, predesignTask.ID, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func ApprovePredesignTask(c *gin.Context) {
	var predesignTask, predesignTaskbak models.PredesignTask
	_ = c.ShouldBindJSON(&predesignTask)
	code = models.GeneralSelect(&predesignTaskbak, predesignTask.ID, nil)

	if code == msg.SUCCESS && predesignTaskbak.Status == magic.PREDESIGN_TASK_STATUS_NOT_APPROVAL {
		var maps = make(map[string]interface{})
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["approve_date"] = predesignTask.ApproveDate

		if predesignTask.IsPass {
			maps["status"] = magic.PREDESIGN_TASK_STATUS_FINAL
			code = models.GeneralUpdate(&models.PredesignTask{}, predesignTask.ID, maps)
		} else {
			maps["status"] = magic.PREDESIGN_TASK_STATUS_FAIL
			predesignTask.PredesignID = predesignTaskbak.PredesignID
			predesignTask.AuditorID = c.MustGet("employeeID").(int)
			code = models.InsertPredesignTask(&predesignTask, maps)
		}
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func QueryPredesignTasks(c *gin.Context) {
	var predesignTaskQuery models.PredesignTask
	_ = c.ShouldBindJSON(&predesignTaskQuery)

	xForms := ginUtil.GinArrayPreprocessing(c)

	xForms.Data, code = models.SelectPredesignTasks(&predesignTaskQuery, &xForms)
	msg.Message(c, code, xForms)
}
