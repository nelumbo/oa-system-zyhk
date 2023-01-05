package models

import (
	"oa-backend/utils/msg"

	"gorm.io/gorm"
)

type Purchasing struct {
	ID                 int     `gorm:"primary_key" json:"id"`
	IsDelete           bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ContractID         int     `gorm:"type:int;comment:合同ID" json:"contractID"`
	TaskID             int     `gorm:"type:int;comment:任务ID" json:"taskID"`
	EmployeeID         int     `gorm:"type:int;comment:创建人ID;default:(-)" json:"employeeID"`
	No                 string  `gorm:"type:varchar(100);comment:合同编号" json:"no"`
	Type               int     `gorm:"type:int;comment:类型(1:合同附属 2:备货装置任务 3:库存配件任务)" json:"type"`
	ProductID          int     `gorm:"type:int;comment:产品ID" json:"productID"`
	ProductAttributeID int     `gorm:"type:int;comment:产品属性ID" json:"productAttributeID"`
	Price              float64 `gorm:"type:decimal(20,6);comment:采购价格" json:"price"`
	Number             int     `gorm:"type:int;comment:申请数量" json:"number"`
	RealNumber         int     `gorm:"type:int;comment:实际数量" json:"realNumber"`
	TotalPrice         float64 `gorm:"type:decimal(20,6);comment:总价" json:"totalPrice"`
	Status             int     `gorm:"type:int;comment:状态(-1:驳回 0:暂存 1:待采购确认 2:待审批 3:未完成 4:已完成)" json:"status"`
	ProductStatus      int     `gorm:"type:int;comment:产品状态(1:待收货 2:已收货)" json:"productStatus"`
	PayStatus          int     `gorm:"type:int;comment:付款状态(1:待付款 2:已付款)" json:"collectionStatus"`
	InvoiceStatus      int     `gorm:"type:int;comment:发票状态(1:发票未收到 2:发票已收到)" json:"invoiceStatus"`
	AuditorID          int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	PurchaseManID      int     `gorm:"type:int;comment:采购负责人ID;default:(-)" json:"purchaseManID"`
	InventoryManID     int     `gorm:"type:int;comment:仓库负责人ID;default:(-)" json:"inventoryManID"`
	PayManID           int     `gorm:"type:int;comment:财务负责人ID;default:(-)" json:"payManID"`
	InvoiceManID       int     `gorm:"type:int;comment:发票负责人ID;default:(-)" json:"invoiceManID"`

	CreateDate    XDate `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EndDate       XDate `gorm:"type:date;comment:期望到达日期;default:(-)" json:"endDate"`
	RealEndDate   XDate `gorm:"type:date;comment:到货日期;default:(-)" json:"realEndDate"`
	AuditDate     XDate `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	PayDate       XDate `gorm:"type:date;comment:付款日期;default:(-)" json:"payDate"`
	PayCreateDate XDate `gorm:"type:date;comment:财务提交日期;default:(-)" json:"payCreateDate"`
	InvoiceDate   XDate `gorm:"type:date;comment:发票到达日期;default:(-)" json:"invoiceDate"`

	IsPass bool `gorm:"-" json:"isPass"`
}

func InsertPurchasing(purchasing *Purchasing, maps map[string]interface{}) (code int) {

	err = db.Model(&Purchasing{}).
		Where("contractID = ? AND taskID = ? AND employeeID = ?", purchasing.ContractID, purchasing.TaskID, purchasing.EmployeeID).
		Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func ApprovePurchasing(purchasing *Purchasing, maps map[string]interface{}) (code int) {
	if purchasing.IsPass {
		err = db.Transaction(func(tx *gorm.DB) error {

			if tErr := tx.Model(&Purchasing{}).Where("id", purchasing.ID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tx.Model(&ProductAttribute{}).Where("id", purchasing.ProductAttributeID).Update("purchase_price", purchasing.Price).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		err = db.Model(&Purchasing{}).Where("id", purchasing.ID).Updates(maps).Error
	}
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}
