package stats

import (
	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

// var theMoney types.Money

func Avg(payments []types.Payment) types.Money {
	var avg types.Money

 for _, v:= range payments{
   if ((v.Status != "FAIL") && (v.Amount > 0)) {
     avg+=v.Amount
   }
 }

	calcIt := int(avg) / len(payments)

	return types.Money(calcIt)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money

	for _, v := range payments {
		if ((v.Category == category) && (v.Amount > 0) && (v.Status != "FAIL")){
			total += v.Amount
		}
	}

	return total
}