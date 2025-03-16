package leetcode

func TimeRequiredToBuy(tickets []int, k int) int {

	var res int
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			res += minInt(tickets[i], tickets[k])
		} else {
			res += minInt(tickets[i], tickets[k]-1)
		}
	}

	return res
}

func minInt(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
