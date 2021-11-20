package stats

import (
	"fmt"

	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

func ExampleAvg() {

	var payments = []types.Payment{
		{
			ID:       1,
			Amount:   2,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   5,
			Category: "auto",
			Status:   types.StatusInProgress,
		},
		{
			ID:       3,
			Amount:   4,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       4,
			Amount:   6,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       5,
			Amount:   3,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       6,
			Amount:   3,
			Category: "auto",
			Status:   types.StatusFail,
		},
		
	}

	fmt.Println(Avg(payments))

	//Output:
	// 4

}

func ExampleTotalInCategory() {

	var payments = []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "auto",
			Status:   types.StatusInProgress,
		},
		{
			ID:       2,
			Amount:   100,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   300,
			Category: "restaurant",
			Status:   types.StatusOk,
		},
		{
			ID:       4,
			Amount:   500,
			Category: "restaurant",
			Status:   types.StatusOk,
		},
		{
			ID:       5,
			Amount:   400,
			Category: "pharmacy",
			Status:   types.StatusInProgress,
		},
		{
			ID:       6,
			Amount:   400,
			Category: "pharmacy",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(TotalInCategory(payments, "pharmacy"))
	fmt.Println(TotalInCategory(payments, "auto"))
	fmt.Println(TotalInCategory(payments, "restaurant"))

	//Output:
	// 400
	// 300
	// 800
}

