package models

import (
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID          int    `gorm:"primary_key" json:"id"`
	IsDelete    bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name        string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Address     string `gorm:"type:varchar(100);comment:地址" json:"address"`
	Web         string `gorm:"type:varchar(100);comment:网站" json:"web"`
	Linkman     string `gorm:"type:varchar(50);comment:联系人" json:"linkman"`
	Phone       string `gorm:"type:varchar(50);comment:联系电话" json:"phone"`
	WechatID    string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email       string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	Description string `gorm:"type:varchar(300);comment:主营产品概述" json:"description"`
	Remark      string `gorm:"type:varchar(300);comment:备注" json:"remark"`
	InvoiceType string `gorm:"type:varchar(50);comment:开票种类" json:"invoiceType"`
}

type Product struct {
	ID            int     `gorm:"primary_key" json:"id"`
	IsDelete      bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	Name          string  `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Version       string  `gorm:"type:varchar(50);comment:型号" json:"version"`
	Brand         string  `gorm:"type:varchar(50);comment:品牌" json:"brand"`
	Specification string  `gorm:"type:varchar(100);comment:规格" json:"specification"`
	Number        int     `gorm:"type:int;comment:库存数量" json:"number"`
	CallNumber    int     `gorm:"type:int;comment:报警数量" json:"callNumber"`
	Unit          string  `gorm:"type:varchar(50);comment:单位" json:"unit"`
	DeliveryCycle string  `gorm:"type:varchar(50);comment:供货周期" json:"deliveryCycle"`
	Remark        string  `gorm:"type:varchar(300);comment:备注" json:"remark"`
	IsFree        bool    `gorm:"type:boolean;comment:是否为小零配件" json:"isFree"`
	TypeID        int     `gorm:"type:int;comment:类型ID;default:(-)" json:"typeID"`
	PurchasePrice float64 `gorm:"type:decimal(20,6);comment:采购价格(元)" json:"purchasePrice"`
	AttributeID   int     `gorm:"type:int;comment:属性ID;default:(-)" json:"attributeID"`

	Type      ProductType      `gorm:"foreignKey:TypeID" json:"type"`
	Attribute ProductAttribute `gorm:"foreignKey:AttributeID" json:"attribute"`
	Suppliers []Supplier       `gorm:"many2many:product_supplier;foreignKey:ID;references:ID" json:"suppliers"`
}

type ProductAttribute struct {
	ID               int     `gorm:"primary_key" json:"id"`
	IsDelete         bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	StandardPrice    float64 `gorm:"type:decimal(20,6);comment:标准价格(元)" json:"standardPrice"`
	StandardPriceUSD float64 `gorm:"type:decimal(20,6);comment:标准价格(美元)" json:"standardPriceUSD"`
}

type ProductType struct {
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
}

type ProductTrial struct {
	ID            int    `gorm:"primary_key" json:"id"`
	IsDelete      bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ProductID     int    `gorm:"type:int;comment:产品ID" json:"productID"`
	Number        int    `gorm:"type:int;comment:数量" json:"number"`
	Status        int    `gorm:"type:int;comment:状态(-1:驳回 1:待审批 2:待发货 3:待归还 4:待确认 5:完成)" json:"status"`
	EmployeeID    int    `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	AuditorID     int    `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	ShipmentID    int    `gorm:"type:int;comment:物流人员ID;default:(-)" json:"shipmentID"`
	InventoryID   int    `gorm:"type:int;comment:仓库负责人ID;default:(-)" json:"inventoryID"`
	FinalID       int    `gorm:"type:int;comment:确认人员ID;default:(-)" json:"finalID"`
	CreateDate    XDate  `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate     XDate  `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	ShipmentDate  XDate  `gorm:"type:date;comment:发货日期;default:(-)" json:"shipmentDate"`
	InventoryDate XDate  `gorm:"type:date;comment:归还日期;default:(-)" json:"inventoryDate"`
	FinalDate     XDate  `gorm:"type:date;comment:确认日期;default:(-)" json:"finalDate"`
	Remark1       string `gorm:"type:varchar(300);comment:备注1" json:"remark1"`
	Remark2       string `gorm:"type:varchar(300);comment:备注2" json:"remark2"`
	Remark3       string `gorm:"type:varchar(300);comment:备注3" json:"remark3"`
	Remark4       string `gorm:"type:varchar(300);comment:备注4" json:"remark4"`
	Remark5       string `gorm:"type:varchar(300);comment:备注5" json:"remark5"`

	Product   Product  `gorm:"foreignKey:ProductID" json:"product"`
	Employee  Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Auditor   Employee `gorm:"foreignKey:AuditorID" json:"auditor"`
	Shipment  Employee `gorm:"foreignKey:ShipmentID" json:"shipment"`
	Inventory Employee `gorm:"foreignKey:InventoryID" json:"inventory"`
	Final     Employee `gorm:"foreignKey:FinalID" json:"final"`
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
	maps["version"] = product.Version
	maps["specification"] = product.Specification
	maps["unit"] = product.Unit
	maps["delivery_cycle"] = product.DeliveryCycle
	maps["remark"] = product.Remark
	maps["is_free"] = product.IsFree
	maps["call_number"] = product.CallNumber

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
		err = db.Transaction(func(tx *gorm.DB) error {

			if productBak.Attribute.StandardPrice != product.Attribute.StandardPrice ||
				productBak.Attribute.StandardPriceUSD != product.Attribute.StandardPriceUSD {
				var productAttribute ProductAttribute
				productAttribute.StandardPrice = product.Attribute.StandardPrice
				productAttribute.StandardPriceUSD = product.Attribute.StandardPriceUSD
				if tErr := tx.Create(&productAttribute).Error; tErr != nil {
					return tErr
				}
				if tErr := tx.Model(&Product{}).Where("id", product.ID).Updates(map[string]interface{}{"purchase_price": product.PurchasePrice, "attribute_id": productAttribute.ID}).Error; tErr != nil {
					return tErr
				}
			} else {
				if tErr := tx.Model(&Product{}).Where("id", product.ID).Update("purchase_price", product.PurchasePrice).Error; tErr != nil {
					return tErr
				}
			}
			return nil
		})
		if err != nil {
			code = msg.ERROR
		}
		code = msg.SUCCESS
	} else {
		code = msg.ERROR
	}
	return
}

func UpdateProductNumber(product *Product) (code int) {
	//TODO日志
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := tx.Model(&Product{}).Where("id", product.ID).Update("number", product.Number).Error; tErr != nil {
			return tErr
		}

		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
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

func NextProductTrial(productTrial *ProductTrial, productTrialBak *ProductTrial, employeeID int) (code int) {
	var maps = make(map[string]interface{})

	if productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_APPROVE {
		if productTrial.Status == magic.PRODUCT_TRIAL_STATUS_NO_SEND {
			err = db.Transaction(func(tx *gorm.DB) error {

				if tErr := tx.Model(&Product{}).Where("id", productTrialBak.ProductID).Update("number", gorm.Expr("number - ?", productTrialBak.Number)).Error; tErr != nil {
					return tErr
				}

				if tErr := tx.Model(&ProductTrial{}).Where("id", productTrialBak.ID).Updates(map[string]interface{}{"status": magic.PRODUCT_TRIAL_STATUS_NO_SEND, "auditor_id": employeeID, "audit_date": XDate{Time: time.Now()}}).Error; tErr != nil {
					return tErr
				}

				return nil
			})

			if err != nil {
				return msg.ERROR
			}
			return msg.SUCCESS

		} else if productTrial.Status == magic.PRODUCT_TRIAL_STATUS_REJECT {
			maps["status"] = magic.PRODUCT_TRIAL_STATUS_REJECT
			maps["auditor_id"] = employeeID
			maps["audit_date"] = XDate{Time: time.Now()}
		}
	} else if productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_SEND && productTrial.Status == magic.PRODUCT_TRIAL_STATUS_NO_BACK {
		maps["status"] = magic.PRODUCT_TRIAL_STATUS_NO_BACK
		maps["shipment_id"] = employeeID
		maps["shipment_date"] = XDate{Time: time.Now()}
		maps["remark3"] = productTrial.Remark3
	} else if productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_BACK && productTrial.Status == magic.PRODUCT_TRIAL_STATUS_NO_FINAL {
		maps["status"] = magic.PRODUCT_TRIAL_STATUS_NO_FINAL
		maps["inventory_id"] = employeeID
		maps["inventory_date"] = XDate{Time: time.Now()}
		maps["remark4"] = productTrial.Remark4
	} else if productTrialBak.Status == magic.PRODUCT_TRIAL_STATUS_NO_FINAL && productTrial.Status == magic.PRODUCT_TRIAL_STATUS_FINAL {
		maps["status"] = magic.PRODUCT_TRIAL_STATUS_FINAL
		maps["final_id"] = employeeID
		maps["final_date"] = XDate{Time: time.Now()}
		maps["remark5"] = productTrial.Remark5
	} else {
		return msg.FAIL
	}

	err = db.Model(&ProductTrial{}).Where("id", productTrialBak.ID).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectProductTrials(productTrialQuery *ProductTrial, xForms *XForms) (productTrials []ProductTrial, code int) {
	var maps = make(map[string]interface{})
	maps["product_trial.is_delete"] = false
	if productTrialQuery.Status != 0 {
		maps["product_trial.status"] = productTrialQuery.Status
	}
	if productTrialQuery.Employee.OfficeID != 0 {
		maps["Employee.office_id"] = productTrialQuery.Employee.OfficeID
	}

	tx := db.Where(maps).Joins("Employee")
	err = tx.Find(&productTrials).Count(&xForms.Total).
		Preload("Product").Preload("Employee.Office").Preload("Auditor").
		Preload("Shipment").Preload("Inventory").Preload("Final").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Order("id desc").Find(&productTrials).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return productTrials, msg.SUCCESS
}
