package twopointers

import "strings"

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1

	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
