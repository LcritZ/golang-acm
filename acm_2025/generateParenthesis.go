package acm_2025

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)

	var backtrack func(curr string, left, right int)
	backtrack = func(curr string, left, right int) {
		if len(curr) == 2*n {
			temp := curr
			res = append(res, temp)
		}

		if left < n {
			curr += "("
			backtrack(curr, left+1, right)
			curr = curr[0 : len(curr)-1]
		}

		if right < left {
			curr += ")"
			backtrack(curr, left, right+1)
			curr = curr[0 : len(curr)-1]
		}
	}

	backtrack("", 0, 0)

	return res
}
