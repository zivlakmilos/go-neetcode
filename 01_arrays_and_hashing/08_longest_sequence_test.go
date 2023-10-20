package arraysandhashing

import (
	"testing"
)

func TestLongestSequence(t *testing.T) {
	inputs := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
	}
	outputs := []int{
		4,
		9,
	}

	for i := range inputs {
		res := longestConsecutive(inputs[i])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
