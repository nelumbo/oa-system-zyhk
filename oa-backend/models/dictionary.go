package models

import "oa-backend/utils/msg"

type DictionaryType struct {
	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name     string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Text     string `gorm:"type:varchar(50);comment:文本;not null" json:"text"`

	Dictionaries []Dictionary `json:"dictionaries"`
}

type Dictionary struct {
	// UID               string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// DictionaryTypeUID string `gorm:"type:varchar(32);comment:类型ID;default:(-)" json:"dictionaryTypeUID"`

	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	TypeID   int    `gorm:"type:int;comment:类型ID;default:(-)" json:"dictionaryTypeID"`
	Text     string `gorm:"type:varchar(50);comment:文本" json:"text"`
	Value    int    `gorm:"type:int;comment:值;default:(-)" json:"value"`

	DictionaryType DictionaryType `gorm:"foreignKey:TypeID" json:"dictionaryType"`
}

func SelectDictionaries(dictionaryQuery *Dictionary) (dictionaries []Dictionary, code int) {
	tx := db.Joins("DictionaryType").Where("dictionary.s_delete = ?", false)
	if dictionaryQuery.DictionaryType.Name != "" {
		tx = tx.Where("DictionaryType.name = ?", dictionaryQuery.DictionaryType.Name)
	}
	err = tx.Find(&dictionaries).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return dictionaries, msg.SUCCESS
}
