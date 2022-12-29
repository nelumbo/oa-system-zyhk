package models

import (
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"gorm.io/gorm"
)

type Predesign struct {
	ID            int    `gorm:"primary_key" json:"id"`
	IsDelete      bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	EmployeeID    int    `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	AuditorID     int    `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateRemark  string `gorm:"type:varchar(300);comment:业务员备注" json:"createRemark"`
	ApproveRemark string `gorm:"type:varchar(300);comment:审核备注" json:"approveRemark"`
	Status        int    `gorm:"type:int;comment:状态(-1:驳回 1:未审批 2:未完成 3:已完成)" json:"status"`
	CreateDate    XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate     XDate  `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	FinalDate     XDate  `gorm:"type:date;comment:完成日期;default:(-)" json:"finalDate"`

	PredesignTasks []PredesignTask `gorm:"-" json:"predesignTasks"`
	Employee       Employee        `gorm:"foreignKey:EmployeeID" json:"employee"`
	Auditor        Employee        `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass        bool          `gorm:"-" json:"isPass"`
	PredesignTask PredesignTask `gorm:"-" json:"predesignTask"`
}

type PredesignTask struct {
	ID           int    `gorm:"primary_key" json:"id"`
	PredesignID  int    `gorm:"type:int;comment:预设计ID;default:(-)" json:"predesignID"`
	EmployeeID   int    `gorm:"type:int;comment:创建者ID;default:(-)" json:"employeeID"`
	TechnicistID int    `gorm:"type:int;comment:技术员ID;default:(-)" json:"technicistID"`
	AuditorID    int    `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateRemark string `gorm:"type:varchar(300);comment:设计要求" json:"createRemark"`
	FinalRemark  string `gorm:"type:varchar(300);comment:完成备注" json:"finalRemark"`
	Days         int    `gorm:"type:int;comment:分配工作天数" json:"days"`
	CreateDate   XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EndDate      XDate  `gorm:"type:date;comment:预计结束工作日期" json:"endDate"`
	FinalDate    XDate  `gorm:"type:date;comment:完成工作日期" json:"finalDate"`
	ApproveDate  XDate  `gorm:"type:date;comment:审核日期" json:"approveDate"`
	// Status       int    `gorm:"type:int;comment:状态( 1:未提交 2:未审核 3:未通过 4:已通过)" json:"status"`
	Status int `gorm:"type:int;comment:状态( -1:未通过 1:未提交 2:未审核 3:已通过)" json:"status"`

	Employee   Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Technicist Employee `gorm:"foreignKey:TechnicistID" json:"technicist"`
	Auditor    Employee `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass          bool   `gorm:"-" json:"isPass"`
	NewTechnicistID int    `gorm:"-" json:"newTechnicistID"`
	NewCreateRemark string `gorm:"-" json:"newCreateRemark"`
	NewDays         int    `gorm:"-" json:"newDays"`
}

func UpdatePredesign(predesign *Predesign, maps map[string]interface{}) (code int) {
	var predesignTask PredesignTask
	predesignTask.PredesignID = predesign.ID
	predesignTask.EmployeeID = predesign.AuditorID
	predesignTask.TechnicistID = predesign.PredesignTask.TechnicistID
	predesignTask.CreateRemark = predesign.PredesignTask.CreateRemark
	predesignTask.Days = predesign.PredesignTask.Days
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
	db.Preload("PredesignTasks").Preload("Employee.Office").Preload("Auditor").
		Where("is_delete = ?", false).
		First(&predesign, id)
	if predesign.ID == 0 {
		return Predesign{}, msg.FAIL
	}
	return predesign, msg.SUCCESS
}

func SelectPredesigns(predesignQuery *Predesign, xForms *XForms) (predesigns []Predesign, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false

	if predesignQuery.Status != 0 {
		maps["predesign.status"] = predesignQuery.Status
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

func InsertPredesignTask(predesignTaskOld *PredesignTask, maps map[string]interface{}) (code int) {
	var predesignTask PredesignTask
	predesignTask.PredesignID = predesignTaskOld.PredesignID
	predesignTask.EmployeeID = predesignTaskOld.AuditorID
	predesignTask.TechnicistID = predesignTaskOld.NewTechnicistID
	predesignTask.CreateRemark = predesignTaskOld.NewCreateRemark
	predesignTask.Days = predesignTaskOld.NewDays
	predesignTask.CreateDate.Time = time.Now()
	predesignTask.EndDate.Time = predesignTask.CreateDate.Time.AddDate(0, 0, predesignTask.Days)
	predesignTask.Status = magic.PREDESIGN_TASK_STATUS_NOT_SUBMIT

	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := tx.Model(&PredesignTask{ID: predesignTaskOld.ID}).Updates(maps).Error; tErr != nil {
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
