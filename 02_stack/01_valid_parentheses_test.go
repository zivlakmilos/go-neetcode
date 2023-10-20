package stack

import "testing"

func TestValidParenthases(t *testing.T) {
	inputs := []string{
		"()",
		"()[]{}",
		"(]",
	}
	outputs := []bool{
		true,
		true,
		false,
	}

	for i := range inputs {
		res := isValid(inputs[i])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
