package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Supplier struct {
	// UID  string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`

	ID          int    `gorm:"primary_key" json:"id"`
	IsDelete    bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name        string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Address     string `gorm:"type:varchar(100);comment:地址" json:"address"`
	Web         string `gorm:"type:varchar(50);comment:网站" json:"web"`
	Linkman     string `gorm:"type:varchar(50);comment:联系人" json:"linkman"`
	Phone       string `gorm:"type:varchar(50);comment:联系电话" json:"phone"`
	WechatID    string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email       string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	Description string `gorm:"type:varchar(300);comment:主营产品概述" json:"description"`
	Remark      string `gorm:"type:varchar(300);comment:备注" json:"remark"`
	InvoiceType string `gorm:"type:varchar(50);comment:开票种类" json:"invoiceType"`
}

type Product struct {
	// UID  string `gorm:"type:varchar(32);comment:唯一标识" json:"UID"`
	// SupplierUID string `gorm:"type:varchar(32);comment:供应商UID;default:(-)" json:"supplierUID"`

	ID            int    `gorm:"primary_key" json:"id"`
	IsDelete      bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name          string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Version       string `gorm:"type:varchar(50);comment:型号;not null" json:"version"`
	Brand         string `gorm:"type:varchar(50);comment:品牌" json:"brand"`
	Specification string `gorm:"type:varchar(100);comment:规格" json:"specification"`
	Number        int    `gorm:"type:int;comment:可售数量(库存数量-订单锁定但未出库的数量)" json:"number"`
	NumberCount   int    `gorm:"type:int;comment:库存数量" json:"numberCount"`
	CallNumber    int    `gorm:"type:int;comment:报警数量" json:"callNumber"`
	Unit          string `gorm:"type:varchar(50);comment:单位" json:"unit"`
	DeliveryCycle string `gorm:"type:varchar(50);comment:供货周期" json:"deliveryCycle"`
	Remark        string `gorm:"type:varchar(300);comment:备注" json:"remark"`
	IsFree        bool   `gorm:"type:boolean;comment:是否为小零配件" json:"isFree"`
	TypeID        int    `gorm:"type:int;comment:类型ID;default:(-)" json:"typeID"`
	AttributeID   int    `gorm:"type:int;comment:属性ID;default:(-)" json:"attributeID"`

	Type      ProductType      `gorm:"foreignKey:TypeID" json:"type"`
	Attribute ProductAttribute `gorm:"foreignKey:AttributeID" json:"attribute"`
	Suppliers []Supplier       `gorm:"many2many:product_supplier;foreignKey:ID;references:ID" json:"suppliers"`
}

type ProductAttribute struct {
	// TypeUID          string  `gorm:"type:varchar(32);comment:类型UID;default:(-)" json:"typeUID"`

	ID               int     `gorm:"primary_key" json:"id"`
	IsDelete         bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	PurchasePrice    float64 `gorm:"type:decimal(20,6);comment:采购价格(元)" json:"purchasePrice"`
	PurchasePriceUSD float64 `gorm:"type:decimal(20,6);comment:采购价格(美元)" json:"purchasePriceUSD"`
	StandardPrice    float64 `gorm:"type:decimal(20,6);comment:标准价格(元)" json:"standardPrice"`
	StandardPriceUSD float64 `gorm:"type:decimal(20,6);comment:标准价格(美元)" json:"standardPriceUSD"`
}

type ProductType struct {
	// UID                        string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`

	ID                         int     `gorm:"primary_key" json:"id"`
	IsDelete                   bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name                       string  `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	PushMoneyPercentages       float64 `gorm:"type:decimal(20,6);comment:标准提成百分比" json:"pushMoneyPercentages"`
	MinPushMoneyPercentages    float64 `gorm:"type:decimal(20,6);comment:标准提成百分比" json:"minPushMoneyPercentages"`
	PushMoneyPercentagesUp     float64 `gorm:"type:decimal(20,6);comment:提成上涨百分比" json:"pushMoneyPercentagesUp"`
	PushMoneyPercentagesDown   float64 `gorm:"type:decimal(20,6);comment:提成下降百分比" json:"pushMoneyPercentagesDown"`
	BusinessMoneyPercentages   float64 `gorm:"type:decimal(20,6);comment:标准业务费用百分比" json:"businessMoneyPercentages"`
	BusinessMoneyPercentagesUp float64 `gorm:"type:decimal(20,6);comment:业务费用上涨百分比" json:"businessMoneyPercentagesUp"`
	Type                       int     `gorm:"type:int;comment:类型(1:原味 2:渠道 3:自研)" json:"type"`
	IsTaskLoad                 bool    `gorm:"type:boolean;comment:是否计算任务量" json:"isTaskLoad"`
	IsDirectOut                bool    `gorm:"type:boolean;comment:是否可以直接出库" json:"isDirectOut"`
}

func SelectSuppliers(supplierQuery *Supplier, xForms *XForms) (suppliers []Supplier, code int) {
	tx := db.Where("is_delete = ?", false)

	if supplierQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+supplierQuery.Name+"%")
	}
	if supplierQuery.Linkman != "" {
		tx = tx.Where("linkman LIKE ?", "%"+supplierQuery.Linkman+"%")
	}
	if supplierQuery.Phone != "" {
		tx = tx.Where("phone LIKE ?", "%"+supplierQuery.Phone+"%")
	}

	err = tx.Find(&suppliers).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&suppliers).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return suppliers, msg.SUCCESS
}

func InsertProduct(product *Product) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := db.Create(&product.Attribute).Error; tErr != nil {
			return tErr
		}
		product.AttributeID = product.Attribute.ID
		if tErr := db.Create(&product).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateProduct(product *Product) (code int) {
	var maps = make(map[string]interface{})
	maps["type_id"] = product.TypeID
	maps["name"] = product.Name
	maps["brand"] = product.Brand
	maps["specification"] = product.Specification
	maps["unit"] = product.Unit
	maps["delivery_cycle"] = product.DeliveryCycle
	maps["remark"] = product.Remark
	maps["is_free"] = product.IsFree

	err = db.Transaction(func(tx *gorm.DB) error {
		if tErr := tx.Model(&Product{ID: product.ID}).Updates(maps).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&product).Association("Suppliers").Replace(product.Suppliers); tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateProductAttribute(product *Product) (code int) {
	var productBak Product
	productBak, code = SelectProduct(product.ID)
	if code == msg.SUCCESS {
		if productBak.Attribute.PurchasePrice != product.Attribute.PurchasePrice ||
			productBak.Attribute.PurchasePriceUSD != product.Attribute.PurchasePriceUSD ||
			productBak.Attribute.StandardPrice != product.Attribute.StandardPrice ||
			productBak.Attribute.StandardPriceUSD != product.Attribute.StandardPriceUSD {
			var productAttribute ProductAttribute
			productAttribute.PurchasePrice = product.Attribute.PurchasePrice
			productAttribute.PurchasePriceUSD = product.Attribute.PurchasePriceUSD
			productAttribute.StandardPrice = product.Attribute.StandardPrice
			productAttribute.StandardPriceUSD = product.Attribute.StandardPriceUSD

			err = db.Transaction(func(tx *gorm.DB) error {
				if tErr := tx.Create(&productAttribute).Error; tErr != nil {
					return tErr
				}
				if tErr := tx.Model(&Product{ID: product.ID}).Update("attribute_id", productAttribute.ID).Error; tErr != nil {
					return tErr
				}
				return nil
			})
			if err != nil {
				return msg.ERROR
			}
			return msg.SUCCESS
		}
	}
	return code
}

func SelectProduct(id int) (product Product, code int) {
	db.Preload("Type").Preload("Attribute").Preload("Suppliers").
		Where("is_delete = ?", false).
		First(&product, id)
	if product.ID == 0 {
		return Product{}, msg.FAIL
	}
	return product, msg.SUCCESS
}

func SelectProducts(productQuery *Product, xForms *XForms) (products []Product, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	if productQuery.TypeID > 0 {
		maps["product.type_id"] = productQuery.TypeID
	}

	tx := db.Where(maps)

	if productQuery.Name != "" {
		tx = tx.Where("product.name LIKE ?", "%"+productQuery.Name+"%")
	}
	if productQuery.Specification != "" {
		tx = tx.Where("product.specification LIKE ?", "%"+productQuery.Specification+"%")
	}

	err = tx.Find(&products).Count(&xForms.Total).
		Preload("Type").Preload("Attribute").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&products).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return products, msg.SUCCESS
}

func SelectProductTypes(productTypeQuery *ProductType, xForms *XForms) (productTypes []ProductType, code int) {
	tx := db.Where("is_delete = ?", false)
	if productTypeQuery.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+productTypeQuery.Name+"%")
	}
	err = tx.Find(&productTypes).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&productTypes).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return productTypes, msg.SUCCESS
}
