package models

import (
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Purchasing struct {
	ID             int     `gorm:"primary_key" json:"id"`
	IsDelete       bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	IsSecond       bool    `gorm:"type:boolean;comment:合同第二次提交" json:"isSecond"`
	IsSecondFinal  bool    `gorm:"type:boolean;comment:合同第二次提交扣除补贴" json:"isSecondFinal"`
	ContractID     int     `gorm:"type:int;comment:合同ID;default:(-)" json:"contractID"`
	TaskID         int     `gorm:"type:int;comment:任务ID;default:(-)" json:"taskID"`
	EmployeeID     int     `gorm:"type:int;comment:创建人ID;default:(-)" json:"employeeID"`
	No             string  `gorm:"type:varchar(100);comment:合同编号" json:"no"`
	Type           int     `gorm:"type:int;comment:类型(1:合同附属 2:备货装置任务 3:库存配件任务)" json:"type"`
	ProductID      int     `gorm:"type:int;comment:产品ID" json:"productID"`
	Price          float64 `gorm:"type:decimal(20,6);comment:采购价格" json:"price"`
	Number         int     `gorm:"type:int;comment:申请数量" json:"number"`
	RealNumber     int     `gorm:"type:int;comment:实际数量" json:"realNumber"`
	TotalPrice     float64 `gorm:"type:decimal(20,6);comment:总价" json:"totalPrice"`
	Status         int     `gorm:"type:int;comment:状态(-1:驳回 0:暂存 1:待采购确认 2:待审批 3:未完成 4:已完成)" json:"status"`
	ProductStatus  int     `gorm:"type:int;comment:产品状态(1:待收货 2:已收货)" json:"productStatus"`
	PayStatus      int     `gorm:"type:int;comment:付款状态(1:待付款 2:已付款)" json:"payStatus"`
	InvoiceStatus  int     `gorm:"type:int;comment:发票状态(1:发票未收到 2:发票已收到)" json:"invoiceStatus"`
	AuditorID      int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	PurchaseManID  int     `gorm:"type:int;comment:采购负责人ID;default:(-)" json:"purchaseManID"`
	InventoryManID int     `gorm:"type:int;comment:仓库负责人ID;default:(-)" json:"inventoryManID"`
	PayManID       int     `gorm:"type:int;comment:财务负责人ID;default:(-)" json:"payManID"`
	InvoiceManID   int     `gorm:"type:int;comment:发票负责人ID;default:(-)" json:"invoiceManID"`
	ProductRemark  string  `gorm:"type:varchar(300);comment:仓库备注" json:"productRemark"`
	PayRemark      string  `gorm:"type:varchar(300);comment:财务备注" json:"payRemark"`
	InvoiceRemark  string  `gorm:"type:varchar(300);comment:发票备注" json:"invoiceRemark"`

	CreateDate    XDate `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	EndDate       XDate `gorm:"type:date;comment:期望到达日期;default:(-)" json:"endDate"`
	RealEndDate   XDate `gorm:"type:date;comment:到货日期;default:(-)" json:"realEndDate"`
	AuditDate     XDate `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`
	PayDate       XDate `gorm:"type:date;comment:付款日期;default:(-)" json:"payDate"`
	PayCreateDate XDate `gorm:"type:date;comment:财务提交日期;default:(-)" json:"payCreateDate"`
	InvoiceDate   XDate `gorm:"type:date;comment:发票到达日期;default:(-)" json:"invoiceDate"`

	Contract Contract `gorm:"foreignKey:ContractID" json:"contract"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Product  Product  `gorm:"foreignKey:ProductID" json:"product"`

	IsPass bool   `gorm:"-" json:"isPass"`
	Remark string `gorm:"-" json:"remark"`
}

func InsertPurchasing(purchasing *Purchasing) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := tx.Create(&purchasing).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&ProductCall{}).Where("product_id", purchasing.ProductID).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SubmitPurchasing(purchasing *Purchasing) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = magic.PURCHASING_STATUS_NO_CHECK
	maps["create_date"] = time.Now()
	maps["is_second_final"] = true

	err = db.Transaction(func(tx *gorm.DB) error {
		//扣除一百补贴
		if tErr := tx.Model(&Employee{}).Where("id = ?", purchasing.EmployeeID).Update("role_credit_del", gorm.Expr("role_credit_del + ?", 100)).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&Purchasing{}).
			Where("contract_id = ? AND task_id = ? AND employee_id = ? AND status = ?", purchasing.ContractID, purchasing.TaskID, purchasing.EmployeeID, magic.PURCHASING_STATUS_SAVE).
			Updates(maps).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeletePurchasing(Purchasing *Purchasing) (code int) {
	err = db.Model(&Purchasing).Where("id = ?", Purchasing.ID).Updates(map[string]interface{}{"employee_id": nil, "contract_id": nil, "task_id": nil, "is_delete": true}).Error
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

			var product Product
			if tErr := tx.First(&product, purchasing.ProductID).Error; tErr != nil {
				return tErr
			}
			if product.PurchasePrice < purchasing.Price {
				if tErr := tx.Model(&Product{}).Where("id", purchasing.ProductID).Update("purchase_price", purchasing.Price).Error; tErr != nil {
					return tErr
				}
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

func UpdatePurchasingProductStatus(purchasing *Purchasing, maps map[string]interface{}, employeeID int) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		if tErr := tx.Model(&Purchasing{}).Where("id", purchasing.ID).Updates(maps).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Exec("UPDATE product SET number = number + ? WHERE id = ?", purchasing.RealNumber, purchasing.ProductID).Error; tErr != nil {
			return tErr
		}
		historyProduct := HistoryProduct{
			CreateDate: XDate{Time: time.Now()},
			EmployeeID: employeeID,
			ProductID:  purchasing.ProductID,
			Number:     purchasing.RealNumber,
			Remark:     "采购(ID:" + strconv.Itoa(purchasing.ID) + ")确认收货",
		}
		if tErr := InsertHistoryProduct(&historyProduct, tx); tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectPurchasings(purchasingQuery *Purchasing, xForms *XForms) (purchasings []Purchasing, code int) {

	var maps = make(map[string]interface{})
	maps["purchasing.is_delete"] = false
	if purchasingQuery.ContractID != 0 {
		maps["purchasing.contract_id"] = purchasingQuery.ContractID
	}
	if purchasingQuery.EmployeeID != 0 {
		maps["purchasing.employee_id"] = purchasingQuery.EmployeeID
	}
	if purchasingQuery.Status != 0 {
		maps["purchasing.status"] = purchasingQuery.Status
	}
	if purchasingQuery.Type != 0 {
		maps["purchasing.type"] = purchasingQuery.Type
	}

	tx := db.Where(maps)

	if purchasingQuery.Status == 0 {
		tx = tx.Where("purchasing.status <> ?", purchasingQuery.Status)
	}

	err = tx.Find(&purchasings).Count(&xForms.Total).
		Preload("Product").Preload("Employee").Preload("Contract").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&purchasings).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return purchasings, msg.SUCCESS
}

func SelectAllSavePurchasings(contractID int, taskID int, employeeID int) (purchasings []Purchasing, code int) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	maps["contract_id"] = contractID
	maps["task_id"] = taskID
	maps["employee_id"] = employeeID
	maps["status"] = magic.PURCHASING_STATUS_SAVE

	err = db.Where(maps).
		Preload("Product").
		Find(&purchasings).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return purchasings, msg.SUCCESS
}

func SelectAllPurchasings(maps map[string]interface{}) (purchasings []Purchasing, code int) {
	err = db.Where(maps).
		Preload("Product").Preload("Employee").
		Find(&purchasings).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return purchasings, msg.SUCCESS

}
