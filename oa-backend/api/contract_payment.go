package api

import (
	"oa-backend/models"
	"oa-backend/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	var payment models.Payment
	_ = c.ShouldBindJSON(&payment)

	payment.CreateDate.Time = time.Now()
	payment.EmployeeID = c.MustGet("employeeID").(int)

	code = models.InsertPayment(&payment)

	msg.Message(c, code, nil)
}

func EditPayment(c *gin.Context) {
	var payment models.Payment
	_ = c.ShouldBindJSON(&payment)
}
