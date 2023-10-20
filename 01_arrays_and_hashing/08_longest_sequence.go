package arraysandhashing

func longestConsecutive(nums []int) int {
	store := map[int]bool{}
	longest := 0

	for _, n := range nums {
		store[n] = true
	}

	for _, n := range nums {
		if store[n-1] {
			continue
		}

		count := 1
		next := n + 1
		for store[next] {
			next++
			count++
		}

		if count > longest {
			longest = count
		}
	}

	return longest
}
