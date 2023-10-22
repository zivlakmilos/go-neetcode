package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	stack := Constructor()
	stack.Push(-2)
	fmt.Printf("%v\n", stack.GetMin())
	stack.Push(0)
	fmt.Printf("%v\n", stack.GetMin())
	stack.Push(-3)
	fmt.Printf("%v\n", stack.GetMin())

	res := stack.GetMin()
	if res != -3 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 0, nil, -3, res)
	}

	stack.Pop()
	fmt.Printf("%v\n", stack.Top())

	res = stack.Top()
	if res != 0 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 1, nil, 0, res)
	}

	res = stack.GetMin()
	if res != -2 {
		t.Fatalf("TEST %d FAILED: input: %v, expected: %v, output: %v", 2, nil, -2, res)
	}
}
