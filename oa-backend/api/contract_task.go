package api

import (
	"oa-backend/models"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 任务分发
func DistributeTask(c *gin.Context) {
	var task, taskBak models.Task
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&task)

	_ = models.GeneralSelect(&taskBak, task.ID, nil)
	_ = models.GeneralSelect(&contractBak, task.ContractID, nil)

	if taskBak.ID != 0 && contractBak.ID != 0 && taskBak.ContractID == task.ContractID &&
		taskBak.Status == magic.TASK_STATUS_NOT_DISTRIBUTE &&
		contractBak.Status == magic.CONTRACT_STATUS_NOT_FINISH {

		taskBak.Contract = contractBak

		var maps = make(map[string]interface{})
		tn := time.Now()
		maps["type"] = task.Type
		maps["auditor_id"] = c.MustGet("employeeID").(int)
		maps["audit_date"] = tn
		maps["inventory_man_id"] = task.InventoryManID
		maps["shipment_man_id"] = task.ShipmentManID

		if contractBak.IsSpecial {
			maps["push_money_percentages"] = task.PushMoneyPercentages
		}

		switch task.Type {
		case magic.TASK_TYPE_1:
			maps["status"] = magic.TASK_STATUS_NOT_STORAGE

			maps["inventory_start_date"] = tn
		case magic.TASK_TYPE_2:
			maps["status"] = magic.TASK_STATUS_NOT_PURCHASE

			maps["purchase_man_id"] = task.PurchaseManID
			maps["purchase_start_date"] = tn
			maps["purchase_days"] = task.PurchaseDays
			maps["purchase_end_date"] = tn.AddDate(0, 0, task.PurchaseDays)
		case magic.TASK_TYPE_3:
			maps["status"] = magic.TASK_STATUS_NOT_DESIGN

			maps["technician_man_id"] = task.TechnicianManID
			maps["technician_start_date"] = tn
			maps["technician_days"] = task.PurchaseDays
			maps["technician_end_date"] = tn.AddDate(0, 0, task.PurchaseDays)

			maps["purchase_man_id"] = task.PurchaseManID
			maps["purchase_days"] = task.PurchaseDays
		}

		code = models.DistributeTask(&taskBak, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

func NextTask(c *gin.Context) {
	var task, taskBak models.Task
	_ = c.ShouldBindJSON(&task)

	_ = models.GeneralSelect(&taskBak, task.ID, nil)

	if taskBak.ID != 0 && task.Status == taskBak.Status {
		var maps = make(map[string]interface{})
		employeeID := c.MustGet("employeeID").(int)
		tn := time.Now()
		if taskBak.Status == magic.TASK_STATUS_NOT_DESIGN && taskBak.TechnicianManID == employeeID {
			maps["stauts"] = magic.TASK_STATUS_NOT_PURCHASE

			maps["technician_final_date"] = tn
			maps["technician_remark"] = task.TechnicianRemark

			maps["purchase_end_date"] = tn.AddDate(0, 0, taskBak.PurchaseDays)
			maps["purchase_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_PURCHASE && taskBak.PurchaseManID == employeeID {
			maps["stauts"] = magic.TASK_STATUS_NOT_STORAGE

			maps["purchase_final_date"] = tn
			maps["purchase_remark"] = task.TechnicianRemark

			maps["inventory_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_STORAGE && taskBak.InventoryManID == employeeID {
			maps["inventory_final_date"] = tn
			maps["inventory_remark"] = task.TechnicianRemark
			if taskBak.Type == magic.TASK_TYPE_3 {
				maps["stauts"] = magic.TASK_STATUS_NOT_ASSEMBLY

				maps["assembly_start_date"] = tn
			} else {
				maps["stauts"] = magic.TASK_STATUS_NOT_SHIPMENT

				maps["shipment_start_date"] = tn
			}
		} else if taskBak.Status == magic.TASK_STATUS_NOT_ASSEMBLY && taskBak.TechnicianManID == employeeID {
			maps["stauts"] = magic.TASK_STATUS_NOT_SHIPMENT

			maps["assembly_final_date"] = tn
			maps["assembly_remark"] = task.TechnicianRemark

			maps["shipment_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_SHIPMENT && taskBak.ShipmentManID == employeeID {
			maps["stauts"] = magic.TASK_STATUS_SHIPMENT

			maps["shipment_final_date"] = tn
			maps["shipment_remark"] = task.TechnicianRemark
		} else {
			code = msg.FAIL
		}

		if code != msg.FAIL {
			code = models.NextTask(&models.Task{}, maps)
		}
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

// 预存款合同添加任务
func AddTask(c *gin.Context) {
	var task models.Task
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&task)

	code = models.GeneralSelect(&contractBak, task.ContractID, nil)

	if code == msg.SUCCESS && contractBak.IsPreDeposit && contractBak.PreDeposit >= task.TotalPrice {
		code = models.InsertTask(&contractBak, &task)
	}

	msg.Message(c, code, nil)
}

// 预存款合同任务驳回接口
func RejectTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		code = msg.ERROR
	} else {
		code = models.RejectTask(id)

	}

	msg.Message(c, code, nil)
}
