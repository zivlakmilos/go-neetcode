package twopointers

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestTwoSum(t *testing.T) {
	inputs := [][][]int{
		{{2, 7, 11, 15}, {9}},
		{{2, 3, 4}, {6}},
		{{-1, 0}, {-1}},
	}
	outputs := [][]int{
		{1, 2},
		{1, 3},
		{1, 2},
	}

	for i := range inputs {
		res := twoSum(inputs[i][0], inputs[i][1][0])
		if !shared.CompareArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
