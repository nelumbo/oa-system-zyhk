package models

import (
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Expense struct {
	// UID           string  `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`
	// EmployeeUID string  `gorm:"type:varchar(32);comment:申请员工UID;default:(-)" json:"employeeUID"`
	// ApproverUID1 string  `gorm:"type:varchar(32);comment:办事处审批员工UID;default:(-)" json:"approverUID1"`
	// ApproverUID2 string  `gorm:"type:varchar(32);comment:财务审批财务员工UID;default:(-)" json:"approverUID2"`
	// ApproverUID3 string  `gorm:"type:varchar(32);comment:出纳员工UID;default:(-)" json:"approverUID3"`

	ID          int     `gorm:"primary_key" json:"id"`
	IsDelete    bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	EmployeeID  int     `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	Type        int     `gorm:"type:int;comment:报销类型(1:补助,2:提成,3:业务费,4:差旅费)" json:"type"`
	Text        string  `gorm:"type:varchar(300);comment:申请理由" json:"text"`
	Amount      float64 `gorm:"type:decimal(20,6);comment:金额(元)" json:"amount"`
	Status      int     `gorm:"type:int;comment:状态(-1:拒绝,1:待办事处审批,2:待财务审批,3:带出纳付款,4:完成)" json:"status"`
	ApproverID  int     `gorm:"type:int;comment:审批员工ID;default:(-)" json:"oapproverID"`
	FinanceID   int     `gorm:"type:int;comment:财务ID;default:(-)" json:"financeID"`
	CashierID   int     `gorm:"type:int;comment:出纳ID;default:(-)" json:"cashierID"`
	CreateDate  XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	ApproveDate XDate   `gorm:"type:date;comment:审批日期;default:(-)" json:"approveDate"`

	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Approver Employee `gorm:"foreignKey:ApproverID" json:"approver"`
	Finance  Employee `gorm:"foreignKey:FinanceID" json:"finance"`
	Cashier  Employee `gorm:"foreignKey:CashierID" json:"cashier"`

	IsPass    bool   `gorm:"-" json:"isPass"`
	StartDate string `gorm:"-" json:"startDate"`
	EndDate   string `gorm:"-" json:"endDate"`
}

func DeleteExpense(id int, employeeID int) (code int) {
	err = db.Model(&Expense{ID: id}).
		Where(map[string]interface{}{"employee_id": employeeID, "status": -1}).
		Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateExpense(expense *Expense, expenseBak *Expense, maps map[string]interface{}) (code int) {

	if expense.Type == magic.EXPENSE_TYPE_CaiLvFei {
		return GeneralUpdate(&Expense{}, expense.ID, maps)
	} else {
		if expense.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_2 && expense.IsPass {
			if expense.Type == magic.EXPENSE_TYPE_BuZhu {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE employee SET money = money - ? WHERE id = ?", expense.Amount, expense.EmployeeID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			} else if expense.Type == magic.EXPENSE_TYPE_TiChen {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE office SET money = money - ? WHERE id = ?", expense.Amount, expense.Employee.Office.ID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			} else if expense.Type == magic.EXPENSE_TYPE_YeWuFei {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE office SET business_money = business_money - ? WHERE id = ?", expense.Amount, expense.Employee.Office.ID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			}
		} else if expense.Status == magic.EXPENSE_STATUS_NOT_PAYMENT && !expense.IsPass {
			if expense.Type == magic.EXPENSE_TYPE_BuZhu {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE employee SET money = money + ? WHERE id = ?", expense.Amount, expense.EmployeeID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			} else if expense.Type == magic.EXPENSE_TYPE_TiChen {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE office SET money = money + ? WHERE id = ?", expense.Amount, expense.Employee.Office.ID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			} else if expense.Type == magic.EXPENSE_TYPE_YeWuFei {
				err = db.Transaction(func(tx *gorm.DB) error {
					if tErr := tx.Exec("UPDATE office SET business_money = business_money + ? WHERE id = ?", expense.Amount, expense.Employee.Office.ID).Error; tErr != nil {
						return tErr
					}
					if tErr := tx.Model(&Expense{ID: expense.ID}).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			}
		}
	}

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectExpense(id int) (expense Expense, code int) {
	db.Preload("Employee.Office").Preload("Approver").Preload("Finance").Preload("Cashier").
		Where("is_delete = ?", false).
		First(&expense, id)
	if expense.ID == 0 {
		return Expense{}, msg.FAIL
	}
	return expense, msg.SUCCESS
}

func SelectExpenses(expenseQuery *Expense, xForms *XForms) (expenses []Expense, code int) {
	var maps = make(map[string]interface{})
	maps["expense.is_delete"] = false
	if expenseQuery.Type != 0 {
		maps["expense.type"] = expenseQuery.Type
	}
	if expenseQuery.Status != 0 {
		maps["expense.status"] = expenseQuery.Status
	}
	if expenseQuery.EmployeeID != 0 {
		maps["expense.employee_id"] = expenseQuery.EmployeeID
	}
	if expenseQuery.Employee.OfficeID != 0 {
		maps["Employee.office_id"] = expenseQuery.Employee.OfficeID
	}

	tx := db.Joins("Employee").Where(maps)

	if expenseQuery.Employee.Name != "" {
		tx = tx.Where("Employee.name LIKE ?", "%"+expenseQuery.Employee.Name+"%")
	}

	if expenseQuery.StartDate != "" && expenseQuery.EndDate != "" {
		tx = tx.Where("expense.create_date BETWEEN ? AND ?", expenseQuery.StartDate, expenseQuery.EndDate)
	} else {
		if expenseQuery.StartDate != "" {
			tx = tx.Where("expense.create_date >= ?", expenseQuery.StartDate)
		}
		if expenseQuery.EndDate != "" {
			tx = tx.Where("expense.create_date <= ?", expenseQuery.EndDate)
		}
	}

	err = tx.Find(&expenses).Count(&xForms.Total).
		Preload("Employee.Office").Preload("Approver").Preload("Finance").Preload("Cashier").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("id desc").Find(&expenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return expenses, msg.SUCCESS
}
