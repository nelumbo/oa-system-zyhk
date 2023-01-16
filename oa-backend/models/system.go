package models

var Ice int

func InitSystem() {
	var iceBak System
	db.First(&iceBak, "text = ?", "ice")
	Ice = iceBak.Value
}

type System struct {
	ID    uint   `gorm:"primary_key" json:"ID"`
	Text  string `gorm:"type:string;comment:Text" json:"text"`
	Value int    `gorm:"type:int;comment:Value" json:"value"`
}
