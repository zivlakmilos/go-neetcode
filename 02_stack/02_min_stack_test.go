package stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {

	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	res := stack.GetMin()
	if res != -3 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 0, nil, -3, res)
	}

	stack.Pop()

	res = stack.Top()
	if res != 0 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 1, nil, 0, res)
	}

	res = stack.GetMin()
	if res != -2 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 2, nil, -2, res)
	}
}
