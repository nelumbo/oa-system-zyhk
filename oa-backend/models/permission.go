package models

type Permission struct {
	// UID  string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`
	// Text string `gorm:"type:varchar(99);comment:文本" json:"text"`
	// No   string `gorm:"type:varchar(9);comment:序号" json:"no"`

	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name     string `gorm:"type:varchar(50);comment:名称" json:"name"`
	Value    string `gorm:"type:varchar(50);comment:值" json:"value"`
}
