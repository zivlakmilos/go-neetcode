package arraysandhashing

func groupAnagrams(strs []string) [][]string {
	letters := map[[26]byte][]string{}

	for _, s := range strs {
		chars := [26]byte{}

		for _, c := range s {
			chars[c-'a']++
		}

		letters[chars] = append(letters[chars], s)
	}

	res := [][]string{}

	for _, val := range letters {
		res = append(res, val)
	}

	return res
}
