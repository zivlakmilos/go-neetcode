package stack

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := 0

	for _, str := range tokens {
		if str == "+" {
			stack[top-2] = stack[top-2] + stack[top-1]
			top--
		} else if str == "-" {
			stack[top-2] = stack[top-2] - stack[top-1]
			top--
		} else if str == "*" {
			stack[top-2] = stack[top-2] * stack[top-1]
			top--
		} else if str == "/" {
			stack[top-2] = stack[top-2] / stack[top-1]
			top--
		} else {
			stack[top], _ = strconv.Atoi(str)
			top++
		}
	}

	return stack[0]
}
