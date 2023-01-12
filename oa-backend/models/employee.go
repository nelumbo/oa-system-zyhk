package models

import (
	"oa-backend/utils/msg"
	"oa-backend/utils/pwd"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	// UID           string  `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`
	// OfficeUID     string  `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	// DepartmentUID string  `gorm:"type:varchar(32);comment:部门UID;default:(-)" json:"departmentUID"`

	ID            int     `gorm:"primary_key" json:"id"`
	IsDelete      bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Phone         string  `gorm:"type:varchar(50);comment:手机号" json:"phone"`
	Name          string  `gorm:"type:varchar(50);comment:姓名" json:"name"`
	Password      string  `gorm:"type:varchar(50);comment:密码" json:"password"`
	WechatID      string  `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email         string  `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	OfficeID      int     `gorm:"type:int;comment:办事处ID;default:(-)" json:"officeID"`
	Number        string  `gorm:"type:varchar(50);comment:编号" json:"number"`
	ContractCount int     `gorm:"type:int;comment:被审批合同总数" json:"contractCount"`
	Money         float64 `gorm:"type:decimal(20,6);comment:补助额度(元)" json:"money"`
	Credit        float64 `gorm:"type:decimal(20,6);comment:每月总部补助额度(元)" json:"credit"`
	OfficeCredit  float64 `gorm:"type:decimal(20,6);comment:每月办事处补助额度(元)" json:"officeCredit"`

	Office Office `gorm:"foreignKey:OfficeID" json:"office"`
	Roles  []Role `gorm:"many2many:employee_role;foreignKey:ID;references:ID" json:"roles"`

	Remark string `gorm:"-" json:"remark"`
}

func UpdateEmployeeOffice(employee *Employee) (code int) {
	var maps = make(map[string]interface{})
	if employee.OfficeID == 0 {
		maps["office_id"] = nil
	} else {
		maps["office_id"] = employee.OfficeID
	}
	maps["number"] = employee.Number
	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := tx.Model(&Employee{ID: employee.ID}).Updates(maps).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&employee).Association("Roles").Replace(employee.Roles); tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateEmployeeMoney(employee *Employee, employeeID int) (code int) {
	var maps = make(map[string]interface{})
	maps["money"] = employee.Money
	maps["credit"] = employee.Credit
	maps["office_credit"] = employee.OfficeCredit
	err = db.Transaction(func(tx *gorm.DB) error {
		var employeeBak Employee
		if tErr := tx.First(&employeeBak, employee.ID).Error; tErr != nil {
			return tErr
		}

		historyEmployee := HistoryEmployee{
			UserID:      employee.ID,
			EmployeeID:  employeeID,
			OldMoney:    employeeBak.Money,
			ChangeMoney: employee.Money - employeeBak.Money,
			NewMoney:    employee.Money,
			CreateDate:  XDate{Time: time.Now()},
			Remark:      "[直接修改] : " + employee.Remark,
		}
		if tErr := InsertHistoryEmployee(&historyEmployee, tx); tErr != nil {
			return tErr
		}

		if tErr := tx.Model(&Employee{ID: employee.ID}).Updates(maps).Error; tErr != nil {
			return tErr
		}

		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectEmployee(id int) (employee Employee, code int) {
	db.Preload("Office").Preload("Roles").
		Where("is_delete = ?", false).
		First(&employee, id)
	if employee.ID == 0 {
		return Employee{}, msg.FAIL
	}
	return employee, msg.SUCCESS
}

func SelectEmployees(employeeQuery *Employee, xForms *XForms) (employees []Employee, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	if employeeQuery.OfficeID > 0 {
		maps["office_id"] = employeeQuery.OfficeID
	}

	tx := db.Where(maps)
	if employeeQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+employeeQuery.Name+"%")
	}
	if employeeQuery.Phone != "" {
		tx = tx.Where("phone LIKE ?", "%"+employeeQuery.Phone+"%")
	}
	err = tx.Find(&employees).Count(&xForms.Total).
		Preload("Office").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&employees).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return employees, msg.SUCCESS
}

func SelectEmployeeByPhone(phone string) (employee Employee, code int) {
	db.Where("phone = ? AND is_delete = ?", phone, false).First(&employee)
	if employee.ID == 0 {
		return Employee{}, msg.FAIL
	}
	return employee, msg.SUCCESS
}

func SelectEmployeeByPhoneAndPwd(phone string, password string) (employee Employee, code int) {
	db.Where("phone = ? AND is_delete = ?", phone, false).First(&employee)
	if employee.ID == 0 {
		return employee, msg.FAIL
	}
	password, err = pwd.Scrypt(password)
	if err != nil {
		return Employee{}, msg.ERROR
	}
	if password != employee.Password {
		return Employee{}, msg.FAIL
	}
	return employee, msg.SUCCESS
}
