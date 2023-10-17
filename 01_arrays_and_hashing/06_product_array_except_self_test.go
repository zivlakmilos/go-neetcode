package arraysandhashing

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestProductExceptSelf(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}
	outputs := [][]int{
		{24, 12, 8, 6},
		{0, 0, 9, 0, 0},
	}

	for i := range inputs {
		res := productExceptSelf(inputs[i])
		if !shared.CompareArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
