package models

type Permission struct {
	// UID  string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`
	// Text string `gorm:"type:varchar(99);comment:文本" json:"text"`
	// No   string `gorm:"type:varchar(9);comment:序号" json:"no"`

	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `gorm:"type:varchar(50);comment:名称" json:"name"`
	Value string `gorm:"type:varchar(50);comment:值" json:"value"`
	UrlID int    `gorm:"type:int;comment:urlID" json:"urlID"`
}

type Url struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Title string `gorm:"type:varchar(50);comment:title" json:"title"`
	Icon  string `gorm:"type:varchar(50);comment:icon" json:"icon"`
	Url   string `gorm:"type:varchar(50);comment:url" json:"url"`
	No    int    `gorm:"type:int;comment:no" json:"no"`
}

func SelectAllPermission(employee *Employee) (code int) {
	db.Raw("SELECT distinct permission_id FROM role_permission WHERE role_id IN (SELECT role_id AS id FROM employee_role WHERE employee_id = ? UNION SELECT role_id AS id FROM office WHERE id = ?)", employee.ID, employee.OfficeID).Scan(&employee.Pids)
	db.Raw("SELECT distinct permission.url_id FROM (SELECT distinct permission_id FROM role_permission WHERE role_id IN (SELECT role_id AS id FROM employee_role WHERE employee_id = ? UNION SELECT role_id AS id FROM office WHERE id = ?)) temp LEFT JOIN permission ON temp.permission_id = permission.id WHERE permission.url_id IS NOT NULL ", employee.ID, employee.OfficeID).Scan(&employee.Urls)
	return
}
