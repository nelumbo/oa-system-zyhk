package api

import (
	"math"
	"oa-backend/middleware"
	"oa-backend/models"
	"oa-backend/utils/msg"
	"sort"

	"github.com/gin-gonic/gin"
)

var code int
var err error

func Login(c *gin.Context) {
	var employee models.Employee
	var token string
	_ = c.ShouldBindJSON(&employee)

	employee, code = models.SelectEmployeeByPhoneAndPwd(employee.Phone, employee.Password)

	if code == msg.SUCCESS {
		token, _ = middleware.GenerateToken(employee.ID)
	}

	msg.Message(c, code, gin.H{
		"employee": employee,
		"token":    token,
	})
}

func TopList(c *gin.Context) {
	//office1:普通合同待回款量
	//office2:预付款合同待回款量
	var offices, offices1, offices2 []models.Office
	var productTypes []models.ProductType
	code = models.GeneralSelectAll(&offices, nil)
	offices1, offices2, productTypes = models.SelectNotPaymentForTopList()

	for i := range offices {

		for j := range offices1 {
			if offices[i].ID == offices1[j].ID {
				offices[i].NotPayment += offices1[j].Money
				break
			}
		}
		for k := range offices2 {
			if offices[i].ID == offices2[k].ID {
				offices[i].NotPayment += offices2[k].Money
				break
			}
		}
		if offices[i].TaskLoad == 0 {
			offices[i].FinalPercentages = 100
		} else {
			offices[i].FinalPercentages = round((offices[i].TargetLoad / offices[i].TaskLoad * 100), 2)
		}
	}
	sort.SliceStable(offices, func(i, j int) bool {
		return offices[i].FinalPercentages > offices[j].FinalPercentages
	})

	msg.Message(c, code, map[string]interface{}{"offices": offices, "productTypes": productTypes})
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
