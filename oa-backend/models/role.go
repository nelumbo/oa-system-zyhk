package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Role struct {
	// UID  string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`

	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name     string `gorm:"type:varchar(50);comment:名称" json:"name"`

	Permissions []Permission `gorm:"many2many:role_permission;foreignKey:ID;References:ID" json:"permissions"`
	Employees   []Employee   `gorm:"many2many:employee_role"`
}

func DeleteRole(id int) (code int) {
	var role Role
	role, code = SelectRole(id)
	if code != msg.SUCCESS {
		return msg.FAIL
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := db.Model(&Office{}).Where("role_id", role.ID).Update("role_id", nil).Error; tErr != nil {
			return tErr
		}
		if tErr := db.Model(&role).Association("Permissions").Clear(); tErr != nil {
			return tErr
		}
		if tErr := db.Model(&role).Association("Employees").Clear(); tErr != nil {
			return tErr
		}
		if tErr := db.Model(&role).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateRole(role *Role) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := tx.Model(&role).Update("name", role.Name).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&role).Association("Permissions").Replace(role.Permissions); tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectRole(id int) (role Role, code int) {
	err = db.Preload("Permissions").Where("is_delete = ?", false).First(&role, id).Error
	if err != nil {
		return Role{}, msg.ERROR
	}
	return role, msg.SUCCESS
}

func SelectRoles(roleQuery *Role, xForms *XForms) (roles []Role, code int) {
	tx := db.Where("is_delete = ?", false)
	if roleQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+roleQuery.Name+"%")
	}
	err = tx.Find(&roles).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&roles).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return roles, msg.SUCCESS
}
