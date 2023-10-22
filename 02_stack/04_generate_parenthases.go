package stack

func generateParenthesisRec(n int, opened int, str string, res *[]string) {
	if n == 0 && opened == 0 {
		*res = append(*res, str)
		return
	}

	if n > 0 {
		generateParenthesisRec(n-1, opened+1, str+"(", res)
	}

	if opened > 0 {
		generateParenthesisRec(n, opened-1, str+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}

	generateParenthesisRec(n, 0, "", &res)

	return res
}
