package api

import (
	"oa-backend/models"
	"oa-backend/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddInvoice(c *gin.Context) {
	var invoice models.Invoice
	_ = c.ShouldBindJSON(&invoice)

	invoice.EmployeeID = c.MustGet("employee").(int)

	code = models.GeneralInsert(&invoice)
	msg.Message(c, code, nil)
}

func DelInvoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		code = models.DeleteInvoice(id)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, nil)
}

func EditInvoice(c *gin.Context) {
	var invoice models.Invoice
	_ = c.ShouldBindJSON(&invoice)

	var maps = make(map[string]interface{})
	maps["employee_id"] = c.MustGet("employee").(int)
	maps["no"] = invoice.No
	maps["money"] = invoice.Money

	code = models.GeneralUpdate(&models.Invoice{}, invoice.ID, maps)

	msg.Message(c, code, nil)
}
