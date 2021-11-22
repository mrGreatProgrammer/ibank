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

func getOnlyCategories(arr []types.Payment) []types.Category {
	var a = []types.Category{}

	for _, v := range arr {
		a = append(a, v.Category)
	}

	return a
}

func printUniqueValue(arr []types.Category) map[types.Category]types.Money {
	var dict = map[types.Category]types.Money{}
	count := types.Money(1)
	for _, v := range arr {
		dict[v] += count
	}

	return dict
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	var onlyCategories = getOnlyCategories(payments)
	var amountOfEachCategories = printUniqueValue(onlyCategories)
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount / amountOfEachCategories[payment.Category]
	}

	return categories
}

func PeriodsDinamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	maxLen := map[types.Category]types.Money{}
	if len(first) > len(second) {
		maxLen = first
	} else {
		maxLen = second
	}

	for i, _ := range maxLen {
		result[i] = second[i] - first[i]
	}
	return result
}