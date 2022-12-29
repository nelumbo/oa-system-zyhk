package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"ID"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GeneralInsert(model interface{}) int {
	err = db.Create(model).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func GeneralDelete(model interface{}, id int) int {
	err = db.Model(model).Where("id = ?", id).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func GeneralUpdate(model interface{}, id int, maps map[string]interface{}) int {
	err = db.Model(model).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func GeneralSelect(model interface{}, id int, maps map[string]interface{}) int {
	tx := db.Where("is_delete", false).Where("id = ?", id)
	if len(maps) > 0 {
		tx = tx.Where(maps)
	}
	err = tx.First(model).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return msg.FAIL
		} else {
			return msg.ERROR
		}
	}
	return msg.SUCCESS
}

func GeneralSelectAll(model interface{}, maps map[string]interface{}) int {
	tx := db.Where("is_delete", false)
	if len(maps) > 0 {
		err = tx.Where(maps).Find(model).Error
	} else {
		err = tx.Find(model).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func GeneralSelectLimit(model interface{}, xForms *XForms, maps map[string]interface{}) int {
	tx := db.Where("1 = ?", 1)
	if len(maps) > 0 {
		tx = tx.Where(maps)
	}
	err = tx.Find(model).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(model).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return msg.ERROR
	}

	xForms.Data = model
	return msg.SUCCESS
}
