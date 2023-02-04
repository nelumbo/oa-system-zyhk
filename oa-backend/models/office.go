package models

import (
	"oa-backend/utils/msg"
	"time"

	"gorm.io/gorm"
)

type Office struct {
	// UID string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`

	ID            int     `gorm:"primary_key" json:"id"`
	IsDelete      bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
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
	RoleID        int     `gorm:"type:int;comment:角色ID;default:(-)" json:"roleID"`
	IsRanking     bool    `gorm:"type:boolean;comment:是否进入排行榜" json:"isRanking"`
	IsSubmit      bool    `gorm:"type:boolean;comment:今年结算是否提交" json:"isSubmit"`
	LastYearMoney float64 `gorm:"type:decimal(20,6);comment:上年结算提成(元)" json:"lastYearMoney"`
	NextTaskLoad  float64 `gorm:"type:decimal(20,6);comment:今年目标量(元)" json:"nextTaskLoad"`
	IsSetSubmit   int     `gorm:"type:int;comment:提成方案是否提交(-1：提交被驳回 0:未提交 1：已提交 2：通过)" json:"isSetSubmit"`

	Role Role `gorm:"foreignKey:RoleID" json:"role"`

	FinalPercentages float64 `gorm:"-" json:"finalPercentages"`
	NotPayment       float64 `gorm:"-" json:"notPayment"`
	Remark           string  `gorm:"-" json:"remark"`
	IsPass           bool    `gorm:"-" json:"isPass"`
}

func UpdateOfficeMoney(office *Office, employeeID int) (code int) {
	var maps = make(map[string]interface{})
	maps["business_money"] = office.BusinessMoney
	maps["money"] = office.Money
	maps["money_cold"] = office.MoneyCold
	maps["target_load"] = office.TargetLoad
	err = db.Transaction(func(tx *gorm.DB) error {
		var officeBak Office
		if tErr := tx.First(&officeBak, office.ID).Error; tErr != nil {
			return tErr
		}

		historyOffice := HistoryOffice{
			OfficeID:            officeBak.ID,
			EmployeeID:          employeeID,
			OldBusinessMoney:    officeBak.BusinessMoney,
			OldMoney:            officeBak.Money,
			OldMoneyCold:        officeBak.MoneyCold,
			OldTargetLoad:       officeBak.TargetLoad,
			ChangeBusinessMoney: office.BusinessMoney - officeBak.BusinessMoney,
			ChangeMoney:         office.Money - officeBak.Money,
			ChangeMoneyCold:     office.MoneyCold - officeBak.MoneyCold,
			ChangeTargetLoad:    office.TargetLoad - officeBak.TargetLoad,
			NewBusinessMoney:    office.BusinessMoney,
			NewMoney:            office.Money,
			NewMoneyCold:        office.MoneyCold,
			NewTargetLoad:       office.TargetLoad,
			CreateDate:          XDate{Time: time.Now()},
			Remark:              "[直接修改] : " + office.Remark,
		}
		if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
			return tErr
		}

		if tErr := tx.Model(&Office{}).Where("id", office.ID).Updates(maps).Error; tErr != nil {
			return tErr
		}

		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectOffices(officeQuery *Office, xForms *XForms) (offices []Office, code int) {
	tx := db.Where("is_delete = ?", false)
	if officeQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+officeQuery.Name+"%")
	}

	err = tx.Find(&offices).Count(&xForms.Total).
		Preload("Role").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&offices).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return offices, msg.SUCCESS
}

func SelectNotPaymentForTopList() (offices1 []Office, offices2 []Office, productTypes []ProductType) {
	db.Raw("SELECT office_id id,sum(total_amount - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS FALSE AND pay_type = 1 AND STATUS > 1 GROUP BY office_id").Scan(&offices1)
	db.Raw("SELECT office_id id,sum(pre_deposit_record - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS TRUE AND STATUS > 1 AND pre_deposit_record > payment_total_amount GROUP BY office_id").Scan(&offices2)
	db.Raw("SELECT product_type.id,product_type.`name` ,sum(payment.money) as push_money_percentages FROM payment LEFT JOIN task ON payment.task_id = task.id LEFT JOIN product ON task.product_id = product.id LEFT JOIN product_type ON product.type_id = product_type.id WHERE payment.task_id IS NOT NULL GROUP BY product_type.id").Scan(&productTypes)
	return
}
