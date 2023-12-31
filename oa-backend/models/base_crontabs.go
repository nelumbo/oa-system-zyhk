package models

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitCronTabs() {
	monthlyAllowanceToEmployee()
	annualAudit()
}

func annualAudit() {
	cronTask := cron.New(cron.WithSeconds())
	_, err := cronTask.AddFunc("0 0 0 1 1 ? *", func() {

		err = db.Model(&System{}).Where("text = ?", "ice").Update("value", 0).Error
		if err != nil {
			fmt.Println(err)
		} else {
			Ice = 0
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	cronTask.Start()
}

func monthlyAllowanceToEmployee() {

	cronTask := cron.New(cron.WithSeconds())
	_, err := cronTask.AddFunc("0 0 1 1 * ?", func() {

		err = db.Transaction(func(tdb *gorm.DB) error {
			var employees []Employee
			if tErr := tdb.Find(&employees, "is_delete = ?", false).Error; tErr != nil {
				return tErr
			}
			for i := range employees {

				var tempRoleCredit = employees[i].RoleCredit - employees[i].RoleCreditDel
				if tempRoleCredit < 0 {
					tempRoleCredit = 0
				}

				historyEmployee := HistoryEmployee{
					UserID:      employees[i].ID,
					ChangeMoney: employees[i].Credit + employees[i].OfficeCredit + tempRoleCredit,
					Remark:      "每月补助",
				}
				if tErr := InsertHistoryEmployee(&historyEmployee, tdb); tErr != nil {
					return tErr
				}

				//职位补贴并重置每月职位补贴扣除
				if tErr := tdb.Exec("UPDATE employee SET money = money + ?,role_credit_del = 0 WHERE id = ?", tempRoleCredit, employees[i].ID).Error; tErr != nil {
					return tErr
				}
			}

			if tErr := tdb.Exec("UPDATE employee SET money = money + credit WHERE is_delete = ?", false).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Exec("UPDATE employee SET money = money + office_credit WHERE is_delete = ?", false).Error; tErr != nil {
				return tErr
			}

			var offices []Office
			if tErr := tdb.Raw("SELECT office_id AS id,sum(office_credit) AS money FROM employee WHERE office_id IS NOT NULL AND is_delete IS false GROUP BY office_id").Scan(&offices).Error; tErr != nil {
				return tErr
			}
			for i := range offices {
				var historyOffice HistoryOffice
				historyOffice.OfficeID = offices[i].ID
				historyOffice.ChangeMoney = 0 - offices[i].Money
				historyOffice.Remark = "每月业务员补贴"
				if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
					return tErr
				}
			}
			if tErr := tdb.Exec("UPDATE office,(SELECT office_id,sum(office_credit) AS sum FROM employee WHERE office_id IS NOT NULL AND is_delete IS false GROUP BY office_id) AS a SET office.money = office.money - a.sum WHERE office.id = a.office_id").Error; tErr != nil {
				return tErr
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}

	})
	if err != nil {
		fmt.Println(err)
	}
	cronTask.Start()

}
