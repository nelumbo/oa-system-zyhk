package models

import "oa-backend/utils/msg"

type Region struct {
	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name     string `gorm:"type:varchar(50);comment:名称" json:"name"`
}

func SelectRegions(regionQuery *Region, xForms *XForms) (regions []Region, code int) {
	tx := db.Where("is_delete = ?", false)
	if regionQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+regionQuery.Name+"%")
	}
	err = tx.Find(&regions).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&regions).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return regions, msg.SUCCESS
}
