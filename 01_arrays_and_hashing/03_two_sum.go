package arraysandhashing

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int, len(nums))

	for i, n := range nums {
		diff := target - n
		if prevI, ok := diffs[diff]; ok {
			if ok {
				return []int{prevI, i}
			}
		}

		diffs[n] = i
	}

	return []int{0, 0}
}
