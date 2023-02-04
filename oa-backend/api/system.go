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
	var maps = make(map[string]interface{})
	maps["is_ranking"] = true
	code = models.GeneralSelectAll(&offices, maps)
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
			offices[i].FinalPercentages = 0
		} else {
			offices[i].FinalPercentages = round((offices[i].TargetLoad / offices[i].TaskLoad * 100), 2)
		}
	}
	sort.SliceStable(offices, func(i, j int) bool {
		return offices[i].FinalPercentages > offices[j].FinalPercentages
	})

	msg.Message(c, code, map[string]interface{}{"offices": offices, "productTypes": productTypes})
}

func IceGet(c *gin.Context) {
	msg.Message(c, msg.SUCCESS, models.Ice)
}

func IceStart(c *gin.Context) {
	code = models.SettlementStart()
	msg.Message(c, code, nil)
}

func IceEnd(c *gin.Context) {
	code = models.SettlementEnd()
	msg.Message(c, code, nil)
}

func IceSubmit(c *gin.Context) {
	var office models.Office
	_ = c.ShouldBindJSON(&office)

	var maps = make(map[string]interface{})
	maps["is_submit"] = true
	maps["next_task_load"] = office.NextTaskLoad
	maps["last_year_money"] = office.LastYearMoney

	code = models.GeneralUpdate(&models.Office{}, office.ID, maps)
	msg.Message(c, code, nil)
}

func SettSubmit(c *gin.Context) {
	var employee models.Employee
	var employees []models.Employee
	_ = c.ShouldBindJSON(&employees)
	models.GeneralSelect(&employee, c.MustGet("employeeID").(int), nil)

	code = models.SettSubmit(&employee, employees)
	msg.Message(c, code, nil)
}

func SettApprove(c *gin.Context) {
	var office models.Office
	var employees []models.Employee
	_ = c.ShouldBindJSON(&office)

	models.GeneralSelectAll(&employees, map[string]interface{}{"office_id": office.ID})

	if office.IsPass {
		code = models.SettApprove(c.MustGet("employeeID").(int), office.ID, employees)
	} else {
		var maps = make(map[string]interface{})
		maps["is_set_submit"] = -1
		code = models.GeneralUpdate(&models.Office{}, office.ID, maps)
	}
	msg.Message(c, code, nil)
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
