package arraysandhashing

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 32)
	globalCount := 0

	for i := range chars {
		chars[i] = 0
	}

	for i := range s {
		chars[s[i]-'a']++
		globalCount++
	}

	for i := range t {
		chars[t[i]-'a']--
		globalCount--

		if chars[t[i]-'a'] < 0 {
			return false
		}
	}

	if globalCount > 0 {
		return false
	}

	return true
}
