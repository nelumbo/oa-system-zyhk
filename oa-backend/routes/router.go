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

		tRouter.POST("employee", api.AddEmployee)
		tRouter.DELETE("employee/:id", api.DelEmployee)
		tRouter.PUT("employee/base", api.EditEmployeeBase)
		tRouter.PUT("employee/expense", api.EditEmployeeExpense)
		tRouter.PUT("employee/office", api.EditEmployeeOffice)
		tRouter.GET("employee/:id", api.QueryEmployee)
		tRouter.POST("employees", api.QueryEmployees)
		tRouter.PUT("employee/reset/:id", api.ResetEmployeePwd)

		tRouter.POST("office", api.AddOffice)
		tRouter.DELETE("office/:id", api.DelOffice)
		tRouter.PUT("office", api.EditOffice)
		tRouter.GET("office/:id", api.QueryOffice)
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
		tRouter.GET("product/:id", api.QueryProduct)
		tRouter.POST("products", api.QueryProducts)

		tRouter.POST("expense", api.AddExpense)
		tRouter.DELETE("expense/:id", api.DelExpense)
		tRouter.PUT("expense", api.EditExpense)
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
		tRouter.GET("predesign/:id", api.QueryPredegisn)
		tRouter.POST("predesigns", api.QueryPredegisns)

		tRouter.PUT("predesignTask/submit", api.SubmitPredesignTask)
		tRouter.PUT("predesignTask/approve", api.ApprovePredesignTask)
	}

	_ = r.Run(config.SystemConfig.Server.Port)

}
