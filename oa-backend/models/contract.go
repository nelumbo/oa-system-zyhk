package models

import (
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
	// PayType               int     `gorm:"type:int;comment:付款类型(1:人民币 2:美元)" json:"payType"`
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
	VendorID              int     `gorm:"type:int;comment:签订单位;default:(-)" json:"verdorID"`
	EstimatedDeliveryDate XDate   `gorm:"type:date;comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       XDate   `gorm:"type:date;comment:实际交货日期;default:(-)" json:"endDeliveryDate"`
	EndPaymentDate        XDate   `gorm:"type:date;comment:最后回款日期;default:(-)" json:"endPaymentDate"`
	InvoiceType           int     `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string  `gorm:"type:varchar(300);comment:开票内容" json:"invoiceContent"`
	PaymentContent        string  `gorm:"type:varchar(300);comment:付款方式" json:"paymentContent"`
	IsSpecial             bool    `gorm:"type:boolean;comment:是否是特殊合同" json:"isSpecial"`
	IsPreDeposit          bool    `gorm:"type:boolean;comment:是否是预存款合同" json:"isPreDeposit"`
	PreDeposit            float64 `gorm:"type:decimal(20,6);comment:可用预存款金额" json:"preDeposit"`
	PreDepositRecord      float64 `gorm:"type:decimal(20,6);comment:合同预存款金额" json:"preDepositRecord"`
	IsRMB                 bool    `gorm:"type:boolean;comment:付款类型是否为人民币" json:"isRMB"`
	MoneyType             string  `gorm:"type:varchar(50);comment:具体付款类型" json:"moneyType"`
	TotalAmount           float64 `gorm:"type:decimal(20,6);comment:总金额" json:"totalAmount"`
	PaymentTotalAmount    float64 `gorm:"type:decimal(20,6);comment:回款总金额(人民币)" json:"paymentTotalAmount"`
	Remarks               string  `gorm:"type:varchar(300);comment:备注" json:"remarks"`
	ProductionStatus      int     `gorm:"type:int;comment:生产状态(1:生产中 2:生产完成)" json:"productionStatus"`
	CollectionStatus      int     `gorm:"type:int;comment:回款状态(1:回款中 2:回款完成)" json:"collectionStatus"`
	Status                int     `gorm:"type:int;comment:状态(-1:审批驳回 0:暂存 1:待审批 2:未完成 3:已完成)" json:"status"`
	AuditorID             int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateDate            XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate             XDate   `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`

	Region   Region    `gorm:"foreignKey:RegionID" json:"region"`
	Office   Office    `gorm:"foreignKey:OfficeID" json:"office"`
	Employee Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	Customer Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Vendor   Vendor    `gorm:"foreignKey:VendorID" json:"verdor"`
	Auditor  Employee  `gorm:"foreignKey:AuditorID" json:"auditor"`
	Tasks    []Task    `json:"tasks"`
	Invoices []Invoice `json:"invoices"`
	Payments []Payment `json:"payments"`

	IsPass bool `gorm:"-" json:"isPass"`
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

	ID                    int     `gorm:"primary_key" json:"id"`
	IsDelete              bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
	ContractID            int     `gorm:"type:int;comment:合同ID" json:"contractID"`
	ProductID             int     `gorm:"type:int;comment:产品ID" json:"productID"`
	ProductAttributeID    int     `gorm:"type:int;comment:产品属性ID" json:"productAttributeID"`
	Number                int     `gorm:"type:int;comment:数量" json:"number"`
	Price                 float64 `gorm:"type:decimal(20,6);comment:单价" json:"price"`
	TotalPrice            float64 `gorm:"type:decimal(20,6);comment:总价" json:"totalPrice"`
	PaymentTotalPrice     float64 `gorm:"type:decimal(20,6);comment:回款总金额(CNY)" json:"paymentTotalPrice"`
	Status                int     `gorm:"type:int;comment:状态(1:待设计 2:待采购 3:待入/出库 4:待装配 5:待发货 6:已发货)" json:"status"`
	Type                  int     `gorm:"type:int;comment:类型(1:标准/第三方有库存 2:标准/第三方无库存 3:非标准定制)" json:"type"`
	TechnicianManID       int     `gorm:"type:int;comment:技术负责人ID;default:(-)" json:"technicianManID"`
	PurchaseManID         int     `gorm:"type:int;comment:采购负责人ID;default:(-)" json:"purchaseManID"`
	InventoryManID        int     `gorm:"type:int;comment:仓库负责人ID;default:(-)" json:"inventoryManID"`
	ShipmentManID         int     `gorm:"type:int;comment:物流人员ID;default:(-)" json:"shipmentManID"`
	Remarks               string  `gorm:"type:varchar(600);comment:业务备注" json:"remarks"`
	PushMoney             float64 `gorm:"type:decimal(20,6);comment:提成(元)" json:"pushMoney"`
	PushMoneyPercentages  float64 `gorm:"type:decimal(20,6);comment:特殊合同提成百分比" json:"pushMoneyPercentages"`
	TechnicianDays        int     `gorm:"type:int;comment:技术预计工作天数;default:(-)" json:"technicianDays"`
	PurchaseDays          int     `gorm:"type:int;comment:采购预计工作天数;default:(-)" json:"purchaseDays"`
	TechnicianStartDate   XDate   `gorm:"type:date;comment:技术接到工作日期;default:(-)" json:"technicianStartDate"`
	TechnicianRealEndDate XDate   `gorm:"type:date;comment:技术实际提交工作日期;default:(-)" json:"technicianRealEndDate"`
	PurchaseStartDate     XDate   `gorm:"type:date;comment:采购接到工作日期;default:(-)" json:"purchaseStartDate"`
	PurchaseRealEndDate   XDate   `gorm:"type:date;comment:采购实际提交工作日期;default:(-)" json:"purchaseRealEndDate"`
	InventoryStartDate    XDate   `gorm:"type:date;comment:仓库接到工作日期;default:(-)" json:"inventoryStartDate"`
	InventoryRealEndDate  XDate   `gorm:"type:date;comment:仓库实际提交工作日期;default:(-)" json:"inventoryRealEndDate"`
	ShipmentStartDate     XDate   `gorm:"type:date;comment:物流接到工作日期;default:(-)" json:"shipmentStartDate"`
	ShipmentRealEndDate   XDate   `gorm:"type:date;comment:物流实际提交工作日期;default:(-)" json:"shipmentRealEndDate"`
	AuditorID             int     `gorm:"type:int;comment:审核员ID;default:(-)" json:"auditorID"`
	CreateDate            XDate   `gorm:"type:date;comment:创建日期;default:(-)" json:"createDate"`
	AuditDate             XDate   `gorm:"type:date;comment:审核日期;default:(-)" json:"auditDate"`

	Contract      Contract `gorm:"foreignKey:ContractID" json:"contract"`
	Product       Product  `gorm:"foreignKey:ProductID" json:"product"`
	TechnicianMan Employee `gorm:"foreignKey:TechnicianManID" json:"technicianMan"`
	PurchaseMan   Employee `gorm:"foreignKey:PurchaseManID" json:"purchaseMan"`
	InventoryMan  Employee `gorm:"foreignKey:InventoryManID" json:"inventoryMan"`
	ShipmentMan   Employee `gorm:"foreignKey:ShipmentManID" json:"shipmentMan"`
	Auditor       Employee `gorm:"foreignKey:AuditorID" json:"auditor"`

	IsPass bool `gorm:"-" json:"isPass"`
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

func ApproveContract(contractBak *Contract, maps map[string]interface{}) (code int) {
	if contractBak.IsPass {
		var no string
		no, code = createNo(contractBak)
		if code == msg.SUCCESS {
			maps["no"] = no
			maps["status"] = magic.CONTRACT_STATUS_NOT_FINISH
			maps["production_status"] = magic.CONTATCT_PRODUCTION_STATUS_ING
			maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_ING

			err = db.Transaction(func(tx *gorm.DB) error {
				var contract Contract
				if tErr := tx.Preload("Tasks").Where("is_delete = ?", false).First(&contract, contractBak.ID).Error; tErr != nil {
					return tErr
				}
				//业务员累计合同数目+1
				if tErr := tx.Exec("UPDATE employee SET contract_count = contract_count + 1 WHERE id = ?", contract.EmployeeID).Error; tErr != nil {
					return tErr
				}
				//更新合同信息
				if tErr := tx.Model(&Contract{ID: contractBak.ID}).Updates(maps).Error; tErr != nil {
					return tErr
				}
				//产品可售库存减少
				for i := range contract.Tasks {
					if tErr := tx.Exec("UPDATE product SET number = number - ? WHERE id = ?", contract.Tasks[i].Number, contract.Tasks[i].ProductID).Error; tErr != nil {
						return tErr
					}
				}
				//修改合同产品任务的开始时间
				t := time.Now().Format("2006-01-02")
				if tErr := tx.Model(&Task{}).Where("contract_id = ? AND type = ?", contractBak.ID, magic.TASK_TYPE_1).Update("inventory_start_date", t).Error; tErr != nil {
					return tErr
				}
				if tErr := tx.Model(&Task{}).Where("contract_id = ? AND type = ?", contractBak.ID, magic.TASK_TYPE_2).Update("purchase_start_date", t).Error; tErr != nil {
					return tErr
				}
				if tErr := tx.Model(&Task{}).Where("contract_id = ? AND type = ?", contractBak.ID, magic.TASK_TYPE_3).Update("technician_start_date", t).Error; tErr != nil {
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
			if tErr := tx.Model(&Contract{ID: contractBak.ID}).Updates(maps).Error; tErr != nil {
				return tErr
			}
			//删除子任务
			if tErr := tx.Model(&Task{ContractID: contractBak.ID}).Update("is_delete", true).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	}
	return
}

// employeeID 记录用
func RejectContract(contract *Contract, employeeID int) (code int) {
	err = db.Transaction(func(tx *gorm.DB) error {
		var payments []Payment
		if tErr := tx.Preload("Tasks.Product.Type").Where("contract_id = ?", contract.ID).Find(&payments).Error; tErr != nil {
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

		//TODO 历史记录

		//办事处记录变更
		if tErr := tx.Exec("UPDATE office SET target_load = target_load - ?,money = money - ?, money_cold = money_cold - ?, business_money = business_money - ? WHERE id = ?", tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney, contract.OfficeID).Error; tErr != nil {
			return tErr
		}

		//合同状态变更为驳回
		if tErr := tx.Model(&Contract{ID: contract.ID}).Update("status", magic.CONTRACT_STATUS_REJECT).Error; tErr != nil {
			return tErr
		}

		//删除任务，回款记录，支票
		if tErr := tx.Model(&Task{ContractID: contract.ID}).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Delete(&Payment{ContractID: contract.ID}).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		if tErr := tx.Model(&Invoice{ContractID: contract.ID}).Update("is_delete", true).Error; tErr != nil {
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
	db.Preload("Region").Preload("Office").Preload("Employee").Preload("Customer").Preload("ContractUnit").
		Preload("Tasks.Product.Type").
		Preload("Tasks.TechnicianMan").Preload("Tasks.PurchaseMan").
		Preload("Tasks.InventoryMan").Preload("Tasks.ShipmentMan").
		Preload("Invoices.Employee").Preload("Payments.Employee").
		Where("contract.is_delete = ?", false).
		First(&contract, id)
	if contract.ID == 0 {
		return Contract{}, msg.FAIL
	}
	return contract, msg.SUCCESS
}

func SelectContracts(contractQuery *Contract, xForms *XForms) (contracts []Contract, code int) {
	var maps = make(map[string]interface{})
	maps["contract.is_delete"] = false
	//TODO 搜索条件
	tx := db.Where(maps)

	err = tx.Find(&contracts).Count(&xForms.Total).
		Limit(xForms.PageSize).Offset((xForms.PageNo - 1) * xForms.PageSize).
		Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR
	}
	return contracts, msg.SUCCESS
}

// 预存款合同添加任务
func InsertTask(contract *Contract, task *Task) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		//创建任务
		if tErr := tdb.Create(&task).Error; tErr != nil {
			return tErr
		}
		//减去合同预存款并加上合同总金额
		if tErr := tdb.Exec("UPDATE contract SET pre_deposit = pre_deposit - ? WHERE id = ?", task.TotalPrice, contract.ID).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func ApproveTask(task *Task, employeeID int) (code int) {
	//TODO
	return
}

func SelectInvoices(invoiceQuery *Invoice) (invoices []Invoice, code int) {
	err = db.Where("is_delete = ?", false).Where("contract_id = ?", invoiceQuery.ContractID).Find(&invoices).Error

	if err != nil {
		return nil, msg.ERROR
	}
	return invoices, msg.SUCCESS
}
