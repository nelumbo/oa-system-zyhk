package api

import (
	"oa-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	var payment models.Payment
	_ = c.ShouldBindJSON(&payment)

	payment.CreateDate.Time = time.Now()
	payment.EmployeeID = c.MustGet("employeeID").(int)
}

func EditPayment(c *gin.Context) {

}
