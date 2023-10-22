package stack

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := 0

	for i := range temperatures {
		for top > 0 && temperatures[stack[top-1]] < temperatures[i] {
			res[stack[top-1]] = i - stack[top-1]
			top--
		}

		stack[top] = i
		top++
	}

	return res
}
