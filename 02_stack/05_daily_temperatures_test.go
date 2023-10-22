package stack

import (
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	inputs := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{30, 60, 90},
	}
	outputs := [][]int{
		{1, 1, 4, 2, 1, 1, 0, 0},
		{1, 1, 1, 0},
		{1, 1, 0},
	}

	for i := range inputs {
		res := dailyTemperatures(inputs[i])
		for j := range res {
			if res[j] != outputs[i][j] {
				t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
			}
		}
	}
}
