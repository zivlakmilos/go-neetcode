package stack

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestGenerateParenthases(t *testing.T) {
	inputs := []int{
		3,
		1,
	}
	outputs := [][]string{
		{"((()))", "(()())", "(())()", "()(())", "()()()"},
		{"()"},
	}

	for i := range inputs {
		res := generateParenthesis(inputs[i])
		if !shared.CompareArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
