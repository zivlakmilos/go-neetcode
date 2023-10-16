package arraysandhashing

import "testing"

func TestValidAnagram(t *testing.T) {
	inputs := [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
	}
	outputs := []bool{
		true,
		false,
	}

	for i := range inputs {
		res := isAnagram(inputs[i][0], inputs[i][1])
		if res != outputs[i] {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
