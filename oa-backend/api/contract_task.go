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
			maps["technician_days"] = task.TechnicianDays
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

	if taskBak.ID != 0 {
		var maps = make(map[string]interface{})
		employeeID := c.MustGet("employeeID").(int)
		tn := time.Now()
		if taskBak.Status == magic.TASK_STATUS_NOT_DESIGN && taskBak.TechnicianManID == employeeID {
			maps["status"] = magic.TASK_STATUS_NOT_PURCHASE

			maps["technician_final_date"] = tn
			maps["technician_remark"] = task.Remark

			maps["purchase_end_date"] = tn.AddDate(0, 0, taskBak.PurchaseDays)
			maps["purchase_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_PURCHASE && taskBak.PurchaseManID == employeeID {
			maps["status"] = magic.TASK_STATUS_NOT_STORAGE

			maps["purchase_final_date"] = tn
			maps["purchase_remark"] = task.Remark

			maps["inventory_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_STORAGE && taskBak.InventoryManID == employeeID {
			maps["inventory_final_date"] = tn
			maps["inventory_remark"] = task.Remark
			if taskBak.Type == magic.TASK_TYPE_3 {
				maps["status"] = magic.TASK_STATUS_NOT_ASSEMBLY

				maps["assembly_start_date"] = tn
			} else {
				maps["stauts"] = magic.TASK_STATUS_NOT_SHIPMENT

				maps["shipment_start_date"] = tn
			}
		} else if taskBak.Status == magic.TASK_STATUS_NOT_ASSEMBLY && taskBak.TechnicianManID == employeeID {
			maps["status"] = magic.TASK_STATUS_NOT_SHIPMENT

			maps["assembly_final_date"] = tn
			maps["assembly_remark"] = task.Remark

			maps["shipment_start_date"] = tn
		} else if taskBak.Status == magic.TASK_STATUS_NOT_SHIPMENT && taskBak.ShipmentManID == employeeID {
			maps["status"] = magic.TASK_STATUS_SHIPMENT

			maps["shipment_final_date"] = tn
			maps["shipment_remark"] = task.Remark
		} else {
			code = msg.FAIL
		}

		if code != msg.FAIL {
			code = models.NextTask(&taskBak, maps)
		}
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

// 合同任务重置
func ResetTask(c *gin.Context) {
	var task, taskBak models.Task
	var contractBak models.Contract
	_ = c.ShouldBindJSON(&task)

	code = models.GeneralSelect(&taskBak, task.ID, nil)
	_ = models.GeneralSelect(&contractBak, task.ContractID, nil)

	if taskBak.ID != 0 &&
		taskBak.Status != magic.TASK_STATUS_REJECT &&
		taskBak.Status != magic.TASK_STATUS_NOT_DISTRIBUTE {

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

			maps["technician_man_id"] = nil
			maps["purchase_man_id"] = nil
			maps["technician_days"] = nil
			maps["purchase_days"] = nil
			maps["technician_start_date"] = nil
			maps["technician_end_date"] = nil
			maps["technician_final_date"] = nil
			maps["purchase_start_date"] = nil
			maps["purchase_end_date"] = nil
			maps["purchase_final_date"] = nil
			maps["inventory_start_date"] = tn
			maps["inventory_final_date"] = nil
			maps["assembly_start_date"] = nil
			maps["assembly_final_date"] = nil
			maps["shipment_start_date"] = nil
			maps["shipment_final_date"] = nil
			maps["purchase_remark"] = nil
			maps["inventory_remark"] = nil
			maps["inventory_remark"] = nil
			maps["assembly_remark"] = nil
			maps["shipment_remark"] = nil
		case magic.TASK_TYPE_2:
			maps["status"] = magic.TASK_STATUS_NOT_PURCHASE

			maps["technician_man_id"] = nil
			maps["purchase_man_id"] = task.PurchaseManID
			maps["technician_days"] = nil
			maps["purchase_days"] = task.PurchaseDays
			maps["technician_start_date"] = nil
			maps["technician_end_date"] = nil
			maps["technician_final_date"] = nil
			maps["purchase_start_date"] = tn
			maps["purchase_end_date"] = tn.AddDate(0, 0, task.PurchaseDays)
			maps["purchase_final_date"] = nil
			maps["inventory_start_date"] = nil
			maps["inventory_final_date"] = nil
			maps["assembly_start_date"] = nil
			maps["assembly_final_date"] = nil
			maps["shipment_start_date"] = nil
			maps["shipment_final_date"] = nil
			maps["purchase_remark"] = nil
			maps["inventory_remark"] = nil
			maps["inventory_remark"] = nil
			maps["assembly_remark"] = nil
			maps["shipment_remark"] = nil
		case magic.TASK_TYPE_3:
			maps["status"] = magic.TASK_STATUS_NOT_DESIGN

			maps["technician_man_id"] = task.TechnicianManID
			maps["purchase_man_id"] = task.PurchaseManID
			maps["technician_days"] = task.TechnicianDays
			maps["purchase_days"] = task.PurchaseDays
			maps["technician_start_date"] = tn
			maps["technician_end_date"] = tn.AddDate(0, 0, task.PurchaseDays)
			maps["technician_final_date"] = nil
			maps["purchase_start_date"] = nil
			maps["purchase_end_date"] = nil
			maps["purchase_final_date"] = nil
			maps["inventory_start_date"] = nil
			maps["inventory_final_date"] = nil
			maps["assembly_start_date"] = nil
			maps["assembly_final_date"] = nil
			maps["shipment_start_date"] = nil
			maps["shipment_final_date"] = nil
			maps["purchase_remark"] = nil
			maps["inventory_remark"] = nil
			maps["inventory_remark"] = nil
			maps["assembly_remark"] = nil
			maps["shipment_remark"] = nil
		}

		code = models.GeneralUpdate(&models.Task{}, task.ID, maps)
	} else {
		code = msg.FAIL
	}

	msg.Message(c, code, nil)
}

// 预存款合同添加任务
func AddTask(c *gin.Context) {
	var task models.Task
	var contractBak models.Contract
	var product models.Product
	_ = c.ShouldBindJSON(&task)

	_ = models.GeneralSelect(&contractBak, task.ContractID, nil)
	_ = models.GeneralSelect(&product, task.ProductID, nil)

	if contractBak.ID != 0 && product.ID != 0 &&
		contractBak.IsPreDeposit && contractBak.PreDeposit >= task.TotalPrice {

		task.TotalPrice = float64(task.Number) * task.Price
		task.ProductAttributeID = product.AttributeID

		code = models.InsertTask(&contractBak, &task)
	}

	msg.Message(c, code, nil)
}

// 预存款合同任务驳回接口
func RejectTask(c *gin.Context) {
	var task models.Task
	_ = c.ShouldBindJSON(&task)

	if err != nil {
		code = msg.ERROR
	} else {
		code = models.RejectTask(task.ID)
	}

	msg.Message(c, code, nil)
}

func DelTask(c *gin.Context) {
	var task models.Task
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		code = msg.ERROR
	} else {
		task, code = models.SelectTask(id)
		if code == msg.SUCCESS && task.Status == magic.TASK_STATUS_REJECT && task.Contract.IsPreDeposit {
			code = models.DeleteTask(id)
		} else {
			code = msg.FAIL
		}
	}
	msg.Message(c, code, nil)
}
