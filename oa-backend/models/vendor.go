package models

import "oa-backend/utils/msg"

type Vendor struct {
	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name     string `gorm:"type:varchar(50);comment:名称" json:"name"`
}

func SelectVendors(vendorQuery *Vendor, xForms *XForms) (vendors []Vendor, code int) {
	tx := db.Where("is_delete = ?", false)
	if vendorQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+vendorQuery.Name+"%")
	}
	err = tx.Find(&vendors).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&vendors).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return vendors, msg.SUCCESS
}
