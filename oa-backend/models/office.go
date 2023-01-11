package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Office struct {
	// UID           string  `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`

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

	Role Role `gorm:"foreignKey:RoleID" json:"role"`

	FinalPercentages float64 `gorm:"-" json:"finalPercentages"`
	NotPayment       float64 `gorm:"-" json:"notPayment"`
}

func UpdateOffice(office *Office, maps map[string]interface{}) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := tx.Model(&Office{}).Where("id", office.ID).Updates(maps).Error; tErr != nil {
			return tErr
		}
		//TODO日志
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

func SelectNotPaymentForTopList() (offices1 []Office, offices2 []Office) {
	db.Raw("SELECT office_id id,sum(total_amount - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS FALSE AND STATUS > 1 GROUP BY office_id").Scan(&offices1)
	db.Raw("SELECT office_id id,sum(pre_deposit_record - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS TRUE AND STATUS > 1 AND pre_deposit_record > payment_total_amount GROUP BY office_id").Scan(&offices2)
	return
}
