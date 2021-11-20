package stats

import (
	"fmt"

	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

func ExampleAvg() {

	var paymetns = []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   100,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   300,
			Category: "auto",
		},
		{
			ID:       4,
			Amount:   500,
			Category: "auto",
		},
		{
			ID:       5,
			Amount:   400,
			Category: "auto",
		},
	}

	fmt.Println(Avg(paymetns))

	//Output:
	// 300

}

func ExampleTotalInCategory() {
	
	var payments = []types.Payment{
		{
			ID: 1,
			Amount: 200,
			Category: "auto",
		},
		{
			ID: 2,
			Amount: 100,
			Category: "auto",
		},
		{
			ID: 3,
			Amount: 300,
			Category: "restaurant",
		},
		{
			ID: 4,
			Amount: 500,
			Category: "restaurant",
		},
		{
			ID: 5,
			Amount: 400,
			Category: "pharmacy",
		},
		{
			ID: 6,
			Amount: 400,
			Category: "pharmacy",
		},
	}

	fmt.Println(TotalInCategory(payments, "pharmacy"))
	fmt.Println(TotalInCategory(payments, "auto"))
	fmt.Println(TotalInCategory(payments, "restaurant"))

	//Output:
	// 800
	// 300
	// 800
}
