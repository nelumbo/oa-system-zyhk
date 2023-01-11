package models

import (
	"math"
	"oa-backend/utils/magic"
)

// 提成计算器
func calculatePercentage(payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64, businessMoney float64) {

	if task.Contract.IsSpecial {
		theoreticalPushMoney, fine, realPushMoney, businessMoney = specialPercentage(payment, task)
	} else {
		theoreticalPushMoney, fine, realPushMoney, businessMoney = commonPercentage(payment, task)
	}

	return
}

// theoreticalPushMoney 理论提成
// fine 延迟罚款
// realPushMoney 实际提成
// businessMoney 业务费
func commonPercentage(payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64, businessMoney float64) {
	var realityPrice float64
	percent := 0.01
	if task.Contract.PayType == magic.CONTRACT_PAY_TYPE_CNY {
		realityPrice = task.ProductAttribute.StandardPrice
	} else {
		realityPrice = task.ProductAttribute.StandardPriceUSD
	}
	if realityPrice == 0 || task.Price == realityPrice {
		//平价卖出
		theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages*percent, 3)
		businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
	} else if task.Price > realityPrice {
		//高价卖出

		//产品任务标准总价
		totalStandardPrice := task.ProductAttribute.StandardPrice * float64(task.Number)

		if task.PaymentTotalPrice < totalStandardPrice {
			//任务回款未达到标准总价

			//到达标准回款的差额
			difference := totalStandardPrice - task.PaymentTotalPrice
			if difference >= payment.Money {
				//该次回款不超出差额
				theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentages*percent, 3)
				businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
			} else {
				//该次回款超出差额
				//1.标准提成
				theoreticalPushMoney1 := round(difference*task.Product.Type.PushMoneyPercentages*percent, 3)
				businessMoney1 := round(difference*task.Product.Type.BusinessMoneyPercentages*percent, 3)
				//2.额外提成
				theoreticalPushMoney2 := round((payment.Money-difference)*task.Product.Type.PushMoneyPercentagesUp*percent, 3)
				theoreticalPushMoney = theoreticalPushMoney1 + theoreticalPushMoney2
				businessMoney2 := round((payment.Money-difference)*task.Product.Type.BusinessMoneyPercentagesUp*percent, 3)
				businessMoney = businessMoney1 + businessMoney2
			}
		} else {
			//任务回款达到了标准总价
			theoreticalPushMoney = round(payment.Money*task.Product.Type.PushMoneyPercentagesUp*percent, 3)
			businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentagesUp*percent, 3)
		}
	} else if task.Price < realityPrice {
		//下降后的提成
		tempPushMoneyPercentages := task.Product.Type.PushMoneyPercentages - (realityPrice-task.Price)/realityPrice*100*task.Product.Type.PushMoneyPercentagesDown
		//提成下降过低保底
		if tempPushMoneyPercentages < task.Product.Type.MinPushMoneyPercentages {
			tempPushMoneyPercentages = task.Product.Type.MinPushMoneyPercentages
		}
		//计算提成
		theoreticalPushMoney = round(payment.Money*tempPushMoneyPercentages*percent, 3)
		businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
	}

	//回款延期罚款(预存款没有)
	if task.Contract.IsPreDeposit {
		fine = 0
	} else {
		tdoa := timeDelayOfArrival(task.Contract.EndDeliveryDate, payment.PaymentDate)
		if tdoa > 60 {
			rate := float64(tdoa-60) * 2 / 1000
			if rate > 0.4 {
				rate = 0.4
			}
			fine = round(theoreticalPushMoney*rate, 3)
		}
	}
	//实际提成
	realPushMoney = theoreticalPushMoney - fine

	return
}

func specialPercentage(payment *Payment, task *Task) (theoreticalPushMoney float64, fine float64, realPushMoney float64, businessMoney float64) {
	percent := 0.01
	theoreticalPushMoney = round(payment.Money*task.PushMoneyPercentages*percent, 3)
	businessMoney = round(payment.Money*task.Product.Type.BusinessMoneyPercentages*percent, 3)
	//回款延期罚款
	if task.Contract.IsPreDeposit {
		fine = 0
	} else {
		tdoa := timeDelayOfArrival(task.Contract.EndDeliveryDate, payment.PaymentDate)
		if tdoa > 60 {
			rate := float64(tdoa-60) * 2 / 1000
			if rate > 0.4 {
				rate = 0.4
			}
			fine = round(theoreticalPushMoney*rate, 3)
		}
	}
	//实际提成
	realPushMoney = theoreticalPushMoney - fine
	return
}

func timeDelayOfArrival(endDeliveryDate XDate, paymentDate XDate) (tdoa int) {

	if endDeliveryDate.Time.IsZero() {
		return 0
	}

	t1x := endDeliveryDate.Unix()
	t2x := paymentDate.Unix()

	if t1x >= t2x || t1x == 0 {
		return 0
	}

	days := (t2x - t1x) / 86400
	daysX := (t2x - t1x) % 86400

	if daysX > 0 {
		tdoa = int(days) + 1
	} else {
		tdoa = int(days)
	}
	return
}

// 保留f的有效n位数
func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
