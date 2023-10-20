package stack

func isValid(s string) bool {
	parenthases := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, len(s))
	top := 0

	for i := range s {
		p, ok := parenthases[s[i]]
		if ok {
			if top == 0 || stack[top-1] != p {
				return false
			}

			top--
		} else {
			stack[top] = s[i]
			top++
		}
	}

	return top == 0
}
