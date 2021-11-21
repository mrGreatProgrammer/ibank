package stats

import (
	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money
	var i int

	for _, v := range payments {
		if v.Status != types.StatusFail {
			i++
			avg += v.Amount
		}
	}

	lenOfTheArr := types.Money(i)
	return avg / lenOfTheArr
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money

	for _, v := range payments {
		if (v.Category == category) && (v.Amount > 0) && (v.Status != types.StatusFail) {
			total += v.Amount
		}
	}

	return total
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	avg := Avg(payments)
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += avg
	}

	return categories
}
