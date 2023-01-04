package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Bidbond struct {
	ID           int     `gorm:"primary_key" json:"id"`
	IsDelete     bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	EmployeeID   int     `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	AuditorID    int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	FinanceID    int     `gorm:"type:int;comment:财务ID;default:(-)" json:"financeID"`
	Money        float64 `gorm:"type:decimal(20,6);comment:保证金金额(元)" json:"money"`
	CreateRemark string  `gorm:"type:varchar(300);comment:备注" json:"createRemark"`
	FinalRemark  string  `gorm:"type:varchar(300);comment:回款备注" json:"finalRemark"`
	Status       int     `gorm:"type:int;comment:状态(-1:驳回 1:待审核 2:审核锁定 3:已回款)" json:"status"`
	CreateDate   XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate    XDate   `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	FinalDate    XDate   `gorm:"type:date;comment:回款日期;default:(-)" json:"finalDate"`

	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Auditor  Employee `gorm:"foreignKey:AuditorID" json:"auditor"`
	Finalce  Employee `gorm:"foreignKey:FinanceID" json:"finalce"`

	IsPass bool `gorm:"-" json:"isPass"`
}

func SelectBidbond(id int) (bidbond Bidbond, code int) {
	db.Preload("Employee").Preload("Auditor").Preload("Finalce").
		Where("is_delete = ?", false).
		First(&bidbond, id)
	if bidbond.ID == 0 {
		return Bidbond{}, msg.FAIL
	}
	return bidbond, msg.SUCCESS
}

func SelectBidbonds(bidbondQuery *Bidbond, xForms *XForms) (bidbonds []Bidbond, code int) {
	var maps = make(map[string]interface{})
	maps["bidbond.is_delete"] = false
	if bidbondQuery.Status != 0 {
		maps["bidbond.status"] = bidbondQuery.Status
	}
	if bidbondQuery.EmployeeID != 0 {
		maps["bidbond.employee_id"] = bidbondQuery.EmployeeID
	}
	if bidbondQuery.Employee.OfficeID != 0 {
		maps["Employee.office_id"] = bidbondQuery.Employee.OfficeID
	}

	tx := db.Where(maps).Joins("Employee")
	if bidbondQuery.Employee.Name != "" {
		tx = tx.Where("Employee.Name LIKE ?", "%"+bidbondQuery.Employee.Name+"%")
	}
	err = tx.Find(&bidbonds).Count(&xForms.Total).
		Preload("Employee.Office").Preload("Auditor").Preload("Finalce").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("id desc").Find(&bidbonds).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return bidbonds, msg.SUCCESS
}
