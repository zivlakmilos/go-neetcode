package arraysandhashing

import "testing"

func TestDuplicates(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	outputs := []bool{
		true,
		false,
		true,
	}

	for i := range inputs {
		res := containsDuplicate(inputs[i])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
