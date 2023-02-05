package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type HistoryEmployee struct {
	ID          int     `gorm:"primary_key" json:"id"`
	CreateDate  XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EmployeeID  int     `gorm:"type:int;comment:操作人ID;default:(-)" json:"employeeID"`
	UserID      int     `gorm:"type:int;comment:员工ID;default:(-)" json:"userID"`
	OldMoney    float64 `gorm:"type:decimal(20,6);comment:原补助额度(元)" json:"oldMoney"`
	ChangeMoney float64 `gorm:"type:decimal(20,6);comment:修改补助额度(元)" json:"changeMoney"`
	NewMoney    float64 `gorm:"type:decimal(20,6);comment:新补助额度(元)" json:"newMoney"`

	Remark string `gorm:"type:varchar(200);comment:备注" json:"remark"`

	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	User     Employee `gorm:"foreignKey:UserID" json:"user"`

	StartDate string `gorm:"-" json:"startDate"`
	EndDate   string `gorm:"-" json:"endDate"`
}

type HistoryOffice struct {
	ID                  int     `gorm:"primary_key" json:"id"`
	CreateDate          XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	OfficeID            int     `gorm:"type:int;comment:办事处ID;default:(-)" json:"officeID"`
	EmployeeID          int     `gorm:"type:int;comment:操作人ID;default:(-)" json:"employeeID"`
	OldBusinessMoney    float64 `gorm:"type:decimal(20,6);comment:旧业务费用(元)" json:"oldBusinessMoney"`
	OldMoney            float64 `gorm:"type:decimal(20,6);comment:旧办事处目前可报销额度(元)" json:"oldMoney"`
	OldMoneyCold        float64 `gorm:"type:decimal(20,6);comment:旧办事处今年冻结报销额度(元)" json:"oldMoneyCold"`
	OldTargetLoad       float64 `gorm:"type:decimal(20,6);comment:旧今年完成量(元)" json:"oldTargetLoad"`
	ChangeBusinessMoney float64 `gorm:"type:decimal(20,6);comment:修改业务费用(元)" json:"changeBusinessMoney"`
	ChangeMoney         float64 `gorm:"type:decimal(20,6);comment:修改办事处目前可报销额度(元)" json:"changeMoney"`
	ChangeMoneyCold     float64 `gorm:"type:decimal(20,6);comment:修改办事处今年冻结报销额度(元)" json:"changeMoneyCold"`
	ChangeTargetLoad    float64 `gorm:"type:decimal(20,6);comment:修改今年目标量(元)" json:"changeTargetLoad"`
	NewBusinessMoney    float64 `gorm:"type:decimal(20,6);comment:新业务费用(元)" json:"newBusinessMoney"`
	NewMoney            float64 `gorm:"type:decimal(20,6);comment:新办事处目前可报销额度(元)" json:"newMoney"`
	NewMoneyCold        float64 `gorm:"type:decimal(20,6);comment:新办事处今年冻结报销额度(元)" json:"newMoneyCold"`
	NewTargetLoad       float64 `gorm:"type:decimal(20,6);comment:新今年目标量(元)" json:"newTargetLoad"`

	Remark string `gorm:"type:varchar(200);comment:备注" json:"remark"`

	Office   Office   `gorm:"foreignKey:OfficeID" json:"office"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`

	StartDate string `gorm:"-" json:"startDate"`
	EndDate   string `gorm:"-" json:"endDate"`
}

type HistoryProduct struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreateDate XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EmployeeID int    `gorm:"type:int;comment:操作人ID;default:(-)" json:"employeeID"`
	ProductID  int    `gorm:"type:int;comment:产品ID" json:"productID"`
	Number     int    `gorm:"type:int;comment:数量" json:"number"`
	Remark     string `gorm:"type:varchar(200);comment:备注" json:"remark"`

	Product  Product  `gorm:"foreignKey:ProductID" json:"product"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
}

func InsertHistoryEmployee(historyEmployee *HistoryEmployee, tx *gorm.DB) (err error) {
	tErr := tx.Create(&historyEmployee).Error
	if tErr != nil {
		return tErr
	}
	return nil
}

func InsertHistoryOffice(historyOffice *HistoryOffice, tx *gorm.DB) (err error) {
	tErr := tx.Create(&historyOffice).Error
	if tErr != nil {
		return tErr
	}
	return nil
}

func InsertHistoryProduct(historyProduct *HistoryProduct, tx *gorm.DB) (err error) {
	tErr := tx.Create(&historyProduct).Error
	if tErr != nil {
		return tErr
	}
	return nil
}

func SelectHistoryEmployees(historyEmployee *HistoryEmployee, xForms *XForms) (historyEmployees []HistoryEmployee, code int) {
	var maps = make(map[string]interface{})
	if historyEmployee.UserID != 0 {
		maps["history_employee.user_id"] = historyEmployee.UserID
	}

	tx := db.Where(maps).Joins("User")

	if historyEmployee.User.OfficeID != 0 {
		tx = tx.Joins("left join office on User.office_id = office.id").Where("office.id = ?", historyEmployee.User.OfficeID)
	}

	if historyEmployee.User.Name != "" {
		tx = tx.Where("User.name Like ?", "%"+historyEmployee.User.Name+"%")
	}

	if historyEmployee.Remark != "" {
		tx = tx.Where("history_employee.remark LIKE ?", "%"+historyEmployee.Remark+"%")
	}

	if historyEmployee.StartDate != "" && historyEmployee.EndDate != "" {
		tx = tx.Where("history_employee.create_date BETWEEN ? AND ?", historyEmployee.StartDate, historyEmployee.EndDate)
	} else {
		if historyEmployee.StartDate != "" {
			tx = tx.Where("history_employee.create_date >= ?", historyEmployee.StartDate)
		}
		if historyEmployee.EndDate != "" {
			tx = tx.Where("history_employee.create_date <= ?", historyEmployee.EndDate)
		}
	}

	err = tx.Find(&historyEmployees).Count(&xForms.Total).
		Preload("User.Office").Preload("Employee").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("history_employee.create_date desc").
		Find(&historyEmployees).Error
	if err != nil {
		return historyEmployees, msg.ERROR
	}
	return historyEmployees, msg.SUCCESS
}

func SelectHistoryOffices(historyOffice *HistoryOffice, xForms *XForms) (historyOffices []HistoryOffice, code int) {
	var maps = make(map[string]interface{})
	if historyOffice.OfficeID != 0 {
		maps["office_id"] = historyOffice.OfficeID
	}

	tx := db.Where(maps)
	if historyOffice.Remark != "" {
		tx = tx.Where("remark LIKE ?", "%"+historyOffice.Remark+"%")
	}

	if historyOffice.StartDate != "" && historyOffice.EndDate != "" {
		tx = tx.Where("create_date BETWEEN ? AND ?", historyOffice.StartDate, historyOffice.EndDate)
	} else {
		if historyOffice.StartDate != "" {
			tx = tx.Where("create_date >= ?", historyOffice.StartDate)
		}
		if historyOffice.EndDate != "" {
			tx = tx.Where("create_date <= ?", historyOffice.EndDate)
		}
	}

	err = tx.Find(&historyOffices).Count(&xForms.Total).
		Order("id desc").
		Preload("Office").Preload("Employee").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("history_office.create_date desc").
		Find(&historyOffices).Error
	if err != nil {
		return historyOffices, msg.ERROR
	}
	return historyOffices, msg.SUCCESS
}

func SelectHistoryProducts(historyProduct *HistoryProduct, xForms *XForms) (historyProducts []HistoryProduct, code int) {
	tx := db.Joins("Product")
	if historyProduct.Product.Name != "" {
		tx = tx.Where("Product.name Like ?", "%"+historyProduct.Product.Name+"%")
	}

	err = tx.Find(&historyProducts).Count(&xForms.Total).
		Preload("Employee").Preload("Product").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("history_product.create_date desc").
		Find(&historyProducts).Error
	if err != nil {
		return historyProducts, msg.ERROR
	}
	return historyProducts, msg.SUCCESS
}
