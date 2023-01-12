package models

import (
	"errors"
	"fmt"
	"oa-backend/utils/magic"
	"oa-backend/utils/msg"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	// UID                   string  `gorm:"type:varchar(32);comment:唯一标识;unique" json:"UID"`
	// RegionUID             string  `gorm:"type:varchar(32);comment:省份UID;default:(-)" json:"regionUID"`
	// OfficeUID             string  `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	// EmployeeUID           string  `gorm:"type:varchar(32);comment:业务员UID;default:(-)" json:"employeeUID"`
	// CustomerUID           string  `gorm:"type:varchar(32);comment:客户ID;default:(-)" json:"customerUID"`
	// ContractUnitUID string `gorm:"type:varchar(32);comment:签订单位;default:(-)" json:"contractUnitUID"`
	// AuditorUID            string  `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditorUID"`
	// ContractUnitID        int     `gorm:"type:int;comment:签订单位;default:(-)" json:"contractUnitID"`

	ID                    int     `gorm:"primary_key" json:"id"`
	IsDelete              bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	No                    string  `gorm:"type:varchar(100);comment:合同编号" json:"no"`
	RegionID              int     `gorm:"type:int;comment:省份ID;default:(-)" json:"regionID"`
	OfficeID              int     `gorm:"type:int;comment:办事处ID;default:(-)" json:"officeID"`
	EmployeeID            int     `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	CustomerID            int     `gorm:"type:int;comment:客户ID;default:(-)" json:"customerID"`
	ContractDate          XDate   `gorm:"type:date;comment:签订日期" json:"contractDate"`
	VendorID              int     `gorm:"type:int;comment:签订单位;default:(-)" json:"vendorID"`
	EstimatedDeliveryDate XDate   `gorm:"type:date;comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       XDate   `gorm:"type:date;comment:实际交货日期;default:(-)" json:"endDeliveryDate"`
	EndPaymentDate        XDate   `gorm:"type:date;comment:完成回款日期;default:(-)" json:"endPaymentDate"`
	InvoiceType           int     `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string  `gorm:"type:varchar(300);comment:开票内容" json:"invoiceContent"`
	PaymentContent        string  `gorm:"type:varchar(300);comment:付款方式" json:"paymentContent"`
	IsSpecial             bool    `gorm:"type:boolean;comment:是否是特殊合同" json:"isSpecial"`
	IsPreDeposit          bool    `gorm:"type:boolean;comment:是否是预存款合同" json:"isPreDeposit"`
	PreDeposit            float64 `gorm:"type:decimal(20,6);comment:可用预存款金额" json:"preDeposit"`
	PreDepositRecord      float64 `gorm:"type:decimal(20,6);comment:合同预存款金额" json:"preDepositRecord"`
	PayType               int     `gorm:"type:int;comment:付款类型(1:人民币 2:美元)" json:"payType"`
	TotalAmount           float64 `gorm:"type:decimal(20,6);comment:总金额" json:"totalAmount"`
	PaymentTotalAmount    float64 `gorm:"type:decimal(20,6);comment:回款总金额(人民币)" json:"paymentTotalAmount"`
	Remark                string  `gorm:"type:varchar(300);comment:备注" json:"remark"`
	ProductionStatus      int     `gorm:"type:int;comment:生产状态(1:生产中 2:生产完成)" json:"productionStatus"`
	CollectionStatus      int     `gorm:"type:int;comment:回款状态(1:回款中 2:回款完成)" json:"collectionStatus"`
	Status                int     `gorm:"type:int;comment:状态(-1:审批驳回 0:暂存 1:待审批 2:未完成 3:已完成)" json:"status"`
	AuditorID             int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateDate            XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate             XDate   `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`

	Tasks    []Task    `json:"tasks"`
	Invoices []Invoice `json:"invoices"`
	Payments []Payment `json:"payments"`

	Region   Region   `gorm:"foreignKey:RegionID" json:"region"`
	Office   Office   `gorm:"foreignKey:OfficeID" json:"office"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Vendor   Vendor   `gorm:"foreignKey:VendorID" json:"vendor"`
	Auditor  Employee `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass           bool   `gorm:"-" json:"isPass"`
	StartDate        string `gorm:"-" json:"startDate"`
	EndDate          string `gorm:"-" json:"endDate"`
	IsSpecialNum     int    `gorm:"-" json:"isSpecialNum"`
	IsPreDepositNum  int    `gorm:"-" json:"isPreDepositNum"`
	HavingInvoiceNum int    `gorm:"-" json:"havingInvoiceNum"`
	ProductName      string `gorm:"-" json:"productName"`
}

type Task struct {
	// UID                  string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// ContractUID          string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	// ProductUID           string  `gorm:"type:varchar(32);comment:产品ID" json:"productUID"`
	// Unit                 string  `gorm:"type:varchar(50);comment:单位" json:"unit"`
	// StandardPrice        float64 `gorm:"type:decimal(20,6);comment:下单时标准价格" json:"standardPrice"`
	// StandardPriceUSD     float64 `gorm:"type:decimal(20,6);comment:下单时标准价格(美元)" json:"standardPriceUSD"`
	// TechnicianManUID     string  `gorm:"type:varchar(32);comment:技术负责人ID;default:(-)" json:"technicianManUID"`
	// PurchaseManUID       string  `gorm:"type:varchar(32);comment:采购负责人ID;default:(-)" json:"purchaseManUID"`
	// InventoryManUID      string  `gorm:"type:varchar(32);comment:仓库负责人ID;default:(-)" json:"inventoryManUID"`
	// ShipmentManUID       string  `gorm:"type:varchar(32);comment:物流人员ID;default:(-)" json:"shipmentManUID"`

	// TechnicianRealEndDate XDate   `gorm:"type:date;comment:技术实际提交工作日期;default:(-)" json:"technicianRealEndDate"`
	// PurchaseRealEndDate  XDate `gorm:"type:date;comment:采购实际提交工作日期;default:(-)" json:"purchaseRealEndDate"`
	// InventoryRealEndDate XDate `gorm:"type:date;comment:仓库实际提交工作日期;default:(-)" json:"inventoryRealEndDate"`
	// ShipmentRealEndDate  XDate `gorm:"type:date;comment:物流实际提交工作日期;default:(-)" json:"shipmentRealEndDate"`

	ID                   int     `gorm:"primary_key" json:"id"`
	IsDelete             bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ContractID           int     `gorm:"type:int;comment:合同ID" json:"contractID"`
	ProductID            int     `gorm:"type:int;comment:产品ID" json:"productID"`
	ProductAttributeID   int     `gorm:"type:int;comment:产品属性ID" json:"productAttributeID"`
	Number               int     `gorm:"type:int;comment:数量" json:"number"`
	Price                float64 `gorm:"type:decimal(20,6);comment:单价" json:"price"`
	TotalPrice           float64 `gorm:"type:decimal(20,6);comment:总价" json:"totalPrice"`
	PaymentTotalPrice    float64 `gorm:"type:decimal(20,6);comment:回款总金额" json:"paymentTotalPrice"`
	Status               int     `gorm:"type:int;comment:状态(1:待设计 2:待采购 3:待入/出库 4:待装配 5:待发货 6:已发货)" json:"status"`
	Type                 int     `gorm:"type:int;comment:类型(1:标准/第三方有库存 2:标准/第三方无库存 3:非标准定制)" json:"type"`
	TechnicianManID      int     `gorm:"type:int;comment:技术负责人ID;default:(-)" json:"technicianManID"`
	PurchaseManID        int     `gorm:"type:int;comment:采购负责人ID;default:(-)" json:"purchaseManID"`
	InventoryManID       int     `gorm:"type:int;comment:仓库负责人ID;default:(-)" json:"inventoryManID"`
	ShipmentManID        int     `gorm:"type:int;comment:物流人员ID;default:(-)" json:"shipmentManID"`
	Remark               string  `gorm:"type:varchar(300);comment:业务备注" json:"remark"`
	PushMoney            float64 `gorm:"type:decimal(20,6);comment:提成(元)" json:"pushMoney"`
	PushMoneyPercentages float64 `gorm:"type:decimal(20,6);comment:特殊合同提成百分比" json:"pushMoneyPercentages"`
	TechnicianDays       int     `gorm:"type:int;comment:技术预计工作天数;default:(-)" json:"technicianDays"`
	PurchaseDays         int     `gorm:"type:int;comment:采购预计工作天数;default:(-)" json:"purchaseDays"`
	TechnicianStartDate  XDate   `gorm:"type:date;comment:技术接到工作日期;default:(-)" json:"technicianStartDate"`
	TechnicianEndDate    XDate   `gorm:"type:date;comment:技术预期完成日期;default:(-)" json:"technicianEndDate"`
	TechnicianFinalDate  XDate   `gorm:"type:date;comment:技术实际完成日期;default:(-)" json:"technicianFinalDate"`
	PurchaseStartDate    XDate   `gorm:"type:date;comment:采购接到工作日期;default:(-)" json:"purchaseStartDate"`
	PurchaseEndDate      XDate   `gorm:"type:date;comment:采购预期完成日期;default:(-)" json:"purchaseEndDate"`
	PurchaseFinalDate    XDate   `gorm:"type:date;comment:采购实际完成日期;default:(-)" json:"purchaseFinalDate"`
	InventoryStartDate   XDate   `gorm:"type:date;comment:仓库接到工作日期;default:(-)" json:"inventoryStartDate"`
	InventoryFinalDate   XDate   `gorm:"type:date;comment:仓库实际完成日期;default:(-)" json:"inventoryFinalDate"`
	AssemblyStartDate    XDate   `gorm:"type:date;comment:装配接到工作日期;default:(-)" json:"assemblyStartDate"`
	AssemblyFinalDate    XDate   `gorm:"type:date;comment:装配实际完成日期;default:(-)" json:"assemblyFinalDate"`
	ShipmentStartDate    XDate   `gorm:"type:date;comment:物流接到工作日期;default:(-)" json:"shipmentStartDate"`
	ShipmentFinalDate    XDate   `gorm:"type:date;comment:物流实际完成日期;default:(-)" json:"shipmentFinalDate"`
	TechnicianRemark     string  `gorm:"type:varchar(300);comment:技术备注" json:"technicianRemark"`
	PurchaseRemark       string  `gorm:"type:varchar(300);comment:采购备注" json:"purchaseRemark"`
	InventoryRemark      string  `gorm:"type:varchar(300);comment:仓库备注" json:"inventoryRemark"`
	AssemblyRemark       string  `gorm:"type:varchar(300);comment:装配备注" json:"assemblyRemark"`
	ShipmentRemark       string  `gorm:"type:varchar(300);comment:物流备注" json:"shipmentRemark"`

	AuditorID  int   `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateDate XDate `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate  XDate `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`

	Contract         Contract         `gorm:"foreignKey:ContractID" json:"contract"`
	Product          Product          `gorm:"foreignKey:ProductID" json:"product"`
	ProductAttribute ProductAttribute `gorm:"foreignKey:ProductAttributeID" json:"productAttribute"`
	TechnicianMan    Employee         `gorm:"foreignKey:TechnicianManID" json:"technicianMan"`
	PurchaseMan      Employee         `gorm:"foreignKey:PurchaseManID" json:"purchaseMan"`
	InventoryMan     Employee         `gorm:"foreignKey:InventoryManID" json:"inventoryMan"`
	ShipmentMan      Employee         `gorm:"foreignKey:ShipmentManID" json:"shipmentMan"`
	Auditor          Employee         `gorm:"foreignKey:AuditorID" json:"auditor"`
	Payments         []Payment        `json:"payments"`

	CustomerName string `gorm:"-" json:"customerName"`
	EndDate      string `gorm:"-" json:"endDate"`
	ProductName  string `gorm:"-" json:"productName"`
}

type Invoice struct {
	// UID         string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// ContractUID string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	// EmployeeUID string  `gorm:"type:varchar(32);comment:添加员工UID;default:(-)" json:"employeeUID"`
	// Code        string  `gorm:"type:varchar(100);comment:发票号" json:"code"`

	ID         int     `gorm:"primary_key" json:"id"`
	IsDelete   bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ContractID int     `gorm:"type:int;comment:合同ID" json:"contractID"`
	EmployeeID int     `gorm:"type:int;comment:财务ID;default:(-)" json:"employeeID"`
	No         string  `gorm:"type:varchar(100);comment:发票号" json:"no"`
	Money      float64 `gorm:"type:decimal(20,6);comment:总金额" json:"money"`
	CreateDate XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`

	Contract Contract `gorm:"foreignKey:ContractID" json:"contract"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
}

type Payment struct {
	// UID         string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	// ContractUID          string  `gorm:"type:varchar(32);comment:合同UID;default:(-)" json:"contractUID"`
	// TaskUID              string  `gorm:"type:varchar(32);comment:任务UID;default:(-)" json:"taskUID"`
	// EmployeeUID          string  `gorm:"type:varchar(32);comment:录入人员ID;default:(-)" json:"employeeUID"`

	ID                   int     `gorm:"primary_key" json:"id"`
	IsDelete             bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ContractID           int     `gorm:"type:int;comment:合同ID;default:(-)" json:"contractID"`
	TaskID               int     `gorm:"type:int;comment:任务ID;default:(-)" json:"taskID"`
	PaymentDate          XDate   `gorm:"type:date;comment:回款日期" json:"paymentDate"`
	EmployeeID           int     `gorm:"type:int;comment:财务ID;default:(-)" json:"employeeID"`
	Money                float64 `gorm:"type:decimal(20,6);comment:回款金额(人民币)" json:"money"`
	TheoreticalPushMoney float64 `gorm:"type:decimal(20,6);comment:理论提成(元)" json:"theoreticalPushMoney"`
	Fine                 float64 `gorm:"type:decimal(20,6);comment:回款延迟扣除(元)" json:"fine"`
	PushMoney            float64 `gorm:"type:decimal(20,6);comment:实际提成(元)" json:"pushMoney"`
	BusinessMoney        float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`
	CreateDate           XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`

	Task     Task     `gorm:"foreignKey:TaskID" json:"task"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
}

func createNo(contract *Contract) (no string, code int) {
	var office Office
	var employee Employee

	_ = GeneralSelect(&office, contract.OfficeID, nil)
	_ = GeneralSelect(&employee, contract.EmployeeID, nil)

	if office.ID != 0 && employee.ID != 0 {
		employeeCount := strconv.Itoa(employee.ContractCount + 1)
		if len(employeeCount) == 1 {
			employeeCount = "00" + employeeCount
		} else if len(employeeCount) == 2 {
			employeeCount = "0" + employeeCount
		}
		no = "Bjscistar" + strings.ReplaceAll(contract.CreateDate.Format("2006-01-02"), "-", "") + "-" + office.Number + employee.Number + employeeCount
		code = msg.SUCCESS
	} else {
		code = msg.ERROR
	}
	return
}

func InsertContract(contract *Contract) (code int) {

	if contract.ID == 0 {
		err = db.Create(&contract).Error
	} else {
		err = db.Transaction(func(tx *gorm.DB) error {

			var contractBak Contract

			GeneralSelect(&contractBak, contract.ID, nil)

			if tErr := tx.Model(&Contract{}).Where("id = ?", contract.ID).Updates(map[string]interface{}{"is_delete": true, "employee_id": nil}).Error; tErr != nil {
				return tErr
			}

			contract.CreateDate = contractBak.ContractDate
			contract.ID = 0

			if tErr := tx.Create(&contract).Error; tErr != nil {
				return tErr
			}
			return nil
		})

		err = db.Updates(&contract).Error
	}

	if err != nil {
		return msg.ERROR
	} else {
		return msg.SUCCESS
	}
}

func ApproveContract(contractBak *Contract, maps map[string]interface{}) (code int) {
	if contractBak.IsPass {
		var no string
		no, code = createNo(contractBak)
		if code == msg.SUCCESS {
			maps["no"] = no
			maps["status"] = magic.CONTRACT_STATUS_NOT_FINISH
			maps["production_status"] = magic.CONTRATCT_PRODUCTION_STATUS_ING
			maps["collection_status"] = magic.CONTRATCT_COLLECTION_STATUS_ING

			err = db.Transaction(func(tx *gorm.DB) error {

				if contractBak.IsPreDeposit {
					var contract Contract
					if tErr := tx.Preload("Tasks").Where("is_delete = ?", false).First(&contract, contractBak.ID).Error; tErr != nil {
						return tErr
					}
					//产品可售库存减少
					for i := range contract.Tasks {
						if tErr := tx.Exec("UPDATE product SET number = number - ? WHERE id = ?", contract.Tasks[i].Number, contract.Tasks[i].ProductID).Error; tErr != nil {
							return tErr
						}
					}
				}

				//业务员累计合同数目+1
				if tErr := tx.Exec("UPDATE employee SET contract_count = contract_count + 1 WHERE id = ?", contractBak.EmployeeID).Error; tErr != nil {
					return tErr
				}
				//更新合同信息
				if tErr := tx.Model(&Contract{ID: contractBak.ID}).Updates(maps).Error; tErr != nil {
					return tErr
				}

				return nil
			})
		} else {
			return
		}
	} else {
		maps["status"] = magic.CONTRACT_STATUS_REJECT
		err = db.Transaction(func(tx *gorm.DB) error {
			//更新合同信息
			if tErr := tx.Model(&Contract{}).Where("id", contractBak.ID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	}
	return
}

// 合同完成
func FinalContract(contract *Contract) (code int) {
	var contractBak Contract
	_ = db.Preload("Tasks").First(&contractBak, contract.ID).Error

	if contractBak.ID != 0 && contractBak.CollectionStatus == magic.CONTRATCT_COLLECTION_STATUS_FINISH {
		isFinish := true
		if contractBak.IsPreDeposit {
			for k := range contractBak.Tasks {
				if contractBak.Tasks[k].Status != magic.TASK_STATUS_SHIPMENT &&
					contractBak.Tasks[k].Status != magic.TASK_STATUS_REJECT {
					isFinish = false
					break
				}
			}
		} else {
			if contractBak.ProductionStatus != magic.CONTRATCT_PRODUCTION_STATUS_FINISH {
				isFinish = false
			}
		}

		if isFinish {
			err = db.Model(&Contract{}).Where("id = ?", contract.ID).Update("status", magic.CONTRACT_STATUS_FINISH).Error
			if err != nil {
				return msg.FAIL
			}
			return msg.SUCCESS
		}
	}
	return msg.FAIL

	// if contractBak.ID != 0 && contractBak.IsPreDeposit &&
	// 	contractBak.CollectionStatus == magic.CONTRATCT_COLLECTION_STATUS_FINISH {
	// 	isFinish := true
	// 	for k := range contractBak.Tasks {
	// 		if contractBak.Tasks[k].Status != magic.TASK_STATUS_SHIPMENT &&
	// 			contractBak.Tasks[k].Status != magic.TASK_STATUS_REJECT {
	// 			isFinish = false
	// 			break
	// 		}
	// 	}

	// 	if isFinish {
	// 		err = db.Model(&Contract{}).Where("id = ?", contract.ID).Update("status", magic.CONTRACT_STATUS_FINISH).Error
	// 		if err != nil {
	// 			return msg.FAIL
	// 		}
	// 		return msg.SUCCESS
	// 	}
	// } else {
	// 	code = msg.FAIL
	// }

}

//回款状态回退

// employeeID 记录用
func RejectContract(contract *Contract, employeeID int) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {
		var payments []Payment
		if tErr := tx.Preload("Task.Product.Type").Where("contract_id = ?", contract.ID).Find(&payments).Error; tErr != nil {
			return tErr
		}
		//回款记录统计
		var tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney float64
		if contract.IsPreDeposit {
			for k := range payments {
				if payments[k].TaskID != 0 {
					//预付款自动生成的分成利润
					tempMoney += payments[k].PushMoney * 0.5
					tempMoneyCold += payments[k].PushMoney - payments[k].PushMoney*0.5
					tempBusinessMoney += payments[k].BusinessMoney
					//产品类型是否计算任务量
					if !payments[k].Task.Product.Type.IsTaskLoad {
						tempTargetLoad -= payments[k].Money
					}
				} else {
					//预付款任务量
					tempTargetLoad += payments[k].Money
				}
			}
		} else {
			for k := range payments {
				tempMoney += payments[k].PushMoney * 0.5
				tempMoneyCold += payments[k].PushMoney - payments[k].PushMoney*0.5
				tempBusinessMoney += payments[k].BusinessMoney
				//产品类型是否计算任务量
				if payments[k].Task.Product.Type.IsTaskLoad {
					tempTargetLoad += payments[k].Money
				}
			}
		}

		//历史记录
		var officeBak Office
		if tErr := tx.First(&officeBak, contract.OfficeID).Error; tErr != nil {
			return tErr
		}
		historyOffice := HistoryOffice{
			OfficeID:            officeBak.ID,
			EmployeeID:          employeeID,
			OldBusinessMoney:    officeBak.BusinessMoney,
			OldMoney:            officeBak.Money,
			OldMoneyCold:        officeBak.MoneyCold,
			OldTargetLoad:       officeBak.TargetLoad,
			ChangeBusinessMoney: -tempBusinessMoney,
			ChangeMoney:         -tempMoney,
			ChangeMoneyCold:     -tempMoneyCold,
			ChangeTargetLoad:    -tempTargetLoad,
			NewBusinessMoney:    officeBak.BusinessMoney - tempBusinessMoney,
			NewMoney:            officeBak.Money - tempMoney,
			NewMoneyCold:        officeBak.MoneyCold - tempMoneyCold,
			NewTargetLoad:       officeBak.TargetLoad - tempTargetLoad,
			CreateDate:          XDate{Time: time.Now()},
			Remark:              "合同(" + contract.No + ")驳回",
		}
		if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
			return tErr
		}

		//办事处记录变更
		if tErr := tx.Exec("UPDATE office SET target_load = target_load - ?,money = money - ?, money_cold = money_cold - ?, business_money = business_money - ? WHERE id = ?", tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney, contract.OfficeID).Error; tErr != nil {
			return tErr
		}

		//合同状态变更为驳回
		if tErr := tx.Model(&Contract{}).Where("id", contract.ID).Update("status", magic.CONTRACT_STATUS_REJECT).Error; tErr != nil {
			return tErr
		}

		//删除任务，回款记录，支票
		if tErr := tx.Model(&Task{}).Where("contract_id", contract.ID).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&Payment{}).Where("contract_id", contract.ID).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&Invoice{}).Where("contract_id", contract.ID).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}

		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectContract(id int) (contract Contract, code int) {
	db.Preload("Region").Preload("Office").Preload("Employee").
		Preload("Customer.CustomerCompany").Preload("Vendor").
		Preload("Tasks.TechnicianMan").Preload("Tasks.PurchaseMan").
		Preload("Tasks.InventoryMan").Preload("Tasks.ShipmentMan").
		Preload("Tasks.Product.Type").
		Preload("Tasks.ProductAttribute").
		Preload("Invoices.Employee").
		Preload("Payments.Task.Product.Type").
		Preload("Payments.Employee").
		Where("contract.is_delete = ?", false).
		First(&contract, id)
	if contract.ID == 0 {
		return Contract{}, msg.FAIL
	}
	return contract, msg.SUCCESS
}

func SelectMySaveContracts(employeeID int, xForms *XForms) (contracts []Contract, code int) {
	var maps = make(map[string]interface{})
	maps["contract.is_delete"] = false
	maps["contract.employee_id"] = employeeID
	maps["contract.status"] = 0

	err = db.Where(maps).Find(&contracts).Count(&xForms.Total).
		Preload("Customer.CustomerCompany.Region").Preload("Employee").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return contracts, msg.SUCCESS
}

func SelectContracts(contractQuery *Contract, xForms *XForms) (contracts []Contract, code int) {
	var maps = make(map[string]interface{})
	maps["contract.is_delete"] = false

	if contractQuery.EmployeeID != 0 {
		maps["contract.employee_id"] = contractQuery.EmployeeID
	}
	if contractQuery.RegionID != 0 {
		maps["contract.region_id"] = contractQuery.RegionID
	}
	if contractQuery.Status != 0 {
		maps["contract.status"] = contractQuery.Status
	}
	if contractQuery.ProductionStatus != 0 {
		maps["contract.production_status"] = contractQuery.ProductionStatus
	}
	if contractQuery.CollectionStatus != 0 {
		maps["contract.collection_status"] = contractQuery.CollectionStatus
	}
	if contractQuery.PayType != 0 {
		maps["contract.pay_type"] = contractQuery.PayType
	}
	if contractQuery.IsSpecialNum == 1 {
		maps["contract.is_special"] = true
	} else if contractQuery.IsSpecialNum == 2 {
		maps["contract.is_special"] = false
	}
	if contractQuery.IsPreDepositNum == 1 {
		maps["contract.is_pre_deposit"] = true
	} else if contractQuery.IsPreDepositNum == 2 {
		maps["contract.is_pre_deposit"] = false
	}

	tx := db.Where(maps)

	if contractQuery.Status == 0 {
		tx = tx.Where("contract.status <> ?", contractQuery.Status)
	}

	if contractQuery.No != "" {
		tx = tx.Where("contract.no LIKE ?", "%"+contractQuery.No+"%")
	}

	if contractQuery.Employee.Name != "" {
		tx = tx.Joins("Employee").Where("Employee.name LIKE ?", "%"+contractQuery.Employee.Name+"%")
	}

	if contractQuery.ProductName != "" {
		tx = tx.Where("contract.id IN (?)", db.Table("task").Select("task.contract_id").
			Joins("left join product on task.product_id = product.id").
			Where("task.contract_id is not null").
			Where("product.name like ?", "%"+contractQuery.ProductName+"%").
			Group("task.contract_id"))
	}

	if contractQuery.Customer.CustomerCompany.Name != "" && contractQuery.Customer.Name != "" {
		tx = tx.Joins("customer").
			Joins("left join customer_company on customer.customer_company_id = customer_company.id").
			Where("customer.name LIKE ? AND customer_company.name LIKE ?", "%"+contractQuery.Customer.Name+"%", "%"+contractQuery.Customer.CustomerCompany.Name+"%")
	} else {
		if contractQuery.Customer.CustomerCompany.Name != "" {
			tx = tx.Joins("customer").
				Joins("left join customer_company on Customer.customer_company_id = customer_company.id").
				Where("customer_company.name LIKE ?", "%"+contractQuery.Customer.CustomerCompany.Name+"%")
		}
		if contractQuery.Customer.Name != "" {
			tx = tx.Joins("customer").
				Where("customer.name LIKE ?", "%"+contractQuery.Customer.Name+"%")
		}
	}

	if contractQuery.StartDate != "" && contractQuery.EndDate != "" {
		tx = tx.Where("contract.contract_date BETWEEN ? AND ?", contractQuery.StartDate, contractQuery.EndDate)
	} else {
		if contractQuery.StartDate != "" {
			tx = tx.Where("contract.contract_date >= ?", contractQuery.StartDate)
		}
		if contractQuery.EndDate != "" {
			tx = tx.Where("contract.contract_date <= ?", contractQuery.EndDate)
		}
	}

	if contractQuery.HavingInvoiceNum == 1 {
		tx = tx.Where("contract.id NOT IN (?)", db.Table("invoice").Select("contract_id").Where("contract_id is not null").Group("contract_id"))
	} else if contractQuery.HavingInvoiceNum == 2 {
		tx = tx.Where("contract.id IN (?)", db.Table("invoice").Select("contract_id").Where("contract_id is not null").Group("contract_id"))
	}

	err = tx.Find(&contracts).Count(&xForms.Total).
		Preload("Office").Preload("Employee").
		Preload("Customer.CustomerCompany").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return contracts, msg.SUCCESS
}

func SelectTask(id int) (task Task, code int) {
	db.Preload("Contract").
		Where("task.is_delete = ?", false).
		First(&task, id)
	if task.ID == 0 {
		return Task{}, msg.FAIL
	}
	return task, msg.SUCCESS
}

// 预存款合同添加任务
func InsertTask(contract *Contract, task *Task) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {
		//创建任务
		if tErr := tx.Create(&task).Error; tErr != nil {
			return tErr
		}
		//减去合同预存款并加上合同总金额
		if tErr := tx.Exec("UPDATE contract SET pre_deposit = pre_deposit - ? WHERE id = ?", task.TotalPrice, contract.ID).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 分发任务
func DistributeTask(task *Task, maps map[string]interface{}, employeeID int) (code int) {

	if task.Contract.IsPreDeposit {
		var taskBak Task
		//预存款合同 需要 同时生成一条回款记录
		err = db.Transaction(func(tx *gorm.DB) error {
			if tErr := tx.Preload("Contract").
				Preload("Product.Type").
				Preload("ProductAttribute").
				First(&taskBak, task.ID).Error; tErr != nil {
				return tErr
			}
			var payment Payment
			payment.ContractID = taskBak.ContractID
			payment.TaskID = taskBak.ID
			payment.CreateDate = task.AuditDate
			payment.PaymentDate = task.AuditDate
			payment.Money = taskBak.TotalPrice

			//计算提成
			payment.TheoreticalPushMoney, payment.Fine, payment.PushMoney, payment.BusinessMoney = calculatePercentage(&payment, &taskBak)
			//创建记录
			if tErr := tx.Create(&payment).Error; tErr != nil {
				return tErr
			}
			//更新合同数据
			if tErr := tx.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE id = ?", payment.Money, payment.ContractID).Error; tErr != nil {
				return tErr
			}
			//更新办事处数据&生成历史记录
			tempPushMoney1 := payment.PushMoney * 0.5
			tempPushMoney2 := payment.PushMoney - tempPushMoney1
			var officeBak Office
			if tErr := tx.First(&officeBak, taskBak.Contract.OfficeID).Error; tErr != nil {
				return tErr
			}
			historyOffice := HistoryOffice{
				OfficeID:            taskBak.Contract.OfficeID,
				EmployeeID:          employeeID,
				OldBusinessMoney:    officeBak.BusinessMoney,
				OldMoney:            officeBak.Money,
				OldMoneyCold:        officeBak.MoneyCold,
				OldTargetLoad:       officeBak.TargetLoad,
				ChangeBusinessMoney: payment.BusinessMoney,
				ChangeMoney:         tempPushMoney1,
				ChangeMoneyCold:     tempPushMoney2,
				ChangeTargetLoad:    0,
				NewBusinessMoney:    officeBak.BusinessMoney + payment.BusinessMoney,
				NewMoney:            officeBak.Money + tempPushMoney1,
				NewMoneyCold:        officeBak.MoneyCold + tempPushMoney2,
				NewTargetLoad:       0,
				CreateDate:          XDate{Time: time.Now()},
				Remark:              "预存款合同(" + taskBak.Contract.No + ")采购商品",
			}
			if task.Product.Type.IsTaskLoad {
				if tErr := tx.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", payment.Money, tempPushMoney1, tempPushMoney2, payment.BusinessMoney, taskBak.Contract.OfficeID).Error; tErr != nil {
					return tErr
				}
				historyOffice.ChangeTargetLoad = payment.Money
				historyOffice.NewTargetLoad = officeBak.TargetLoad + payment.Money
			} else {
				if tErr := tx.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempPushMoney1, tempPushMoney2, payment.BusinessMoney, taskBak.Contract.OfficeID).Error; tErr != nil {
					return tErr
				}
				historyOffice.NewTargetLoad = officeBak.TargetLoad
			}
			if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
				return tErr
			}
			//跟新task记录
			if tErr := tx.Model(&Task{}).Where("id = ?", taskBak.ID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			return nil
		})

		if err != nil {
			code = msg.ERROR
		} else {
			code = msg.SUCCESS
		}

	} else {
		//普通合同直接修改
		code = GeneralUpdate(&Task{}, task.ID, maps)
	}

	return
}

func NextTask(task *Task, maps map[string]interface{}) (code int) {

	err = db.Transaction(func(tx *gorm.DB) error {

		if maps["status"] == magic.TASK_STATUS_SHIPMENT {
			var contract Contract

			if tErr := tx.Preload("Tasks").
				First(&contract, task.ContractID).Error; tErr != nil {
				return tErr
			}

			if contract.ID != 0 && !contract.IsPreDeposit {
				isFinish := true
				for _, taskBak := range contract.Tasks {
					if taskBak.ID != task.ID && taskBak.Status != magic.TASK_STATUS_SHIPMENT {
						isFinish = false
						break
					}
				}

				if isFinish {
					if tErr := tx.Model(&Contract{}).
						Where("id = ?", task.ContractID).
						Update("production_status", magic.CONTRATCT_PRODUCTION_STATUS_FINISH).Error; tErr != nil {
						return tErr
					}

					// if contract.CollectionStatus == magic.CONTRATCT_COLLECTION_STATUS_FINISH {
					// 	if tErr := tx.Model(&Contract{}).
					// 		Where("id = ?", task.ContractID).
					// 		Updates(map[string]interface{}{"production_status": magic.CONTRATCT_PRODUCTION_STATUS_FINISH, "status": magic.CONTRACT_STATUS_FINISH}).Error; tErr != nil {
					// 		return tErr
					// 	}
					// } else {
					// 	if tErr := tx.Model(&Contract{}).
					// 		Where("id = ?", task.ContractID).
					// 		Update("production_status", magic.CONTRATCT_PRODUCTION_STATUS_FINISH).Error; tErr != nil {
					// 		return tErr
					// 	}
					// }
				}

			}
		}

		if tErr := tx.Model(&Task{}).Where("id", task.ID).Updates(maps).Error; tErr != nil {
			return tErr
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	}

	return msg.SUCCESS
}

func ResetTask(task *Task, maps map[string]interface{}) (code int) {

	if task.Contract.IsSpecial && maps["push_money_percentages"] != task.PushMoneyPercentages {
		//特殊合同重置要修改利润提成payment
		//TODO
		code = GeneralUpdate(&Task{}, task.ID, maps)
	} else {
		//非特殊合同直接修改
		code = GeneralUpdate(&Task{}, task.ID, maps)
	}

	return
}

// 预存款合同驳回请求
func RejectTask(id int) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {

		var task Task

		if tErr := tx.Preload("Product.Type").First(&task, id).Error; tErr != nil {
			return tErr
		}

		//驳回该任务
		if tErr := tx.Model(&Task{}).Where("id = ?", task.ID).Update("status", magic.TASK_STATUS_REJECT).Error; tErr != nil {
			return tErr
		}
		//退回该任务的预存款金额
		if tErr := tx.Exec("UPDATE contract SET pre_deposit = pre_deposit + ? WHERE id = ?", task.TotalPrice, task.ContractID).Error; tErr != nil {
			return tErr
		}
		//任务若已经分配了退报销
		if task.Status != magic.TASK_STATUS_NOT_DISTRIBUTE {
			var payment Payment

			if tErr := tx.Where("task_id = ? AND is_delete = ?", task.ID, false).First(&payment).Error; tErr != nil {
				return tErr
			}

			if payment.ID != 0 {
				var tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney float64
				tempTargetLoad = payment.Money
				tempMoney = payment.PushMoney
				tempMoneyCold = payment.PushMoney - payment.PushMoney*0.5
				tempBusinessMoney = payment.BusinessMoney
				if task.Product.Type.IsTaskLoad {
					if tErr := tx.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempMoneyCold, tempMoneyCold, tempBusinessMoney, task.Contract.OfficeID).Error; tErr != nil {
						return tErr
					}
				} else {
					if tErr := tx.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney, task.Contract.OfficeID).Error; tErr != nil {
						return tErr
					}
				}

				if tErr := tx.Model(&Payment{}).Where("id", payment.ID).
					Updates(map[string]interface{}{"contract_id": nil, "is_delete": true}).Error; tErr != nil {
					return tErr
				}
				//删除支票
			} else {
				return errors.New("not find payment")
			}
		}
		return nil
	})

	if err != nil {
		return msg.ERROR
	} else {
		return msg.SUCCESS
	}
}

func DeleteTask(id int) (code int) {
	err = db.Model(&Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{"contract_id": nil, "is_delete": true}).Error

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectMyTasks(taskQuery *Task, employeeID int, xForms *XForms) (tasks []Task, code int) {
	var maps = make(map[string]interface{})
	maps["task.is_delete"] = false

	if taskQuery.Status != 0 {
		maps["task.status"] = taskQuery.Status
	}

	tx := db.Where(maps).Where(db.Where("technician_man_id = ?", employeeID).
		Or("purchase_man_id = ?", employeeID).
		Or("inventory_man_id = ?", employeeID).
		Or("shipment_man_id = ?", employeeID))

	tx = tx.Joins("Contract")

	if taskQuery.Contract.No != "" {
		tx = tx.Where("Contract.no LIKE ?", "%"+taskQuery.Contract.No+"%")
	}

	if taskQuery.Status == 0 {
		tx = tx.Where("task.status > 0")
	}

	if taskQuery.ProductName != "" {
		tx = tx.Joins("Product").Where("Product.name LIKE ?", "%"+taskQuery.ProductName+"%")
	}

	if taskQuery.CustomerName != "" {
		tx = tx.Joins("left join customer on Contract.customer_id = customer.id").Where("customer.name LIKE ?", "%"+taskQuery.CustomerName+"%")
	}

	if taskQuery.EndDate != "" {
		tempTime, _ := time.Parse("2006-01-02T15:04:05Z", taskQuery.EndDate)
		fmt.Println(tempTime)
		tx = tx.Where("Contract.estimated_delivery_date <= ?", tempTime)
	}

	err = tx.Find(&tasks).Count(&xForms.Total).
		Preload("Contract").
		Preload("Product.Type").Preload("TechnicianMan").
		Preload("PurchaseMan").Preload("InventoryMan").Preload("ShipmentMan").
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&tasks).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return tasks, msg.SUCCESS
}

func DeleteInvoice(id int) (code int) {
	err = db.Model(&Invoice{}).Where("id", id).
		Updates(map[string]interface{}{"contract_id": nil, "is_delete": true}).Error

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectInvoices(invoiceQuery *Invoice) (invoices []Invoice, code int) {
	err = db.Where("is_delete = ?", false).Where("contract_id = ?", invoiceQuery.ContractID).Find(&invoices).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return invoices, msg.SUCCESS
}

func InsertPayment(payment *Payment) (code int) {
	var contract Contract
	var officeBak Office
	db.First(&contract, payment.ContractID)
	db.First(&officeBak, contract.OfficeID)
	if contract.ID != 0 {
		if contract.IsPreDeposit {
			//预存款合同回款不产生利润
			err = db.Transaction(func(tx *gorm.DB) error {
				//生成记录
				if tErr := tx.Create(&payment).Error; tErr != nil {
					return tErr
				}
				//修改合同可用预存款和回款总额
				if tErr := tx.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ?,pre_deposit = pre_deposit + ? WHERE id = ?", payment.Money, payment.Money, payment.ContractID).Error; tErr != nil {
					return tErr
				}
				//办事处任务量增加
				if tErr := tx.Exec("UPDATE office SET target_load = target_load + ? WHERE id = ?", payment.Money, contract.OfficeID).Error; tErr != nil {
					return tErr
				}
				//生成记录
				historyOffice := HistoryOffice{
					OfficeID:            contract.OfficeID,
					EmployeeID:          payment.EmployeeID,
					OldBusinessMoney:    officeBak.BusinessMoney,
					OldMoney:            officeBak.Money,
					OldMoneyCold:        officeBak.MoneyCold,
					OldTargetLoad:       officeBak.TargetLoad,
					ChangeBusinessMoney: 0,
					ChangeMoney:         0,
					ChangeMoneyCold:     0,
					ChangeTargetLoad:    payment.Money,
					NewBusinessMoney:    officeBak.BusinessMoney,
					NewMoney:            officeBak.Money,
					NewMoneyCold:        officeBak.MoneyCold,
					NewTargetLoad:       officeBak.TargetLoad + payment.Money,
					CreateDate:          XDate{Time: time.Now()},
					Remark:              "预存款合同(" + contract.No + ")添加回款",
				}
				if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
					return tErr
				}
				return nil
			})
		} else {
			err = db.Transaction(func(tx *gorm.DB) error {
				var task Task

				if tErr := tx.Preload("Contract").
					Preload("Product.Type").
					Preload("ProductAttribute").
					First(&task, payment.TaskID).Error; tErr != nil {
					return tErr
				}
				//计算提成
				payment.TheoreticalPushMoney, payment.Fine, payment.PushMoney, payment.BusinessMoney = calculatePercentage(payment, &task)
				//创建记录
				if tErr := tx.Create(&payment).Error; tErr != nil {
					return tErr
				}
				//更新合同数据
				if tErr := tx.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE id = ?", payment.Money, payment.ContractID).Error; tErr != nil {
					return tErr
				}
				//更新任务数据
				if tErr := tx.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE id = ?", payment.Money, payment.TaskID).Error; tErr != nil {
					return tErr
				}
				//更新办事处数据&生成历史记录
				tempPushMoney1 := payment.PushMoney * 0.5
				tempPushMoney2 := payment.PushMoney - tempPushMoney1
				var officeBak Office
				if tErr := tx.First(&officeBak, task.Contract.OfficeID).Error; tErr != nil {
					return tErr
				}
				historyOffice := HistoryOffice{
					OfficeID:            task.Contract.OfficeID,
					EmployeeID:          payment.EmployeeID,
					OldBusinessMoney:    officeBak.BusinessMoney,
					OldMoney:            officeBak.Money,
					OldMoneyCold:        officeBak.MoneyCold,
					OldTargetLoad:       officeBak.TargetLoad,
					ChangeBusinessMoney: payment.BusinessMoney,
					ChangeMoney:         tempPushMoney1,
					ChangeMoneyCold:     tempPushMoney2,
					ChangeTargetLoad:    0,
					NewBusinessMoney:    officeBak.BusinessMoney + payment.BusinessMoney,
					NewMoney:            officeBak.Money + tempPushMoney1,
					NewMoneyCold:        officeBak.MoneyCold + tempPushMoney2,
					NewTargetLoad:       0,
					CreateDate:          XDate{Time: time.Now()},
					Remark:              "普通合同(" + task.Contract.No + ")添加回款",
				}
				if task.Product.Type.IsTaskLoad {
					if tErr := tx.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", payment.Money, tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeID).Error; tErr != nil {
						return tErr
					}
					historyOffice.ChangeTargetLoad = payment.Money
					historyOffice.NewTargetLoad = officeBak.TargetLoad + payment.Money
				} else {
					if tErr := tx.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempPushMoney1, tempPushMoney2, payment.BusinessMoney, contract.OfficeID).Error; tErr != nil {
						return tErr
					}
					historyOffice.NewTargetLoad = officeBak.TargetLoad
				}
				if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
					return tErr
				}
				return nil
			})
		}
		if err != nil {
			code = msg.ERROR
		}
		code = msg.SUCCESS
	} else {
		code = msg.FAIL
	}
	return
}

func UpdatePayment(payment *Payment) (code int) {
	var paymentBak Payment
	var contractBak Contract
	_ = db.First(&paymentBak, payment.ID)
	_ = db.Preload("Office").First(&contractBak, paymentBak.ContractID)
	if contractBak.ID != 0 {
		if contractBak.IsPreDeposit {
			err = db.Transaction(func(tx *gorm.DB) error {
				//合同回款金额更新，总付款金额更新
				tempPaymentTotalAmount := payment.Money - paymentBak.Money
				if tErr := tx.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ?,pre_deposit = pre_deposit + ? WHERE id = ?", tempPaymentTotalAmount, tempPaymentTotalAmount, paymentBak.ContractID).Error; tErr != nil {
					return tErr
				}
				//办事处任务量更新
				if tErr := tx.Exec("UPDATE office SET target_load = target_load + ? WHERE id = ?", tempPaymentTotalAmount, contractBak.OfficeID).Error; tErr != nil {
					return tErr
				}
				//更新回款记录
				var maps = make(map[string]interface{})
				maps["employee_id"] = payment.EmployeeID
				maps["payment_date"] = payment.PaymentDate
				maps["money"] = payment.Money
				if tErr := tx.Model(&Payment{}).Where("id = ?", paymentBak.ID).Updates(maps).Error; tErr != nil {
					return tErr
				}

				historyOffice := HistoryOffice{
					OfficeID:            contractBak.OfficeID,
					EmployeeID:          payment.EmployeeID,
					OldBusinessMoney:    contractBak.Office.BusinessMoney,
					OldMoney:            contractBak.Office.Money,
					OldMoneyCold:        contractBak.Office.MoneyCold,
					OldTargetLoad:       contractBak.Office.TargetLoad,
					ChangeBusinessMoney: 0,
					ChangeMoney:         0,
					ChangeMoneyCold:     0,
					ChangeTargetLoad:    tempPaymentTotalAmount,
					NewBusinessMoney:    contractBak.Office.BusinessMoney,
					NewMoney:            contractBak.Office.Money,
					NewMoneyCold:        contractBak.Office.MoneyCold,
					NewTargetLoad:       contractBak.Office.TargetLoad + tempPaymentTotalAmount,
					CreateDate:          XDate{Time: time.Now()},
					Remark:              "预存款合同(" + contractBak.No + ")回款编辑",
				}
				if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
					return tErr
				}
				return nil
			})
		} else {
			err = db.Transaction(func(tx *gorm.DB) error {
				var task Task

				if tErr := tx.Preload("Contract").
					Preload("Product.Type").
					Preload("ProductAttribute").
					First(&task, payment.TaskID).Error; tErr != nil {
					return tErr
				}

				//重新计算提成
				payment.TheoreticalPushMoney, payment.Fine, payment.PushMoney, payment.BusinessMoney = calculatePercentage(payment, &task)
				tempMoney := payment.Money - paymentBak.Money
				tempOldPushMoney1 := paymentBak.PushMoney * 0.5
				tempOldPushMoney2 := paymentBak.PushMoney - tempOldPushMoney1
				tempPushMoney1 := payment.PushMoney*0.5 - tempOldPushMoney1
				tempPushMoney2 := payment.PushMoney - payment.PushMoney*0.5 - tempOldPushMoney2
				tempBusinessMoney := payment.BusinessMoney - paymentBak.BusinessMoney
				////生成历史记录
				var officeBak Office
				if tErr := tx.First(&officeBak, task.Contract.OfficeID).Error; tErr != nil {
					return tErr
				}
				historyOffice := HistoryOffice{
					OfficeID:            task.Contract.OfficeID,
					EmployeeID:          payment.EmployeeID,
					OldBusinessMoney:    officeBak.BusinessMoney,
					OldMoney:            officeBak.Money,
					OldMoneyCold:        officeBak.MoneyCold,
					OldTargetLoad:       officeBak.TargetLoad,
					ChangeBusinessMoney: tempBusinessMoney,
					ChangeMoney:         tempPushMoney1,
					ChangeMoneyCold:     tempPushMoney2,
					ChangeTargetLoad:    0,
					NewBusinessMoney:    officeBak.BusinessMoney + payment.BusinessMoney,
					NewMoney:            officeBak.Money + tempPushMoney1,
					NewMoneyCold:        officeBak.MoneyCold + tempPushMoney2,
					NewTargetLoad:       0,
					CreateDate:          XDate{Time: time.Now()},
					Remark:              "普通合同(" + task.Contract.No + ")回款编辑",
				}
				//更新回款记录
				var maps = make(map[string]interface{})
				maps["employee_id"] = payment.EmployeeID
				maps["payment_date"] = payment.PaymentDate
				maps["money"] = payment.Money
				maps["theoretical_push_money"] = payment.TheoreticalPushMoney
				maps["fine"] = payment.Fine
				maps["push_money"] = payment.PushMoney
				maps["business_money"] = payment.BusinessMoney
				if tErr := tx.Model(&Payment{}).Where("id = ?", paymentBak.ID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				//更新合同数据
				if tErr := tx.Exec("UPDATE contract SET payment_total_amount = payment_total_amount + ? WHERE id = ?", tempMoney, paymentBak.ContractID).Error; tErr != nil {
					return tErr
				}
				//更新任务数据
				if tErr := tx.Exec("UPDATE task SET payment_total_price = payment_total_price + ? WHERE id = ?", tempMoney, paymentBak.TaskID).Error; tErr != nil {
					return tErr
				}
				//更新办事处数据
				if task.Product.Type.IsTaskLoad {
					if tErr := tx.Exec("UPDATE office SET target_load = target_load + ?, money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempMoney, tempPushMoney1, tempPushMoney2, tempBusinessMoney, contractBak.OfficeID).Error; tErr != nil {
						return tErr
					}
					historyOffice.ChangeTargetLoad = tempMoney
					historyOffice.NewTargetLoad = officeBak.TargetLoad + tempMoney
				} else {
					if tErr := tx.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE id = ?", tempPushMoney1, tempPushMoney2, tempBusinessMoney, contractBak.OfficeID).Error; tErr != nil {
						return tErr
					}
					historyOffice.NewTargetLoad = officeBak.TargetLoad
				}
				if tErr := InsertHistoryOffice(&historyOffice, tx); tErr != nil {
					return tErr
				}
				return nil
			})
		}
	} else {
		code = msg.FAIL
	}
	if err != nil {
		code = msg.ERROR
	}
	code = msg.SUCCESS
	return
}
