package stack

import "testing"

func TestEvalRPN(t *testing.T) {
	inputs := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}
	outputs := []int{
		9,
		6,
		22,
	}

	for i := range inputs {
		res := evalRPN(inputs[i])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
