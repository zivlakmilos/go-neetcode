package arraysandhashing

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 26)

	for i := range s {
		chars[s[i]-'a']++
	}

	for i := range t {
		chars[t[i]-'a']--
	}

	for _, count := range chars {
		if count != 0 {
			return false
		}
	}

	return true
}
