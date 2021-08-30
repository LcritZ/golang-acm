package leetcode

import "fmt"

var pick = 5

func guess(n int) int {
	if n > pick {
		return -1
	} else if n == pick {
		return 0
	} else {
		return 1
	}
}

func GuessNumber(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right - left)/2
		fmt.Println(left, right, mid)
		if guess(mid) < 0 {
			right = mid-1
		} else if guess(mid) > 0{
			left = mid+1
		} else {
			return mid
		}
	}
	return left
}
