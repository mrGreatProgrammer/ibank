package stats

import (
	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

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
