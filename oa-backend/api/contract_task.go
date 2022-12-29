package api

import (
	"oa-backend/models"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"

	"github.com/gin-gonic/gin"
)

// 预存款合同添加任务
func AddTask(c *gin.Context) {
	var task models.Task
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&task)

	code = models.GeneralSelect(contractBak, task.ContractID, nil)

	if code == msg.SUCCESS && contractBak.IsPreDeposit && contractBak.PreDeposit >= task.TotalPrice {
		code = models.InsertTask(&contractBak, &task)
	}

	msg.Message(c, code, nil)
}

func ApproveTask(c *gin.Context) {
	var task, taskBak models.Task
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&task)

	_ = models.GeneralSelect(taskBak, task.ID, nil)
	_ = models.GeneralSelect(contractBak, task.ContractID, nil)

	if taskBak.ID != 0 && contractBak.ID != 0 && taskBak.Status != magic.TASK_STATUS_REJECT &&
		contractBak.Status == magic.CONTATCT_PRODUCTION_STATUS_ING {
		code = models.ApproveTask(&task, c.MustGet("employeeID").(int))
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}
