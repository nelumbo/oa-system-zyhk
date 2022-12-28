package magic

const (

	//合同人民币回款
	CONTRACT_PAYTYPE_CNY = 1
	//合同美元回款
	CONTRACT_PAYTYPE_USD = 2

	//审批驳回
	CONTRACT_STATUS_REJECT = -1
	//审批中
	CONTRACT_STATUS_SAVE = 0
	//审批中
	CONTRACT_STATUS_ING = 1
	//未完成
	CONTRACT_STATUS_UNFINISHED = 2
	//已完成
	CONTARCT_STATUS_FINISH = 3

	//生产中
	CONTATCT_PRODUCTION_STATUS_ING = 1
	//生产完成
	CONTATCT_PRODUCTION_STATUS_FINISH = 2

	//回款中
	CONTATCT_COLLECTION_STATUS_ING = 1
	//回款完成
	CONTATCT_COLLECTION_STATUS_FINISH = 2

	//合同提成系统结算
	CONTATCT_PUSHMONEY_AUTO = 1
	//合同提成手动结算
	CONTATCT_PUSHMONEY_MANUAL = 2

	//驳回
	TASK_STATUS_REJECT = -1
	//待设计
	TASK_STATUS_NOT_DESIGN = 1
	//待采购
	TASK_STATUS_NOT_PURCHASE = 2
	//待入/出库
	TASK_STATUS_NOT_STORAGE = 3
	//待装配（非标设计流程独占）
	TASK_STATUS_NOT_ASSEMBLY = 4
	//待发货
	TASK_STATUS_NOT_SHIPMENT = 5
	//已发货
	TASK_STATUS_SHIPMENT = 6

	//标准/第三方有库存
	TASK_TYPE_1 = 1
	//标准/第三方无库存
	TASK_TYPE_2 = 2
	//非标准定制
	TASK_TYPE_3 = 3

	// 财务审批状态
	EXPENSE_STATUS_FAIL           = -1
	EXPENSE_STATUS_NOT_APPROVAL_1 = 1
	EXPENSE_STATUS_NOT_APPROVAL_2 = 2
	EXPENSE_STATUS_NOT_PAYMENT    = 3
	EXPENSE_STATUS_FINISH         = 4

	// 财务类型
	//补助
	EXPENSE_TYPE_BuZhu = 1
	//提成
	EXPENSE_TYPE_TiChen = 2
	//业务费
	EXPENSE_TYPE_YeWuFei = 3
	//差旅费
	EXPENSE_TYPE_CaiLvFei = 4

	//保证金状态
	BIDBOND_STATUS_FAIL         = -1
	BIDBOND_STATUS_NOT_APPROVAL = 1
	BIDBOND_STATUS_NOT_FINAL    = 2
	BIDBOND_STATUS_FINAL        = 3
)
