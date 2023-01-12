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
		tRouter.GET("topList", api.TopList)

		tRouter.POST("employee", api.AddEmployee)
		tRouter.DELETE("employee/:id", api.DelEmployee)
		tRouter.PUT("employee/base", api.EditEmployeeBase)
		tRouter.PUT("employee/expense", api.EditEmployeeExpense)
		tRouter.PUT("employee/office", api.EditEmployeeOffice)
		tRouter.GET("employee/:id", api.QueryEmployee)
		tRouter.POST("employees", api.QueryEmployees)
		tRouter.POST("employee/all", api.QueryAllEmployee)
		tRouter.PUT("employee/reset/:id", api.ResetEmployeePwd)

		tRouter.POST("office", api.AddOffice)
		tRouter.DELETE("office/:id", api.DelOffice)
		tRouter.PUT("office/base", api.EditOfficeBase)
		tRouter.PUT("office/money", api.EditOfficeMoney)
		tRouter.POST("offices", api.QueryOffices)
		tRouter.GET("offices", api.QueryAllOffice)

		tRouter.POST("role", api.AddRole)
		tRouter.DELETE("role/:id", api.DelRole)
		tRouter.PUT("role", api.EditRole)
		tRouter.GET("role/:id", api.QueryRole)
		tRouter.POST("roles", api.QueryRoles)
		tRouter.GET("roles", api.QueryAllRole)

		tRouter.GET("permissions", api.QueryAllPermission)

		tRouter.POST("region", api.AddRegion)
		tRouter.PUT("region", api.EditRegion)
		tRouter.POST("regions", api.QueryRegions)
		tRouter.GET("regions", api.QueryAllRegion)

		tRouter.POST("vendor", api.AddVendor)
		tRouter.PUT("vendor", api.EditVendor)
		tRouter.POST("vendors", api.QueryVendors)
		tRouter.GET("vendors", api.QueryAllVendor)

		tRouter.POST("customerCompany", api.AddCustomerCompany)
		tRouter.DELETE("customerCompany/:id", api.DelCustomerCompany)
		tRouter.PUT("customerCompany", api.EditCustomerCompany)
		tRouter.POST("customerCompanys", api.QueryCustomerCompanys)

		tRouter.POST("customer", api.AddCustomer)
		tRouter.DELETE("customer/:id", api.DelCustomer)
		tRouter.PUT("customer", api.EditCustomer)
		tRouter.POST("customers", api.QueryCustomers)

		tRouter.POST("supplier", api.AddSupplier)
		tRouter.PUT("supplier", api.EditSupplier)
		tRouter.POST("suppliers", api.QuerySuppliers)

		tRouter.POST("productType", api.AddProductType)
		tRouter.PUT("productType", api.EditProductType)
		tRouter.POST("productTypes", api.QueryProductTypes)
		tRouter.GET("productTypes", api.QueryAllProductType)

		tRouter.POST("product", api.AddProduct)
		tRouter.PUT("product/base", api.EditProductBase)
		tRouter.PUT("product/attribute", api.EditProductAttribute)
		tRouter.PUT("product/number", api.EditProductNumber)
		tRouter.GET("product/:id", api.QueryProduct)
		tRouter.POST("products", api.QueryProducts)

		tRouter.POST("expense", api.AddExpense)
		tRouter.DELETE("expense/:id", api.DelExpense)
		tRouter.PUT("expense/approve", api.ApprovalExpense)
		tRouter.POST("expenses", api.QueryExpenses)

		tRouter.POST("bidbond", api.AddBidbond)
		tRouter.DELETE("bidbond/:id", api.DelBidbond)
		tRouter.PUT("bidbond", api.EditBidbond)
		tRouter.PUT("bidbond/approve", api.ApproveBidbond)
		tRouter.POST("bidbonds", api.QueryBidbonds)

		tRouter.POST("predesign", api.AddPredesign)
		tRouter.DELETE("predesign/:id", api.DelPredesign)
		tRouter.PUT("predesign", api.EditPredesign)
		tRouter.PUT("predesign/approve", api.ApprovePredesign)
		tRouter.GET("predesign/:id", api.QueryPredesign)
		tRouter.POST("predesigns", api.QueryPredesigns)

		tRouter.PUT("predesignTask/submit", api.SubmitPredesignTask)
		tRouter.PUT("predesignTask/approve", api.ApprovePredesignTask)
		tRouter.POST("predesignTasks", api.QueryPredesignTasks)

		tRouter.POST("contract/save", api.SaveContract)
		tRouter.POST("contract", api.AddContract)
		tRouter.DELETE("contract/:id", api.DelContract)
		tRouter.PUT("contract/edd", api.EditContractEstimatedDeliveryDate)
		tRouter.PUT("contract/approve", api.ApproveContract)
		tRouter.PUT("contract/final", api.FinalContract)
		tRouter.PUT("contract/reset", api.ResetContractCollectionStatus)
		tRouter.PUT("contract/reject", api.RejectContract)
		tRouter.GET("contract/:id", api.QueryContract)
		tRouter.POST("contracts", api.QueryContracts)

		tRouter.POST("task", api.AddTask)
		tRouter.PUT("task/distribute", api.DistributeTask)
		tRouter.PUT("task/next", api.NextTask)
		tRouter.PUT("task/reset", api.ResetTask)
		tRouter.PUT("task/reject", api.RejectTask)
		tRouter.DELETE("task/:id", api.DelTask)

		tRouter.POST("payment", api.AddPayment)
		tRouter.PUT("payment", api.EditPayment)
		tRouter.PUT("payment/final", api.FinalPayment)

		tRouter.POST("invoice", api.AddInvoice)
		tRouter.DELETE("invoice/:id", api.DelInvoice)
		tRouter.PUT("invoice", api.EditInvoice)

		tRouter.POST("purchasing/save", api.SavePurchasing)
		tRouter.PUT("purchasing/submit", api.SubmitPurchasing)
		tRouter.POST("purchasing", api.AddPurchasing)
		tRouter.DELETE("purchasing/:id", api.DelPurchasing)
		tRouter.PUT("purchasing/approve", api.ApprovePurchasing)
		tRouter.PUT("purchasing/product/final", api.FinalPurchasingProductStatus)
		tRouter.PUT("purchasing/pay/final", api.FinalPurchasingPayStatus)
		tRouter.PUT("purchasing/invoice/final", api.FinalPurchasingInvoiceStatus)
		tRouter.PUT("purchasing/final", api.FinalPurchasing)
		tRouter.POST("purchasings", api.QueryPurchasings)
		tRouter.POST("purchasing/all", api.QueryAllPurchasing)

		tRouter.POST("historyOffices", api.QueryHistoryOffices)
		tRouter.POST("historyEmployees", api.QueryHistoryEmployees)
	}

	_ = r.Run(config.SystemConfig.Server.Port)

}
