package models

import (
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"gorm.io/gorm"
)

type Predesign struct {
	ID         int    `gorm:"primary_key" json:"id"`
	IsDelete   bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	EmployeeID int    `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	AuditorID  int    `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	Remark     string `gorm:"type:varchar(300);comment:设计需求" json:"remark"`
	Status     int    `gorm:"type:int;comment:状态(-1:驳回 1:未审批 2:未完成 3:已完成)" json:"status"`
	CreateDate XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate  XDate  `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	FinalDate  XDate  `gorm:"type:date;comment:完成日期;default:(-)" json:"finalDate"`

	PredesignTasks []PredesignTask `json:"predesignTasks"`
	Employee       Employee        `gorm:"foreignKey:EmployeeID" json:"employee"`
	Auditor        Employee        `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass        bool          `gorm:"-" json:"isPass"`
	PredesignTask PredesignTask `gorm:"-" json:"predesignTask"`
}

type PredesignTask struct {
	ID           int    `gorm:"primary_key" json:"id"`
	IsDelete     bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	PredesignID  int    `gorm:"type:int;comment:预设计ID;default:(-)" json:"predesignID"`
	CreaterID    int    `gorm:"type:int;comment:创建者ID;default:(-)" json:"createrID"`
	EmployeeID   int    `gorm:"type:int;comment:技术员ID;default:(-)" json:"employeeID"`
	AuditorID    int    `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateRemark string `gorm:"type:varchar(300);comment:设计要求" json:"createRemark"`
	FinalRemark  string `gorm:"type:varchar(300);comment:技术员备注" json:"finalRemark"`
	Days         int    `gorm:"type:int;comment:工作天数" json:"days"`
	CreateDate   XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EndDate      XDate  `gorm:"type:date;comment:预期完成时间" json:"endDate"`
	FinalDate    XDate  `gorm:"type:date;comment:提交日期" json:"finalDate"`
	ApproveDate  XDate  `gorm:"type:date;comment:审核日期" json:"approveDate"`
	Status       int    `gorm:"type:int;comment:状态(1:未提交 2:未审核 3:未通过 4:已通过)" json:"status"`

	Creater  Employee `gorm:"foreignKey:CreaterID" json:"creater"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Auditor  Employee `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass          bool   `gorm:"-" json:"isPass"`
	NewCreateRemark string `gorm:"-" json:"newCreateRemark"`
	NewDays         int    `gorm:"-" json:"newDays"`
}

func UpdatePredesign(predesign *Predesign, maps map[string]interface{}) (code int) {
	var predesignTask PredesignTask
	predesignTask.PredesignID = predesign.ID
	predesignTask.CreaterID = predesign.AuditorID
	predesignTask.EmployeeID = predesign.PredesignTask.EmployeeID
	predesignTask.Days = predesign.PredesignTask.Days
	predesignTask.CreateRemark = predesign.PredesignTask.CreateRemark
	predesignTask.CreateDate.Time = time.Now()
	predesignTask.EndDate.Time = predesignTask.CreateDate.Time.AddDate(0, 0, predesignTask.Days)
	predesignTask.Status = magic.PREDESIGN_TASK_STATUS_NOT_SUBMIT

	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := tx.Model(&Predesign{ID: predesign.ID}).Updates(maps).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Create(&predesignTask).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectPredesign(id int) (predesign Predesign, code int) {
	db.Preload("PredesignTasks.Employee").Preload("Employee.Office").Preload("Auditor").
		Where("is_delete = ?", false).
		First(&predesign, id)
	if predesign.ID == 0 {
		return Predesign{}, msg.FAIL
	}
	return predesign, msg.SUCCESS
}

func SelectPredesigns(predesignQuery *Predesign, xForms *XForms) (predesigns []Predesign, code int) {
	var maps = make(map[string]interface{})
	maps["predesign.is_delete"] = false

	if predesignQuery.Status != 0 {
		maps["predesign.status"] = predesignQuery.Status
	}
	if predesignQuery.EmployeeID != 0 {
		maps["predesign.employee_id"] = predesignQuery.EmployeeID
	}

	tx := db.Where(maps)

	if predesignQuery.Employee.Name != "" {
		tx = tx.Joins("Employee").Where("Employee.Name LIKE ?", "%"+predesignQuery.Employee.Name+"%")
	}

	err = tx.Find(&predesigns).Count(&xForms.Total).
		Preload("Employee.Office").Preload("Auditor").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("id desc").Find(&predesigns).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return predesigns, msg.SUCCESS
}

func ApprovePredesignTask(predesignTaskBak *PredesignTask, maps map[string]interface{}) (code int) {

	if predesignTaskBak.IsPass {
		err = db.Transaction(func(tx *gorm.DB) error {
			if tErr := tx.Model(&PredesignTask{ID: predesignTaskBak.ID}).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tx.Model(&Predesign{}).
				Where("id", predesignTaskBak.PredesignID).
				Updates(map[string]interface{}{"status": magic.PREDESIGN_STATUS_FINAL, "final_date": time.Now()}).
				Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		var predesignTask PredesignTask
		predesignTask.PredesignID = predesignTaskBak.PredesignID
		predesignTask.CreaterID = predesignTaskBak.AuditorID
		predesignTask.EmployeeID = predesignTaskBak.EmployeeID
		predesignTask.CreateRemark = predesignTaskBak.NewCreateRemark
		predesignTask.Days = predesignTaskBak.NewDays
		predesignTask.CreateDate.Time = time.Now()
		predesignTask.EndDate.Time = predesignTask.CreateDate.Time.AddDate(0, 0, predesignTask.Days)
		predesignTask.Status = magic.PREDESIGN_TASK_STATUS_NOT_SUBMIT

		err = db.Transaction(func(tx *gorm.DB) error {
			if tErr := tx.Model(&PredesignTask{ID: predesignTaskBak.ID}).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tx.Create(&predesignTask).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	}
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectPredesignTasks(predesignTaskQuery *PredesignTask, xForms *XForms) (predesignTasks []PredesignTask, code int) {
	var maps = make(map[string]interface{})
	maps["predesign_task.is_delete"] = false

	if predesignTaskQuery.Status != 0 {
		maps["predesign_task.status"] = predesignTaskQuery.Status
	}
	if predesignTaskQuery.EmployeeID != 0 {
		maps["predesign_task.employee_id"] = predesignTaskQuery.EmployeeID
	}

	tx := db.Where(maps)

	if predesignTaskQuery.Employee.Name != "" {
		tx = tx.Joins("Employee").Where("Employee.Name LIKE ?", "%"+predesignTaskQuery.Employee.Name+"%")
	}

	err = tx.Find(&predesignTasks).Count(&xForms.Total).
		Preload("Creater").Preload("Employee").Preload("Auditor").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("id desc").Find(&predesignTasks).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return predesignTasks, msg.SUCCESS
}
