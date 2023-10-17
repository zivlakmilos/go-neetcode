package arraysandhashing

func topKFrequent(nums []int, k int) []int {
	store := map[int]int{}

	for _, num := range nums {
		store[num]++
	}

	bucket := make([][]int, len(nums)+1)

	for key, val := range store {
		bucket[val] = append(bucket[val], key)
	}

	res := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		for _, val := range bucket[i] {
			res = append(res, val)
			k--
		}
	}

	return res
}
