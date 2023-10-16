package arraysandhashing

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}

	for _, n := range nums {
		_, ok := set[n]
		if ok {
			return true
		}

		set[n] = struct{}{}
	}

	return false
}
