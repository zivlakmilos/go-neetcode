package twopointers

import "testing"

func TestDailyTemperatures(t *testing.T) {
	inputs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}
	outputs := []bool{
		true,
		false,
		true,
	}

	for i := range inputs {
		res := isPalindrome(inputs[i])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
