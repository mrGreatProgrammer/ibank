package stats

import (
	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

var theMoney types.Money

func Avg(paymetns []types.Payment) types.Money {
	for _, v := range paymetns {
		theMoney += v.Amount
	}

	calcIt := int(theMoney) / len(paymetns)

	return types.Money(calcIt)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money

	for _, v := range payments {
		if v.Category == category {
			total += v.Amount
		}
		if v.Amount <= 0 {
			total += 0
		}
	}

	return total
}