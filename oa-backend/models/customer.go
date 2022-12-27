package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type CustomerCompany struct {
	// UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// RegionUID string `gorm:"type:varchar(32);comment:省份UID;default:(-)" json:"regionUID"`

	ID       int    `gorm:"primary_key" json:"id"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	RegionID int    `gorm:"type:int;comment:省份ID;default:(-)" json:"regionID"`
	Name     string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Address  string `gorm:"type:varchar(100);comment:地址" json:"address"`

	Region    Region     `gorm:"foreignKey:RegionID" json:"region"`
	Customers []Customer `json:"customers"`
}

type Customer struct {
	// UID           string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// CompanyUID    string `gorm:"type:varchar(32);comment:公司UID;default:(-)" json:"companyUID"`
	// Status        int    `gorm:"type:int;comment:状态(0:未审核,1:通过审核)" json:"status"`

	ID                int    `gorm:"primary_key" json:"id"`
	IsDelete          bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	CustomerCompanyID int    `gorm:"type:int;comment:公司ID;default:(-)" json:"customerCompanyID"`
	Name              string `gorm:"type:varchar(50);comment:姓名;not null" json:"name"`
	ResearchGroup     string `gorm:"type:varchar(100);comment:课题组" json:"researchGroup"`
	Phone             string `gorm:"type:varchar(50);comment:电话" json:"phone"`
	WechatID          string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email             string `gorm:"type:varchar(50);comment:邮箱" json:"email"`

	CustomerCompany CustomerCompany `gorm:"foreignKey:CustomerCompanyID" json:"customerCompany"`
	Contracts       []Contract      `json:"contracts"`
}

func DeleteCustomerCompany(id int) (code int) {
	var customerCompany CustomerCompany
	customerCompany, code = SelectCustomerCompany(id)
	if code != msg.SUCCESS || len(customerCompany.Customers) > 0 {
		return msg.FAIL
	}
	code = GeneralDelete(&customerCompany, id)
	return code
}

func SelectCustomerCompany(id int) (customerCompany CustomerCompany, code int) {
	db.Preload("Customers").Where("id = ? AND is_delete = ?", id, false).First(&customerCompany)
	if customerCompany.ID == 0 {
		return CustomerCompany{}, msg.FAIL
	}
	return customerCompany, msg.SUCCESS
}

func SelectCustomerCompanys(customerCompanyQuery *CustomerCompany, xForms *XForms) (customerCompanys []CustomerCompany, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	if customerCompanyQuery.RegionID != 0 {
		maps["region_id"] = customerCompanyQuery.RegionID
	}

	tx := db.Where(maps)
	if customerCompanyQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+customerCompanyQuery.Name+"%")
	}
	err = tx.Find(&customerCompanys).Count(&xForms.Total).
		Preload("Region").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&customerCompanys).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return customerCompanys, msg.SUCCESS
}

func DeleteCustomer(id int) (code int) {
	var customer Customer
	customer, code = SelectCustomer(id)
	if code != msg.SUCCESS || len(customer.Contracts) > 0 {
		return msg.FAIL
	}
	code = GeneralDelete(&customer, id)
	return code
}

func SelectCustomer(id int) (customer Customer, code int) {
	db.Preload("Contracts").Where("id = ? AND is_delete = ?", id, false).First(&customer)
	if customer.ID == 0 {
		return Customer{}, msg.FAIL
	}
	return customer, msg.SUCCESS
}

func SelectCustomers(customerQuery *Customer, xForms *XForms) (customers []Customer, code int) {
	var maps = make(map[string]interface{})
	maps["customer.is_delete"] = false
	if customerQuery.CustomerCompany.RegionID != 0 {
		maps["CustomerCompany.region_id"] = customerQuery.CustomerCompany.RegionID
	}

	tx := db.Joins("CustomerCompany").Where(maps)
	if customerQuery.CustomerCompany.Name != "" {
		tx = tx.Where("CustomerCompany.name LIKE ?", "%"+customerQuery.CustomerCompany.Name+"%")
	}
	if customerQuery.Name != "" {
		tx = tx.Where("customer.name LIKE ?", "%"+customerQuery.Name+"%")
	}
	if customerQuery.ResearchGroup != "" {
		tx = tx.Where("customer.research_group LIKE ?", "%"+customerQuery.ResearchGroup+"%")
	}
	err = tx.Find(&customers).Count(&xForms.Total).
		Preload("CustomerCompany.Region").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&customers).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return customers, msg.SUCCESS
}
