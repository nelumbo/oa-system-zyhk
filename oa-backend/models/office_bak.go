package models

import (
	"oa-backend/utils/msg"
	"time"

	"gorm.io/gorm"
)

type OfficeBak struct {
	ID            int     `gorm:"primary_key" json:"id"`
	Year          int     `gorm:"type:int;comment:年份" json:"year"`
	Name          string  `gorm:"type:varchar(50);comment:名称" json:"name"`
	Number        string  `gorm:"type:varchar(50);comment:编号" json:"number"`
	BusinessMoney float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`
	Money         float64 `gorm:"type:decimal(20,6);comment:办事处目前可报销额度(元)" json:"money"`
	MoneyCold     float64 `gorm:"type:decimal(20,6);comment:办事处今年冻结报销额度(元)" json:"moneyCold"`
	TaskLoad      float64 `gorm:"type:decimal(20,6);comment:今年目标量(元)" json:"taskLoad"`
	TargetLoad    float64 `gorm:"type:decimal(20,6);comment:今年完成量(元)" json:"targetLoad"`
	YWTargetLoad  float64 `gorm:"type:decimal(20,6);comment:原位目标量(元)" json:"ywTargetLoad"`
	ZYTargetLoad  float64 `gorm:"type:decimal(20,6);comment:直研目标量(元)" json:"zyTargetLoad"`
	QDTargetLoad  float64 `gorm:"type:decimal(20,6);comment:渠道目标量(元)" json:"qdTargetLoad"`
	LastYearMoney float64 `gorm:"type:decimal(20,6);comment:上年结算提成(元)" json:"lastYearMoney"`
}

func SettlementStart() (code int) {

	err = db.Transaction(func(tdb *gorm.DB) error {

		if tErr := tdb.Model(&System{}).Where("text = ?", "ice").Update("value", 1).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Model(&Office{}).Where("1=1").Update("is_set_submit", 0).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	Ice = 1
	return msg.SUCCESS
}

func SettlementEnd() (code int) {
	var offices []Office
	err = db.Where("is_delete is false").Find(&offices).Error
	if err != nil {
		return msg.ERROR
	}

	var can = true
	for i := range offices {
		if !offices[i].IsSubmit {
			can = false
			break
		}
	}

	if can {
		year := time.Now().AddDate(-1, 0, 0).Year()
		err = db.Transaction(func(tx *gorm.DB) error {
			//创建备份记录
			for i := range offices {
				officeBak := OfficeBak{
					Year:          year,
					Name:          offices[i].Name,
					Number:        offices[i].Number,
					BusinessMoney: offices[i].BusinessMoney,
					Money:         offices[i].Money,
					MoneyCold:     offices[i].MoneyCold,
					TaskLoad:      offices[i].TaskLoad,
					TargetLoad:    offices[i].TargetLoad,
					YWTargetLoad:  offices[i].YWTargetLoad,
					ZYTargetLoad:  offices[i].ZYTargetLoad,
					QDTargetLoad:  offices[i].QDTargetLoad,
					LastYearMoney: offices[i].LastYearMoney,
				}
				if tErr := tx.Create(&officeBak).Error; tErr != nil {
					return tErr
				}

				//重置办事处数据
				var maps = make(map[string]interface{})
				maps["money"] = 0
				maps["money_cold"] = 0
				maps["target_load"] = 0
				maps["yw_target_load"] = 0
				maps["zy_target_load"] = 0
				maps["qd_target_load"] = 0
				maps["is_submit"] = false
				maps["task_load"] = offices[i].NextTaskLoad
				maps["next_task_load"] = 0
				if tErr := tx.Model(&Office{}).Where("id = ?", offices[i].ID).Updates(maps).Error; tErr != nil {
					return tErr
				}
			}

			//重置员工合同记录数
			if tErr := tx.Model(&Employee{}).Where("1 = 1").Update("contract_count", 0).Error; tErr != nil {
				return tErr
			}
			//重置员工预设记录数
			if tErr := tx.Model(&Employee{}).Where("1 = 1").Update("pre_count", 0).Error; tErr != nil {
				return tErr
			}
			//系统Ice解除
			if tErr := tx.Model(&System{}).Where("text = ?", "ice").Update("value", 2).Error; tErr != nil {
				return tErr
			}
			return nil
		})
		if err != nil {
			return msg.ERROR
		} else {
			Ice = 2
			return msg.SUCCESS
		}
	}
	return msg.FAIL
}

func SettSubmit(employee *Employee, employees []Employee) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := tx.Model(&Office{}).Where("id = ?", employee.OfficeID).Update("is_set_submit", 1).Error; tErr != nil {
			return tErr
		}

		for i := range employees {
			if tErr := tx.Model(&Employee{}).Where("id = ?", employees[i].ID).Update("year_money", employees[i].YearMoney).Error; tErr != nil {
				return tErr
			}
		}

		return nil
	})

	if err != nil {
		return msg.ERROR
	} else {
		return msg.SUCCESS
	}
}

func SettApprove(emplyeeID int, officeID int, employees []Employee) (code int) {

	err = db.Transaction(func(tx *gorm.DB) error {

		for i := range employees {
			historyEmployee := HistoryEmployee{
				CreateDate:  XDate{Time: time.Now()},
				EmployeeID:  emplyeeID,
				UserID:      employees[i].ID,
				OldMoney:    employees[i].Money,
				ChangeMoney: employees[i].YearMoney,
				NewMoney:    employees[i].Money + employees[i].YearMoney,
				Remark:      "年度提成结算",
			}
			if tErr := InsertHistoryEmployee(&historyEmployee, tx); tErr != nil {
				return tErr
			}
			if tErr := tx.Model(&Employee{}).Where("id = ?", employees[i].ID).Update("money", employees[i].Money+employees[i].YearMoney).Error; tErr != nil {
				return tErr
			}
		}
		if tErr := tx.Model(&Office{}).Where("id = ?", officeID).Updates(map[string]interface{}{"is_set_submit": 2, "last_year_money": 0}).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}
