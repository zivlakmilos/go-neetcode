package twopointers

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestThreeSum(t *testing.T) {
	inputs := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}
	outputs := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{},
		{{0, 0, 0}},
	}

	for i := range inputs {
		res := threeSum(inputs[i])
		if !shared.CompareArrayOfArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
