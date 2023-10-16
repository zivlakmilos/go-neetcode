package arraysandhashing

import "testing"

func TestTwoSum(t *testing.T) {
	inputs := [][][]int{
		{{2, 7, 11, 15}, {9}},
		{{3, 3}, {6}},
		{{3, 2, 4}, {6}},
	}
	outputs := [][]int{
		{0, 1},
		{0, 1},
		{1, 2},
	}

	for i := range inputs {
		res := twoSum(inputs[i][0], inputs[i][1][0])
		if res[0] != outputs[i][0] || res[1] != outputs[i][1] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
