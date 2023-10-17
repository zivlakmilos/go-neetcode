package arraysandhashing

import (
	"testing"

	"github.com/zivlakmilos/go-neetcode/shared"
)

func TestTopKFrequent(t *testing.T) {
	inputs := [][][]int{
		{{1, 1, 1, 2, 2, 3}, {2}},
		{{1}, {1}},
	}
	outputs := [][]int{
		{1, 2},
		{1},
	}

	for i := range inputs {
		res := topKFrequent(inputs[i][0], inputs[i][1][0])
		if !shared.CompareArrays(res, outputs[i]) {
			t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", i, inputs[i], outputs[i], res)
		}
	}
}
