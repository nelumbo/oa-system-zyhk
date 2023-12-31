package routes

import (
	"oa-backend/api"
	"oa-backend/config"
	"oa-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(config.SystemConfig.Server.Mode)

	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Cors())
	router := r.Group("api")
	{
		router.POST("login", api.Login)

	}

	tRouter := r.Group("api")
	tRouter.Use(middleware.CheckToken())
	// {
	// 	tRouter.GET("me", api.QueryMe)
	// 	tRouter.POST("my/expenses", api.QueryMyExpenses)
	// 	tRouter.POST("my/bidbonds", api.QueryMyBidbonds)
	// 	tRouter.POST("my/predesigns", api.QueryMyPredesigns)
	// 	tRouter.POST("my/predesignTasks", api.QueryMyPredesignTasks)
	// 	tRouter.POST("my/contracts", api.QueryMyContracts)
	// 	tRouter.POST("my/contracts/save", api.QueryMySaveContracts)
	// 	tRouter.POST("my/tasks", api.QueryMyTasks)
	// 	tRouter.POST("my/purchasings/save", api.QueryMySavePurchasings)
	// 	tRouter.POST("my/historys", api.QueryMyHistorys)
	// 	tRouter.POST("my/productTrials", api.QueryMyProductTrials)
	// 	tRouter.POST("my/pwd", api.UpdatePwd)
	// 	tRouter.GET("topList", api.TopList)

	// 	tRouter.POST("employee", api.AddEmployee)
	// 	tRouter.DELETE("employee/:id", api.DelEmployee)
	// 	tRouter.PUT("employee/base", api.EditEmployeeBase)
	// 	tRouter.PUT("employee/expense", api.EditEmployeeExpense)
	// 	tRouter.PUT("employee/office", api.EditEmployeeOffice)
	// 	tRouter.GET("employee/:id", api.QueryEmployee)
	// 	tRouter.POST("employees", api.QueryEmployees)
	// 	tRouter.POST("employee/all", api.QueryAllEmployee)
	// 	tRouter.PUT("employee/reset/:id", api.ResetEmployeePwd)

	// 	tRouter.POST("office", api.AddOffice)
	// 	tRouter.DELETE("office/:id", api.DelOffice)
	// 	tRouter.PUT("office/base", api.EditOfficeBase)
	// 	tRouter.PUT("office/money", api.EditOfficeMoney)
	// 	tRouter.POST("offices", api.QueryOffices)
	// 	tRouter.GET("offices", api.QueryAllOffice)

	// 	tRouter.POST("role", api.AddRole)
	// 	tRouter.DELETE("role/:id", api.DelRole)
	// 	tRouter.PUT("role", api.EditRole)
	// 	tRouter.GET("role/:id", api.QueryRole)
	// 	tRouter.POST("roles", api.QueryRoles)
	// 	tRouter.GET("roles", api.QueryAllRole)

	// 	tRouter.GET("permissions", api.QueryAllPermission)

	// 	tRouter.POST("region", api.AddRegion)
	// 	tRouter.PUT("region", api.EditRegion)
	// 	tRouter.POST("regions", api.QueryRegions)
	// 	tRouter.GET("regions", api.QueryAllRegion)

	// 	tRouter.POST("vendor", api.AddVendor)
	// 	tRouter.PUT("vendor", api.EditVendor)
	// 	tRouter.POST("vendors", api.QueryVendors)
	// 	tRouter.GET("vendors", api.QueryAllVendor)

	// 	tRouter.POST("customerCompany", api.AddCustomerCompany)
	// 	tRouter.DELETE("customerCompany/:id", api.DelCustomerCompany)
	// 	tRouter.PUT("customerCompany", api.EditCustomerCompany)
	// 	tRouter.POST("customerCompanys", api.QueryCustomerCompanys)

	// 	tRouter.POST("customer", api.AddCustomer)
	// 	tRouter.DELETE("customer/:id", api.DelCustomer)
	// 	tRouter.PUT("customer", api.EditCustomer)
	// 	tRouter.POST("customers", api.QueryCustomers)

	// 	tRouter.POST("supplier", api.AddSupplier)
	// 	tRouter.PUT("supplier", api.EditSupplier)
	// 	tRouter.POST("suppliers", api.QuerySuppliers)

	// 	tRouter.POST("productType", api.AddProductType)
	// 	tRouter.PUT("productType", api.EditProductType)
	// 	tRouter.POST("productTypes", api.QueryProductTypes)
	// 	tRouter.GET("productTypes", api.QueryAllProductType)

	// 	tRouter.POST("product", api.AddProduct)
	// 	tRouter.PUT("product/base", api.EditProductBase)
	// 	tRouter.PUT("product/attribute", api.EditProductAttribute)
	// 	tRouter.PUT("product/number", api.EditProductNumber)
	// 	tRouter.GET("product/:id", api.QueryProduct)
	// 	tRouter.POST("products", api.QueryProducts)

	// 	tRouter.POST("expense", api.AddExpense)
	// 	tRouter.DELETE("expense/:id", api.DelExpense)
	// 	tRouter.PUT("expense/approve", api.ApprovalExpense)
	// 	tRouter.POST("expenses", api.QueryExpenses)

	// 	tRouter.POST("bidbond", api.AddBidbond)
	// 	tRouter.DELETE("bidbond/:id", api.DelBidbond)
	// 	tRouter.PUT("bidbond", api.EditBidbond)
	// 	tRouter.PUT("bidbond/approve", api.ApproveBidbond)
	// 	tRouter.POST("bidbonds", api.QueryBidbonds)

	// 	tRouter.POST("predesign", api.AddPredesign)
	// 	tRouter.DELETE("predesign/:id", api.DelPredesign)
	// 	tRouter.PUT("predesign", api.EditPredesign)
	// 	tRouter.PUT("predesign/approve", api.ApprovePredesign)
	// 	tRouter.GET("predesign/:id", api.QueryPredesign)
	// 	tRouter.POST("predesigns", api.QueryPredesigns)

	// 	tRouter.PUT("predesignTask/submit", api.SubmitPredesignTask)
	// 	tRouter.PUT("predesignTask/approve", api.ApprovePredesignTask)
	// 	tRouter.POST("predesignTasks", api.QueryPredesignTasks)

	// 	tRouter.POST("contract/save", api.SaveContract)
	// 	tRouter.POST("contract", api.AddContract)
	// 	tRouter.DELETE("contract/:id", api.DelContract)
	// 	tRouter.PUT("contract/edd", api.EditContractEstimatedDeliveryDate)
	// 	tRouter.PUT("contract/approve", api.ApproveContract)
	// 	tRouter.PUT("contract/final", api.FinalContract)
	// 	tRouter.PUT("contract/reset", api.ResetContractCollectionStatus)
	// 	tRouter.PUT("contract/reject", api.RejectContract)
	// 	tRouter.GET("contract/:id", api.QueryContract)
	// 	tRouter.POST("contracts", api.QueryContracts)
	// 	tRouter.PUT("contract/preDeposit", api.EditContractPreDeposit)

	// 	tRouter.POST("task", api.AddTask)
	// 	tRouter.PUT("task/distribute", api.DistributeTask)
	// 	tRouter.PUT("task/next", api.NextTask)
	// 	tRouter.PUT("task/reset", api.ResetTask)
	// 	tRouter.PUT("task/reject", api.RejectTask)
	// 	tRouter.DELETE("task/:id", api.DelTask)

	// 	tRouter.POST("payment", api.AddPayment)
	// 	tRouter.PUT("payment", api.EditPayment)
	// 	tRouter.PUT("payment/final", api.FinalPayment)

	// 	tRouter.POST("invoice", api.AddInvoice)
	// 	tRouter.DELETE("invoice/:id", api.DelInvoice)
	// 	tRouter.PUT("invoice", api.EditInvoice)

	// 	tRouter.POST("purchasing/save", api.SavePurchasing)
	// 	tRouter.PUT("purchasing/submit", api.SubmitPurchasing)
	// 	tRouter.POST("purchasing", api.AddPurchasing)
	// 	tRouter.DELETE("purchasing/:id", api.DelPurchasing)
	// 	tRouter.PUT("purchasing/approve", api.ApprovePurchasing)
	// 	tRouter.PUT("purchasing/product/final", api.FinalPurchasingProductStatus)
	// 	tRouter.PUT("purchasing/pay/final", api.FinalPurchasingPayStatus)
	// 	tRouter.PUT("purchasing/invoice/final", api.FinalPurchasingInvoiceStatus)
	// 	tRouter.PUT("purchasing/final", api.FinalPurchasing)
	// 	tRouter.GET("purchasing/:id", api.QueryPurchasing)
	// 	tRouter.POST("purchasings", api.QueryPurchasings)
	// 	tRouter.POST("purchasing/all", api.QueryAllPurchasing)

	// 	tRouter.POST("historyOffices", api.QueryHistoryOffices)
	// 	tRouter.POST("historyEmployees", api.QueryHistoryEmployees)
	// 	tRouter.POST("historyProducts", api.QueryHistoryProducts)

	// 	tRouter.GET("ice", api.IceGet)
	// 	tRouter.PUT("ice/start", api.IceStart)
	// 	tRouter.PUT("ice/end", api.IceEnd)
	// 	tRouter.PUT("ice/submit", api.IceSubmit)

	// 	tRouter.PUT("sett/submit", api.SettSubmit)
	// 	tRouter.PUT("sett/approve", api.SettApprove)

	// 	tRouter.POST("productTrial", api.AddProductTrial)
	// 	tRouter.DELETE("productTrial/:id", api.DelProductTrial)
	// 	tRouter.PUT("productTrial", api.EditProductTrial)
	// 	tRouter.PUT("productTrial/next", api.NextProductTrial)
	// 	tRouter.POST("productTrials", api.QueryProductTrials)

	// 	tRouter.POST("productCalls", api.QueryProductCalls)
	// }

	{
		tRouter.GET("me", api.QueryMe)
		tRouter.POST("my/expenses", api.QueryMyExpenses)
		tRouter.POST("my/bidbonds", api.QueryMyBidbonds)
		tRouter.POST("my/predesigns", api.QueryMyPredesigns)
		tRouter.POST("my/predesignTasks", api.QueryMyPredesignTasks)
		tRouter.POST("my/contracts", api.QueryMyContracts)
		tRouter.POST("my/contracts/save", api.QueryMySaveContracts)
		tRouter.POST("my/tasks", api.QueryMyTasks)
		tRouter.POST("my/purchasings/save", api.QueryMySavePurchasings)
		tRouter.POST("my/historys", api.QueryMyHistorys)
		tRouter.POST("my/productTrials", api.QueryMyProductTrials)
		tRouter.GET("topList", api.TopList)

		tRouter.GET("employee/:id", api.QueryEmployee)
		tRouter.POST("employees", api.QueryEmployees)
		tRouter.POST("employee/all", api.QueryAllEmployee)

		tRouter.POST("offices", api.QueryOffices)
		tRouter.GET("offices", api.QueryAllOffice)

		tRouter.GET("role/:id", api.QueryRole)
		tRouter.POST("roles", api.QueryRoles)
		tRouter.GET("roles", api.QueryAllRole)

		tRouter.GET("permissions", api.QueryAllPermission)

		tRouter.POST("regions", api.QueryRegions)
		tRouter.GET("regions", api.QueryAllRegion)

		tRouter.POST("vendors", api.QueryVendors)
		tRouter.GET("vendors", api.QueryAllVendor)

		tRouter.POST("customerCompanys", api.QueryCustomerCompanys)

		tRouter.POST("customers", api.QueryCustomers)

		tRouter.POST("suppliers", api.QuerySuppliers)

		tRouter.POST("productTypes", api.QueryProductTypes)
		tRouter.GET("productTypes", api.QueryAllProductType)

		tRouter.GET("product/:id", api.QueryProduct)
		tRouter.POST("products", api.QueryProducts)

		tRouter.POST("expenses", api.QueryExpenses)

		tRouter.POST("bidbonds", api.QueryBidbonds)

		tRouter.GET("predesign/:id", api.QueryPredesign)
		tRouter.POST("predesigns", api.QueryPredesigns)

		tRouter.POST("predesignTasks", api.QueryPredesignTasks)

		tRouter.GET("contract/:id", api.QueryContract)
		tRouter.POST("contracts", api.QueryContracts)

		tRouter.GET("purchasing/:id", api.QueryPurchasing)
		tRouter.POST("purchasings", api.QueryPurchasings)
		tRouter.POST("purchasing/all", api.QueryAllPurchasing)

		tRouter.POST("historyOffices", api.QueryHistoryOffices)
		tRouter.POST("historyEmployees", api.QueryHistoryEmployees)
		tRouter.POST("historyProducts", api.QueryHistoryProducts)

		tRouter.GET("ice", api.IceGet)
		tRouter.PUT("ice/start", api.IceStart)
		tRouter.PUT("ice/end", api.IceEnd)
		tRouter.PUT("ice/submit", api.IceSubmit)

		tRouter.PUT("sett/submit", api.SettSubmit)
		tRouter.PUT("sett/approve", api.SettApprove)

		tRouter.POST("productTrials", api.QueryProductTrials)

		tRouter.POST("productCalls", api.QueryProductCalls)
	}

	t1Router := r.Group("api")
	t1Router.Use(middleware.CheckToken())
	t1Router.Use(middleware.CheckPre())

	{
		t1Router.POST("my/pwd", api.UpdatePwd)

		t1Router.POST("employee", api.AddEmployee)
		t1Router.DELETE("employee/:id", api.DelEmployee)
		t1Router.PUT("employee/base", api.EditEmployeeBase)
		t1Router.PUT("employee/expense", api.EditEmployeeExpense)
		t1Router.PUT("employee/office", api.EditEmployeeOffice)
		t1Router.PUT("employee/reset/:id", api.ResetEmployeePwd)

		t1Router.POST("office", api.AddOffice)
		t1Router.DELETE("office/:id", api.DelOffice)
		t1Router.PUT("office/base", api.EditOfficeBase)
		t1Router.PUT("office/money", api.EditOfficeMoney)

		t1Router.POST("role", api.AddRole)
		t1Router.DELETE("role/:id", api.DelRole)
		t1Router.PUT("role", api.EditRole)

		t1Router.POST("region", api.AddRegion)
		t1Router.PUT("region", api.EditRegion)

		t1Router.POST("vendor", api.AddVendor)
		t1Router.PUT("vendor", api.EditVendor)

		t1Router.POST("customerCompany", api.AddCustomerCompany)
		t1Router.DELETE("customerCompany/:id", api.DelCustomerCompany)
		t1Router.PUT("customerCompany", api.EditCustomerCompany)

		t1Router.POST("customer", api.AddCustomer)
		t1Router.DELETE("customer/:id", api.DelCustomer)
		t1Router.PUT("customer", api.EditCustomer)

		t1Router.POST("supplier", api.AddSupplier)
		t1Router.PUT("supplier", api.EditSupplier)

		t1Router.POST("productType", api.AddProductType)
		t1Router.PUT("productType", api.EditProductType)

		t1Router.POST("product", api.AddProduct)
		t1Router.PUT("product/base", api.EditProductBase)
		t1Router.PUT("product/attribute", api.EditProductAttribute)
		t1Router.PUT("product/number", api.EditProductNumber)

		t1Router.POST("expense", api.AddExpense)
		t1Router.DELETE("expense/:id", api.DelExpense)
		t1Router.PUT("expense/approve", api.ApprovalExpense)

		t1Router.POST("bidbond", api.AddBidbond)
		t1Router.DELETE("bidbond/:id", api.DelBidbond)
		t1Router.PUT("bidbond", api.EditBidbond)
		t1Router.PUT("bidbond/approve", api.ApproveBidbond)

		t1Router.POST("predesign", api.AddPredesign)
		t1Router.DELETE("predesign/:id", api.DelPredesign)
		t1Router.PUT("predesign", api.EditPredesign)
		t1Router.PUT("predesign/approve", api.ApprovePredesign)

		t1Router.PUT("predesignTask/submit", api.SubmitPredesignTask)
		t1Router.PUT("predesignTask/approve", api.ApprovePredesignTask)

		t1Router.POST("contract/save", api.SaveContract)
		t1Router.POST("contract", api.AddContract)
		t1Router.DELETE("contract/:id", api.DelContract)
		t1Router.PUT("contract/edd", api.EditContractEstimatedDeliveryDate)
		t1Router.PUT("contract/approve", api.ApproveContract)
		t1Router.PUT("contract/final", api.FinalContract)
		t1Router.PUT("contract/reset", api.ResetContractCollectionStatus)
		t1Router.PUT("contract/reject", api.RejectContract)
		t1Router.PUT("contract/preDeposit", api.EditContractPreDeposit)

		t1Router.POST("task", api.AddTask)
		t1Router.PUT("task/distribute", api.DistributeTask)
		t1Router.PUT("task/next", api.NextTask)
		t1Router.PUT("task/reset", api.ResetTask)
		t1Router.PUT("task/reject", api.RejectTask)
		t1Router.DELETE("task/:id", api.DelTask)

		t1Router.POST("payment", api.AddPayment)
		t1Router.PUT("payment", api.EditPayment)
		t1Router.PUT("payment/final", api.FinalPayment)

		t1Router.POST("invoice", api.AddInvoice)
		t1Router.DELETE("invoice/:id", api.DelInvoice)
		t1Router.PUT("invoice", api.EditInvoice)

		t1Router.POST("purchasing/save", api.SavePurchasing)
		t1Router.PUT("purchasing/submit", api.SubmitPurchasing)
		t1Router.POST("purchasing", api.AddPurchasing)
		t1Router.DELETE("purchasing/:id", api.DelPurchasing)
		t1Router.PUT("purchasing/approve", api.ApprovePurchasing)
		t1Router.PUT("purchasing/product/final", api.FinalPurchasingProductStatus)
		t1Router.PUT("purchasing/pay/final", api.FinalPurchasingPayStatus)
		t1Router.PUT("purchasing/invoice/final", api.FinalPurchasingInvoiceStatus)
		t1Router.PUT("purchasing/final", api.FinalPurchasing)

		t1Router.POST("productTrial", api.AddProductTrial)
		t1Router.DELETE("productTrial/:id", api.DelProductTrial)
		t1Router.PUT("productTrial", api.EditProductTrial)
		t1Router.PUT("productTrial/next", api.NextProductTrial)
	}

	_ = r.Run(config.SystemConfig.Server.Port)

}
