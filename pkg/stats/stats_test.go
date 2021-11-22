package stats

import (
	"reflect"
	"testing"

	"github.com/mrGreatProgrammer/payment-types/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 200, Category: "auto"},
		{ID: 2, Amount: 300, Category: "food"},
		{ID: 3, Amount: 100, Category: "auto"},
		{ID: 4, Amount: 300, Category: "auto"},
		{ID: 5, Amount: 100, Category: "fun"},
	}
	expected := map[types.Category]types.Money{
		"auto": 199,
		"food": 300,
		"fun": 100,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_one(t *testing.T) {
	first := map[types.Category]types.Money{"auto":10, "food": 20}
	second := map[types.Category]types.Money{"auto": 5, "food": 3}

	expected := map[types.Category]types.Money{"auto":-5, "food":-17}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_two(t *testing.T) {
	first := map[types.Category]types.Money{"auto":10, "food": 20}
	second:= map[types.Category]types.Money{"auto": 20, "food": 20}

	expected := map[types.Category]types.Money{"auto":10,  "food": 0}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_three(t *testing.T) {
	first := map[types.Category]types.Money{"auto":10, "food": 20}
	second := map[types.Category]types.Money{"food": 20}

	expected := map[types.Category]types.Money{"auto":-10, "food":0}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_four(t *testing.T) {
	first := map[types.Category]types.Money{"auto":10, "food": 20}
	second := map[types.Category]types.Money{"auto": 10, "food": 25, "mobile": 5}

	expected := map[types.Category]types.Money{"auto":0, "food": 5, "mobile":5}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}