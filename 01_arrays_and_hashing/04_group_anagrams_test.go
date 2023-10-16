package arraysandhashing

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestGroupAnagrams(t *testing.T) {
	inputs := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}
	outputs := [][][]string{
		{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		{{""}},
		{{"a"}},
	}

	for i := range inputs {
		res := groupAnagrams(inputs[i])
		if len(res) != len(outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}

		if !shared.CompareArrayOfArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
